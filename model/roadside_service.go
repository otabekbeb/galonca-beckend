package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"html"
	"log"
	"strings"
	"tsl_server/api"
)

type RoadsideService struct {
	gorm.Model
	Type        uint32
	Title       string
	Description string `gorm:"type:text"`
	CityId      uint64
	City        *City
	RegionId    uint64
	Region      *Region
	CountryId   uint64
	Country     *Country
	Address     string
	Images      string
	Phone       string
	Email       string
	Keys        string
	Lat         float64
	Lon         float64
	OwnerID     uint
	Owner       *TslUser `gorm:"foreignKey:OwnerID"`
}

type RoadsideServiceStorage struct {
	db *gorm.DB
}

var rss *RoadsideServiceStorage

func NewRoadsideServiceStorage() *RoadsideServiceStorage {
	if rss == nil {
		rss = &RoadsideServiceStorage{
			db: GetDB(),
		}
	}
	return rss
}

func (cs *RoadsideService) ToAPI() *api.FindRoadsideServiceResult {

	roadsideService := api.FindRoadsideServiceResult{
		Id:          uint64(cs.ID),
		CreatedAt:   cs.CreatedAt.Format("02.01.2006"),
		UpdatedAt:   cs.UpdatedAt.Format("02.01.2006"),
		Type:        cs.Type,
		Title:       html.EscapeString(cs.Title),
		Description: html.EscapeString(cs.Description),
		Phone:       strings.Split(cs.Phone, ";"),
		Email:       strings.Split(cs.Email, ";"),
		Images:      strings.Split(cs.Images, ";"),
		Keys:        strings.Split(cs.Keys, ";"),
	}
	if cs.Owner != nil {
		owner := api.FindRoadsideServiceResult_Owner{
			Id:   uint64(cs.Owner.ID),
			Name: cs.Owner.Name,
		}
		roadsideService.Owner = &owner
	}

	if cs.City != nil {
		point := api.FindRoadsideServiceResult_ShortGeo{
			Id:      uint64(cs.City.ID),
			Name:    cs.City.Name,
			Type:    0,
			Address: html.EscapeString(cs.Address),
			Lat:     cs.Lat,
			Lon:     cs.Lon,
		}
		roadsideService.Location = &point
	} else if cs.Region != nil {
		point := api.FindRoadsideServiceResult_ShortGeo{
			Id:      uint64(cs.Region.ID),
			Name:    cs.Region.Name,
			Type:    1,
			Address: html.EscapeString(cs.Address),
			Lat:     cs.Lat,
			Lon:     cs.Lon,
		}
		roadsideService.Location = &point
	} else if cs.Country != nil {
		point := api.FindRoadsideServiceResult_ShortGeo{
			Id:      uint64(cs.Country.ID),
			Name:    cs.Country.Name,
			Type:    2,
			Address: html.EscapeString(cs.Address),
			Lat:     cs.Lat,
			Lon:     cs.Lon,
		}
		roadsideService.Location = &point
	}
	return &roadsideService
}

func (c *RoadsideServiceStorage) Create(roadsideService *api.RoadsideService, userid uint64) (*RoadsideService, error) {
	user := NewUserStorage().GetUser(userid)
	if user == nil {
		log.Printf("User %d not found", userid)
		return nil, nil
	}
	newRoadsideService := RoadsideService{
		Type:        uint32(roadsideService.Type),
		Title:       html.EscapeString(roadsideService.Title),
		Description: html.EscapeString(roadsideService.Description),
		OwnerID:     user.ID,
		Address:     html.EscapeString(roadsideService.Address),
		Images:      strings.Join(roadsideService.Images, ";"),
		Phone:       strings.Join(roadsideService.Phone, ";"),
		Email:       strings.Join(roadsideService.Email, ";"),
		Lat:         roadsideService.Lat,
		Lon:         roadsideService.Lon,
	}
	geo := NewGeoStorage()
	switch roadsideService.Location.GetType() {
	case 0:
		city := geo.GetCity(uint(roadsideService.Location.GetGeoId()))
		if city != nil {
			newRoadsideService.City = city
			newRoadsideService.RegionId = uint64(city.RegionID)
			newRoadsideService.CountryId = uint64(city.CountryID)
		} else {
			log.Printf("City %d not found", roadsideService.Location.GetGeoId())
		}
		break
	case 1:
		region := geo.GetRegion(uint(roadsideService.Location.GetGeoId()))
		if region != nil {
			newRoadsideService.Region = region
		} else {
			log.Printf("Region %d not found", roadsideService.Location.GetGeoId())
		}
		break
	case 2:
		country := geo.GetCountry(uint(roadsideService.Location.GetGeoId()))
		if country != nil {
			newRoadsideService.Country = country
		} else {
			log.Printf("Country %d not found", roadsideService.Location.GetGeoId())
		}
		break
	}
	log.Printf("newRoadsideService: %v", newRoadsideService)
	result := db.Create(&newRoadsideService)
	if result.Error != nil {
		log.Printf("Can not ctreate RoadsideService %v", result.Error.Error())
		return nil, result.Error
	}
	return &newRoadsideService, nil
}

func (c *RoadsideServiceStorage) Update(roadsideService *api.RoadsideService, userid uint64) (*RoadsideService, error) {
	var oldRoadsideService *RoadsideService
	result := c.db.First(&oldRoadsideService, roadsideService.GetId())
	if result.Error != nil {
		log.Printf("RoadsideService %d not found", roadsideService.GetId())
		return nil, result.Error
	}
	user := NewUserStorage().GetUser(userid)
	if user == nil {
		log.Printf("User %d not found", userid)
		return nil, nil
	}
	newRoadsideService := &RoadsideService{
		Model:       gorm.Model{ID: oldRoadsideService.ID, CreatedAt: oldRoadsideService.CreatedAt},
		Type:        uint32(roadsideService.Type),
		Title:       html.EscapeString(roadsideService.Title),
		Description: html.EscapeString(roadsideService.Description),
		OwnerID:     user.ID,
		Address:     html.EscapeString(roadsideService.Address),
		Images:      strings.Join(roadsideService.Images, ";"),
		Phone:       strings.Join(roadsideService.Phone, ";"),
		Email:       strings.Join(roadsideService.Email, ";"),
		Lat:         roadsideService.Lat,
		Lon:         roadsideService.Lon,
	}
	geo := NewGeoStorage()
	switch roadsideService.Location.GetType() {
	case 0:
		city := geo.GetCity(uint(roadsideService.Location.GetGeoId()))
		if city != nil {
			newRoadsideService.City = city
			newRoadsideService.RegionId = uint64(city.RegionID)
			newRoadsideService.CountryId = uint64(city.CountryID)
		} else {
			log.Printf("City %d not found", roadsideService.Location.GetGeoId())
		}
		break
	case 1:
		region := geo.GetRegion(uint(roadsideService.Location.GetGeoId()))
		if region != nil {
			newRoadsideService.Region = region
		} else {
			log.Printf("Region %d not found", roadsideService.Location.GetGeoId())
		}
		break
	case 2:
		country := geo.GetCountry(uint(roadsideService.Location.GetGeoId()))
		if country != nil {
			newRoadsideService.Country = country
		} else {
			log.Printf("Country %d not found", roadsideService.Location.GetGeoId())
		}
		break
	}
	log.Printf("newRoadsideService: %v", newRoadsideService)
	result = db.Save(&newRoadsideService)
	if result.Error != nil {
		log.Printf("Can not ctreate RoadsideService %v", result.Error.Error())
		return nil, result.Error
	}
	return newRoadsideService, nil
}

func (c *RoadsideServiceStorage) Delete(id uint64) uint64 {
	result := c.db.Delete(&RoadsideService{}, id)
	if result.Error != nil {
		return 0
	}
	return id
}

func (c *RoadsideServiceStorage) Get(id uint64) *RoadsideService {
	var roadsideService RoadsideService
	result := db.Preload(clause.Associations).First(&roadsideService, id)
	if result.Error != nil {
		return nil
	}
	return &roadsideService
}

func (c *RoadsideServiceStorage) Find(
	serviceType int32,
	location *api.Geo,
	limit uint32,
	offset uint32,
) ([]RoadsideService, int64) {
	var cty, reg, cou []uint64
	if location != nil {
		switch location.GetType() {
		case api.Geo_city:
			var city City
			c.db.First(&city, location.GetGeoId())
			cty = append(cty, uint64(city.ID))
			break
		case api.Geo_region:
			var region Region
			c.db.First(&region, location.GetGeoId())
			reg = append(reg, uint64(region.ID))
			break
		case api.Geo_country:
			var country Country
			c.db.First(&country, location.GetGeoId())
			cou = append(cou, uint64(country.ID))
			break
		}
	}
	tx := c.db.Limit(100) //.Session(&gorm.Session{QueryFields: true})
	if serviceType != 10 {
		tx.Where("type = ?", serviceType)
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
	if limit != 0 {
		tx.Limit(int(limit))
	}
	tx.Offset(int(offset))
	var count int64
	tx.Model(RoadsideService{}).Distinct("id").Count(&count)
	var roadsideServices []RoadsideService
	tx.Preload(clause.Associations).Distinct("roadside_services.*").Order("created_at desc").Find(&roadsideServices)
	return roadsideServices, count
}
