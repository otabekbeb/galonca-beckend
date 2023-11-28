package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"html"
	"log"
	"strings"
	"tsl_server/api"
)

type ServiceStation struct {
	gorm.Model
	Title       string
	Description string            `gorm:"type:text"`
	Categories  []ServiceCategory `gorm:"many2many:service_station_category"`
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
	Lat         float64
	Lon         float64
	OwnerID     uint
	Owner       *TslUser `gorm:"foreignKey:OwnerID"`
}

type ServiceStationStorage struct {
	db *gorm.DB
}

var sss *ServiceStationStorage

func NewServiceStationStorage() *ServiceStationStorage {
	if sss == nil {
		sss = &ServiceStationStorage{
			db: GetDB(),
		}
	}
	return sss
}

func (cs *ServiceStation) ToAPI() *api.FindServiceStationResult {

	serviceStation := api.FindServiceStationResult{
		Id:          uint64(cs.ID),
		CreatedAt:   cs.CreatedAt.Format("02.01.2006"),
		UpdatedAt:   cs.UpdatedAt.Format("02.01.2006"),
		Title:       html.EscapeString(cs.Title),
		Description: html.EscapeString(cs.Description),
		Phone:       strings.Split(cs.Phone, ";"),
		Email:       strings.Split(cs.Email, ";"),
	}
	if cs.Owner != nil {
		owner := api.FindServiceStationResult_Owner{
			Id:   uint64(cs.Owner.ID),
			Name: cs.Owner.Name,
		}
		serviceStation.Owner = &owner
	}

	for _, img := range strings.Split(cs.Images, ";") {
		serviceStation.Images = append(serviceStation.Images, img)
	}

	for _, cat := range cs.Categories {
		serviceStation.Categories = append(serviceStation.Categories, cat.Name)
	}

	if cs.City != nil {
		point := api.FindServiceStationResult_ShortGeo{
			Id:      uint64(cs.City.ID),
			Name:    cs.City.Name,
			Type:    0,
			Address: html.EscapeString(cs.Address),
			Lat:     cs.Lat,
			Lon:     cs.Lon,
		}
		serviceStation.Location = &point
	} else if cs.Region != nil {
		point := api.FindServiceStationResult_ShortGeo{
			Id:      uint64(cs.Region.ID),
			Name:    cs.Region.Name,
			Type:    1,
			Address: html.EscapeString(cs.Address),
			Lat:     cs.Lat,
			Lon:     cs.Lon,
		}
		serviceStation.Location = &point
	} else if cs.Country != nil {
		point := api.FindServiceStationResult_ShortGeo{
			Id:      uint64(cs.Country.ID),
			Name:    cs.Country.Name,
			Type:    2,
			Address: html.EscapeString(cs.Address),
			Lat:     cs.Lat,
			Lon:     cs.Lon,
		}
		serviceStation.Location = &point
	}
	return &serviceStation
}

func (c *ServiceStationStorage) Create(serviceStation *api.ServiceStation, userid uint64) (*ServiceStation, error) {
	user := NewUserStorage().GetUser(userid)
	if user == nil {
		log.Printf("User %d not found", userid)
		return nil, nil
	}
	newServiceStation := ServiceStation{
		Title:       html.EscapeString(serviceStation.Title),
		Description: html.EscapeString(serviceStation.Description),
		OwnerID:     user.ID,
		Address:     html.EscapeString(serviceStation.Address),
		Images:      strings.Join(serviceStation.Images, ";"),
		Phone:       strings.Join(serviceStation.Phone, ";"),
		Email:       strings.Join(serviceStation.Email, ";"),
		Lat:         serviceStation.Lat,
		Lon:         serviceStation.Lon,
	}
	for _, category := range serviceStation.Categories {
		storage := NewReferenceStorage()
		c := storage.GetServiceCategory(category)
		if c != nil {
			newServiceStation.Categories = append(newServiceStation.Categories, *c)
		} else {
			log.Printf("can not find category: %d", category)
		}
	}
	geo := NewGeoStorage()
	switch serviceStation.Location.GetType() {
	case 0:
		city := geo.GetCity(uint(serviceStation.Location.GetGeoId()))
		if city != nil {
			newServiceStation.City = city
			newServiceStation.RegionId = uint64(city.RegionID)
			newServiceStation.CountryId = uint64(city.CountryID)
		} else {
			log.Printf("City %d not found", serviceStation.Location.GetGeoId())
		}
		break
	case 1:
		region := geo.GetRegion(uint(serviceStation.Location.GetGeoId()))
		if region != nil {
			newServiceStation.Region = region
		} else {
			log.Printf("Region %d not found", serviceStation.Location.GetGeoId())
		}
		break
	case 2:
		country := geo.GetCountry(uint(serviceStation.Location.GetGeoId()))
		if country != nil {
			newServiceStation.Country = country
		} else {
			log.Printf("Country %d not found", serviceStation.Location.GetGeoId())
		}
		break
	}
	log.Printf("newServiceStation: %v", newServiceStation)
	result := db.Create(&newServiceStation)
	if result.Error != nil {
		log.Printf("Can not ctreate ServiceStation %v", result.Error.Error())
		return nil, result.Error
	}
	return &newServiceStation, nil
}

func (c *ServiceStationStorage) Update(serviceStation *api.ServiceStation, userid uint64) (*ServiceStation, error) {
	var oldServiceStation *ServiceStation
	result := c.db.First(&oldServiceStation, serviceStation.GetId())
	if result.Error != nil {
		log.Printf("ServiceStation %d not found", serviceStation.GetId())
		return nil, result.Error
	}
	user := NewUserStorage().GetUser(userid)
	if user == nil {
		log.Printf("User %d not found", userid)
		return nil, nil
	}
	newServiceStation := &ServiceStation{
		Model: gorm.Model{
			ID:        oldServiceStation.ID,
			CreatedAt: oldServiceStation.CreatedAt,
		},
		Title:       html.EscapeString(serviceStation.Title),
		Description: html.EscapeString(serviceStation.Description),
		OwnerID:     user.ID,
		Address:     html.EscapeString(serviceStation.Address),
		Images:      strings.Join(serviceStation.Images, ";"),
		Phone:       strings.Join(serviceStation.Phone, ";"),
		Email:       strings.Join(serviceStation.Email, ";"),
		Lat:         serviceStation.Lat,
		Lon:         serviceStation.Lon,
	}
	_ = db.Model(&newServiceStation).Association("Categories").Clear()
	for _, category := range serviceStation.Categories {
		storage := NewReferenceStorage()
		c := storage.GetServiceCategory(category)
		if c != nil {
			newServiceStation.Categories = append(newServiceStation.Categories, *c)
		} else {
			log.Printf("can not find category: %d", category)
		}
	}
	geo := NewGeoStorage()
	switch serviceStation.Location.GetType() {
	case 0:
		city := geo.GetCity(uint(serviceStation.Location.GetGeoId()))
		if city != nil {
			newServiceStation.City = city
			newServiceStation.RegionId = uint64(city.RegionID)
			newServiceStation.CountryId = uint64(city.CountryID)
		} else {
			log.Printf("City %d not found", serviceStation.Location.GetGeoId())
		}
		break
	case 1:
		region := geo.GetRegion(uint(serviceStation.Location.GetGeoId()))
		if region != nil {
			newServiceStation.Region = region
		} else {
			log.Printf("Region %d not found", serviceStation.Location.GetGeoId())
		}
		break
	case 2:
		country := geo.GetCountry(uint(serviceStation.Location.GetGeoId()))
		if country != nil {
			newServiceStation.Country = country
		} else {
			log.Printf("Country %d not found", serviceStation.Location.GetGeoId())
		}
		break
	}
	log.Printf("newServiceStation: %v", newServiceStation)
	result = db.Save(&newServiceStation)
	if result.Error != nil {
		log.Printf("Can not ctreate ServiceStation %v", result.Error.Error())
		return nil, result.Error
	}
	return newServiceStation, nil
}

func (c *ServiceStationStorage) Delete(id uint64) uint64 {
	result := c.db.Delete(&ServiceStation{}, id)
	if result.Error != nil {
		return 0
	}
	return id
}

func (c *ServiceStationStorage) Get(id uint64) *ServiceStation {
	var serviceStation ServiceStation
	result := db.Preload(clause.Associations).First(&serviceStation, id)
	if result.Error != nil {
		return nil
	}
	return &serviceStation
}

func (c *ServiceStationStorage) Find(
	categories []string,
	location []*api.Geo,
	limit uint32,
	offset uint32,
) ([]ServiceStation, int64) {
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

	tx := c.db.Joins("left join service_station_category on service_stations.id = service_station_category.service_station_id")
	if len(cty) > 0 {
		tx.Where("city_id in ?", cty)
	}
	if len(reg) > 0 {
		tx.Where("region_id in ?", reg)
	}
	if len(cou) > 0 {
		tx.Where("country_id in ?", cou)
	}
	/*var cat []uint
	rs := NewReferenceStorage()
	for _, c := range categories {
		sc := rs.GetServiceCategoryByName(c)
		cat = append(cat, sc.ID)
	}*/
	if len(categories) > 0 {
		tx.Where("service_station_category.service_category_id in ?", categories)
	}
	var count int64
	tx.Model(ServiceStation{}).Distinct("id").Count(&count)
	if limit != 0 {
		tx.Limit(int(limit))
	}
	tx.Offset(int(offset))
	var serviceStations []ServiceStation
	tx.Preload(clause.Associations).Distinct("service_stations.*").Order("created_at desc").Find(&serviceStations)
	return serviceStations, count
}
