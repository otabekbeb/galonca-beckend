package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"html"
	"log"
	"strings"
	"tsl_server/api"
)

type SparePart struct {
	gorm.Model
	Title         string
	Description   string            `gorm:"type:text"`
	Categories    []ServiceCategory `gorm:"many2many:parts_category"`
	TransportType []TransportType   `gorm:"many2many:parts_transport_type"`
	CityId        uint64
	City          *City
	RegionId      uint64
	Region        *Region
	CountryId     uint64
	Country       *Country
	Address       string
	Images        string
	Phone         string
	Email         string
	Lat           float64
	Lon           float64
	OwnerID       uint
	Owner         *TslUser `gorm:"foreignKey:OwnerID"`
	Cost          uint32
	Currency      string
}

type SparePartStorage struct {
	db *gorm.DB
}

var sp *SparePartStorage

func NewSparePartStorage() *SparePartStorage {
	if sp == nil {
		sp = &SparePartStorage{
			db: GetDB(),
		}
	}
	return sp
}

func (sp *SparePart) ToAPI() *api.FindSparePartResult {

	sparePart := api.FindSparePartResult{
		Id:          uint64(sp.ID),
		CreatedAt:   sp.CreatedAt.Format("02.01.2006"),
		UpdatedAt:   sp.UpdatedAt.Format("02.01.2006"),
		Title:       html.EscapeString(sp.Title),
		Description: html.EscapeString(sp.Description),
		Phone:       strings.Split(sp.Phone, ";"),
		Email:       strings.Split(sp.Email, ";"),
		Cost:        sp.Cost,
		Currency:    sp.Currency,
	}
	for _, c := range sp.Categories {
		sparePart.Categories = append(sparePart.Categories, c.Name)
	}
	for _, tt := range sp.TransportType {
		sparePart.TransportTypes = append(sparePart.TransportTypes, tt.Name)
	}

	for _, img := range strings.Split(sp.Images, ";") {
		sparePart.Images = append(sparePart.Images, img)
	}

	if sp.Owner != nil {
		owner := api.FindSparePartResult_Owner{
			Id:   uint64(sp.Owner.ID),
			Name: sp.Owner.Name,
		}
		sparePart.Owner = &owner
	}

	if sp.City != nil {
		point := api.FindSparePartResult_ShortGeo{
			Id:      uint64(sp.City.ID),
			Name:    sp.City.Name,
			Type:    0,
			Address: html.EscapeString(sp.Address),
			Lat:     sp.Lat,
			Lon:     sp.Lon,
		}
		sparePart.Location = &point
	} else if sp.Region != nil {
		point := api.FindSparePartResult_ShortGeo{
			Id:      uint64(sp.Region.ID),
			Name:    html.EscapeString(sp.Region.Name),
			Type:    1,
			Address: html.EscapeString(sp.Address),
			Lat:     sp.Lat,
			Lon:     sp.Lon,
		}
		sparePart.Location = &point
	} else if sp.Country != nil {
		point := api.FindSparePartResult_ShortGeo{
			Id:      uint64(sp.Country.ID),
			Name:    html.EscapeString(sp.Country.Name),
			Type:    2,
			Address: html.EscapeString(sp.Address),
			Lat:     sp.Lat,
			Lon:     sp.Lon,
		}
		sparePart.Location = &point
	}
	return &sparePart
}

func (c *SparePartStorage) Create(sparePart *api.SparePart, userid uint64) (*SparePart, error) {
	user := NewUserStorage().GetUser(userid)
	if user == nil {
		log.Printf("User %d not found", userid)
		return nil, nil
	}
	newSparePart := SparePart{
		Title:       html.EscapeString(sparePart.Title),
		Description: html.EscapeString(sparePart.Description),
		OwnerID:     user.ID,
		Address:     html.EscapeString(sparePart.Address),
		Images:      strings.Join(sparePart.Images, ";"),
		Phone:       strings.Join(sparePart.Phone, ";"),
		Email:       strings.Join(sparePart.Email, ";"),
		Lat:         sparePart.Lat,
		Lon:         sparePart.Lon,
		Cost:        sparePart.Cost,
		Currency:    sparePart.Currency,
	}
	for _, category := range sparePart.Categories {
		storage := NewReferenceStorage()
		c := storage.GetServiceCategory(category)
		if c != nil {
			newSparePart.Categories = append(newSparePart.Categories, *c)
		} else {
			log.Printf("can not find category: %d", category)
		}
	}

	for _, transportType := range sparePart.TransportTypes {
		storage := NewReferenceStorage()
		t := storage.GetTransportTypeById(transportType)
		if c != nil {
			newSparePart.TransportType = append(newSparePart.TransportType, *t)
		} else {
			log.Printf("can not find category: %d", transportType)
		}
	}

	geo := NewGeoStorage()
	switch sparePart.Location.GetType() {
	case 0:
		city := geo.GetCity(uint(sparePart.Location.GetGeoId()))
		if city != nil {
			newSparePart.City = city
			newSparePart.RegionId = uint64(city.RegionID)
			newSparePart.CountryId = uint64(city.CountryID)
		} else {
			log.Printf("City %d not found", sparePart.Location.GetGeoId())
		}
		break
	case 1:
		region := geo.GetRegion(uint(sparePart.Location.GetGeoId()))
		if region != nil {
			newSparePart.Region = region
		} else {
			log.Printf("Region %d not found", sparePart.Location.GetGeoId())
		}
		break
	case 2:
		country := geo.GetCountry(uint(sparePart.Location.GetGeoId()))
		if country != nil {
			newSparePart.Country = country
		} else {
			log.Printf("Country %d not found", sparePart.Location.GetGeoId())
		}
		break
	}
	log.Printf("newSparePart: %v", newSparePart)
	result := db.Create(&newSparePart)
	if result.Error != nil {
		log.Printf("Can not ctreate SparePart %v", result.Error.Error())
		return nil, result.Error
	}
	return &newSparePart, nil
}

func (c *SparePartStorage) Update(sparePart *api.SparePart, userid uint64) (*SparePart, error) {
	var oldSparePart *SparePart
	result := c.db.First(&oldSparePart, sparePart.GetId())
	if result.Error != nil {
		log.Printf("SparePart %d not found", sparePart.GetId())
		return nil, result.Error
	}
	user := NewUserStorage().GetUser(userid)
	if user == nil {
		log.Printf("User %d not found", userid)
		return nil, nil
	}
	newSparePart := &SparePart{
		Model:       gorm.Model{ID: oldSparePart.ID, CreatedAt: oldSparePart.CreatedAt},
		Title:       html.EscapeString(sparePart.Title),
		Description: html.EscapeString(sparePart.Description),
		OwnerID:     user.ID,
		Address:     html.EscapeString(sparePart.Address),
		Images:      strings.Join(sparePart.Images, ";"),
		Phone:       strings.Join(sparePart.Phone, ";"),
		Email:       strings.Join(sparePart.Email, ";"),
		Lat:         sparePart.Lat,
		Lon:         sparePart.Lon,
		Cost:        sparePart.Cost,
		Currency:    sparePart.Currency,
	}
	_ = db.Model(&newSparePart).Association("Categories").Clear()
	for _, category := range sparePart.Categories {
		storage := NewReferenceStorage()
		c := storage.GetServiceCategory(category)
		if c != nil {
			newSparePart.Categories = append(newSparePart.Categories, *c)
		} else {
			log.Printf("can not find category: %d", category)
		}
	}
	geo := NewGeoStorage()
	switch sparePart.Location.GetType() {
	case 0:
		city := geo.GetCity(uint(sparePart.Location.GetGeoId()))
		if city != nil {
			newSparePart.City = city
			newSparePart.RegionId = uint64(city.RegionID)
			newSparePart.CountryId = uint64(city.CountryID)

		} else {
			log.Printf("City %d not found", sparePart.Location.GetGeoId())
		}
		break
	case 1:
		region := geo.GetRegion(uint(sparePart.Location.GetGeoId()))
		if region != nil {
			newSparePart.Region = region
		} else {
			log.Printf("Region %d not found", sparePart.Location.GetGeoId())
		}
		break
	case 2:
		country := geo.GetCountry(uint(sparePart.Location.GetGeoId()))
		if country != nil {
			newSparePart.Country = country
		} else {
			log.Printf("Country %d not found", sparePart.Location.GetGeoId())
		}
		break
	}
	log.Printf("newSparePart: %v", newSparePart)
	result = db.Save(&newSparePart)
	if result.Error != nil {
		log.Printf("Can not ctreate SparePart %v", result.Error.Error())
		return nil, result.Error
	}
	return newSparePart, nil
}

func (c *SparePartStorage) Delete(id uint64) uint64 {
	result := c.db.Delete(&SparePart{}, id)
	if result.Error != nil {
		return 0
	}
	return id
}

func (c *SparePartStorage) Get(id uint64) *SparePart {
	var sparePart SparePart
	result := db.Preload(clause.Associations).First(&sparePart, id)
	if result.Error != nil {
		return nil
	}
	return &sparePart
}

func (c *SparePartStorage) Find(
	categories []string,
	transportType []string,
	location []*api.Geo,
	limit uint32,
	offset uint32,
) ([]SparePart, int64) {
	var cty, reg, cou []uint64
	for _, geo := range location {
		switch geo.GetType() {
		case api.Geo_city:
			var city City
			c.db.First(&city, geo.GetGeoId())
			cty = append(cty, uint64(city.ID))
			break
		case api.Geo_region:
			var region Region
			c.db.First(&region, geo.GetGeoId())
			reg = append(reg, uint64(region.ID))
			break
		case api.Geo_country:
			var country Country
			c.db.First(&country, geo.GetGeoId())
			cou = append(cou, uint64(country.ID))
			break
		}
	}

	tx := c.db.Joins("left join parts_transport_type on spare_parts.id = parts_transport_type.spare_part_id").
		Joins("left join parts_category on spare_parts.id = parts_category.spare_part_id")
	var tt []uint
	/*rs := NewReferenceStorage()
	for _, c := range categories {
		sc := rs.GetServiceCategoryByName(c)
		cat = append(cat, sc.ID)
	}*/
	if len(categories) > 0 {
		tx.Where("parts_category.service_category_id in ?", categories)
	}
	for _, t := range transportType {
		sc := rs.GetTransportType(t, 0)
		tt = append(tt, sc.ID)
	}

	if len(tt) > 0 {
		tx.Where("parts_transport_type.transport_type_id in ?", tt)
	}

	if len(cty) > 0 {
		tx.Where("city_id in ?", cty)
	}
	if len(reg) > 0 {
		tx.Where("region_id in ?", reg)
	}
	if len(cou) > 0 {
		tx.Where("country_id in ?", cou)
	}

	var count int64
	tx.Model(SparePart{}).Distinct("id").Count(&count)
	if limit != 0 {
		tx.Limit(int(limit))
	}
	tx.Offset(int(offset))
	var spareParts []SparePart
	tx.Preload(clause.Associations).Distinct("spare_parts.*").Order("created_at desc").Find(&spareParts)
	return spareParts, count
}
