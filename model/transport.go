package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"html"
	"log"
	"strconv"
	"strings"
	"tsl_server/api"
)

type Transport struct {
	gorm.Model
	SalesType        api.SalesType `gorm:"type:INT"`
	BrandId          uint64
	Brand            *Brand
	TransportTypeId  uint64
	TransportType    *TransportType
	ModelName        string
	ReleaseYEar      string
	Engine           uint32
	Transmission     string
	Mileage          uint64
	FuelType         string
	Availability     string
	Condition        string
	Cost             uint64
	Information      string
	CityId           uint64
	City             *City
	RegionId         uint64
	Region           *Region
	CountryId        uint64
	Country          *Country
	OwnerID          uint
	Owner            *TslUser `gorm:"foreignKey:OwnerID"`
	Images           string
	Currency         string
	Phones           string
	Emails           string
	NumberPassengers uint32
	Powder           uint32
}

type TransportStorage struct {
	db *gorm.DB
}

var t *TransportStorage

func NewTransportStorage() *TransportStorage {
	if t == nil {
		t = &TransportStorage{
			db: GetDB(),
		}
	}
	return t
}

func (cs *Transport) ToAPI() *api.FindTransportResult {

	transport := api.FindTransportResult{
		Id:               uint64(cs.ID),
		CreatedAt:        cs.CreatedAt.Format("02.01.2006"),
		UpdatedAt:        cs.UpdatedAt.Format("02.01.2006"),
		Brand:            html.EscapeString(cs.Brand.Name),
		TransportType:    html.EscapeString(cs.TransportType.Name),
		Transmission:     html.EscapeString(cs.Transmission),
		Model:            html.EscapeString(cs.ModelName),
		ReleaseYear:      cs.ReleaseYEar,
		Engine:           strconv.Itoa(int(cs.Engine)),
		Mileage:          cs.Mileage,
		FuelType:         cs.FuelType,
		Availability:     html.EscapeString(cs.Availability),
		Condition:        html.EscapeString(cs.Condition),
		Cost:             cs.Cost,
		Information:      html.EscapeString(cs.Information),
		Images:           strings.Split(cs.Images, ";"),
		Phones:           strings.Split(cs.Phones, ";"),
		Emails:           strings.Split(cs.Emails, ";"),
		NumberPassengers: cs.NumberPassengers,
		Powder:           cs.Powder,
		Currency:         cs.Currency,
	}
	if cs.Owner != nil {
		owner := api.FindTransportResult_Owner{
			Id:   uint64(cs.Owner.ID),
			Name: cs.Owner.Name,
		}
		transport.Owner = &owner
	}

	if cs.City != nil {
		point := api.FindTransportResult_ShortGeo{
			Id:   uint64(cs.City.ID),
			Name: cs.City.Name,
			Type: 0,
		}
		transport.Location = &point
	} else if cs.Region != nil {
		point := api.FindTransportResult_ShortGeo{
			Id:   uint64(cs.Region.ID),
			Name: cs.Region.Name,
			Type: 1,
		}
		transport.Location = &point
	} else if cs.Country != nil {
		point := api.FindTransportResult_ShortGeo{
			Id:   uint64(cs.Country.ID),
			Name: cs.Country.Name,
			Type: 2,
		}
		transport.Location = &point
	}
	switch cs.SalesType {
	case api.SalesType_sale:
		transport.SalesType = "Продажа"
		break
	case api.SalesType_rent:
		transport.SalesType = "Аренда"
		break
	case api.SalesType_leasing:
		transport.SalesType = "Лизинг"
		break
	}
	return &transport
}

func (c *TransportStorage) Create(transport *api.Transport, userid uint64) (*Transport, error) {
	user := NewUserStorage().GetUser(userid)
	if user == nil {
		log.Printf("User %d not found", userid)
		return nil, nil
	}
	newTransport := Transport{
		SalesType:        transport.GetType(),
		BrandId:          transport.GetBrand(),
		TransportTypeId:  transport.GetTransportType(),
		ModelName:        html.EscapeString(transport.GetModel()),
		ReleaseYEar:      transport.GetReleaseYear(),
		Engine:           transport.GetEngine(),
		Transmission:     transport.GetTransmission(),
		Mileage:          transport.GetMileage(),
		FuelType:         transport.GetFuelType(),
		Availability:     transport.GetAvailability(),
		Condition:        transport.GetCondition(),
		Cost:             transport.GetCost(),
		Information:      html.EscapeString(transport.GetInformation()),
		CityId:           0,
		City:             nil,
		RegionId:         0,
		Region:           nil,
		CountryId:        0,
		Country:          nil,
		OwnerID:          user.ID,
		Images:           strings.Join(transport.GetImages(), ";"),
		Phones:           strings.Join(transport.GetPhones(), ";"),
		Emails:           strings.Join(transport.GetEmails(), ";"),
		Currency:         transport.GetCurrency(),
		NumberPassengers: transport.GetNumberPassengers(),
		Powder:           transport.GetPowder(),
	}
	geo := NewGeoStorage()
	switch transport.Placement.GetType() {
	case 0:
		city := geo.GetCity(uint(transport.Placement.GetGeoId()))
		if city != nil {
			newTransport.City = city
			newTransport.RegionId = uint64(city.RegionID)
			newTransport.CountryId = uint64(city.CountryID)
		} else {
			log.Printf("City %d not found", transport.Placement.GetGeoId())
		}
		break
	case 1:
		region := geo.GetRegion(uint(transport.Placement.GetGeoId()))
		if region != nil {
			newTransport.Region = region
		} else {
			log.Printf("Region %d not found", transport.Placement.GetGeoId())
		}
		break
	case 2:
		country := geo.GetCountry(uint(transport.Placement.GetGeoId()))
		if country != nil {
			newTransport.Country = country
		} else {
			log.Printf("Country %d not found", transport.Placement.GetGeoId())
		}
		break
	}
	log.Printf("newTransport: %v", newTransport)
	result := db.Create(&newTransport)
	if result.Error != nil {
		log.Printf("Can not ctreate Transport %v", result.Error.Error())
		return nil, result.Error
	}
	return &newTransport, nil
}

func (c *TransportStorage) Update(transport *api.Transport, userid uint64) (*Transport, error) {
	var oldTransport *Transport
	result := c.db.First(&oldTransport, transport.GetId())
	if result.Error != nil {
		log.Printf("Transport %d not found", transport.GetId())
		return nil, result.Error
	}
	user := NewUserStorage().GetUser(userid)
	if user == nil {
		log.Printf("User %d not found", userid)
		return nil, nil
	}
	newTransport := &Transport{
		Model:            gorm.Model{ID: oldTransport.ID, CreatedAt: oldTransport.CreatedAt},
		SalesType:        transport.GetType(),
		BrandId:          transport.GetBrand(),
		TransportTypeId:  transport.GetTransportType(),
		ModelName:        html.EscapeString(transport.GetModel()),
		ReleaseYEar:      transport.GetReleaseYear(),
		Engine:           transport.GetEngine(),
		Transmission:     transport.GetTransmission(),
		Mileage:          transport.GetMileage(),
		FuelType:         transport.GetFuelType(),
		Availability:     transport.GetAvailability(),
		Condition:        transport.GetCondition(),
		Cost:             transport.GetCost(),
		Information:      html.EscapeString(transport.GetInformation()),
		CityId:           0,
		City:             nil,
		RegionId:         0,
		Region:           nil,
		CountryId:        0,
		Country:          nil,
		OwnerID:          user.ID,
		Images:           strings.Join(transport.GetImages(), ";"),
		Phones:           strings.Join(transport.GetPhones(), ";"),
		Emails:           strings.Join(transport.GetEmails(), ";"),
		Currency:         transport.GetCurrency(),
		NumberPassengers: transport.GetNumberPassengers(),
		Powder:           transport.GetPowder(),
	}
	geo := NewGeoStorage()
	switch transport.Placement.GetType() {
	case 0:
		city := geo.GetCity(uint(transport.Placement.GetGeoId()))
		if city != nil {
			newTransport.City = city
			newTransport.RegionId = uint64(city.RegionID)
			newTransport.CountryId = uint64(city.CountryID)

		} else {
			log.Printf("City %d not found", transport.Placement.GetGeoId())
		}
		break
	case 1:
		region := geo.GetRegion(uint(transport.Placement.GetGeoId()))
		if region != nil {
			newTransport.Region = region
		} else {
			log.Printf("Region %d not found", transport.Placement.GetGeoId())
		}
		break
	case 2:
		country := geo.GetCountry(uint(transport.Placement.GetGeoId()))
		if country != nil {
			newTransport.Country = country
		} else {
			log.Printf("Country %d not found", transport.Placement.GetGeoId())
		}
		break
	}
	result = db.Save(&newTransport)
	if result.Error != nil {
		log.Printf("Can not ctreate Transport %v", result.Error.Error())
		return nil, result.Error
	}
	return newTransport, nil
}

func (c *TransportStorage) Delete(id uint64) uint64 {
	result := c.db.Delete(&Transport{}, id)
	if result.Error != nil {
		return 0
	}
	return id
}

func (c *TransportStorage) Get(id uint64) *Transport {
	var transport Transport
	result := db.Preload(clause.Associations).First(&transport, id)
	if result.Error != nil {
		return nil
	}
	return &transport
}

func (c *TransportStorage) Find(
	salesType api.SalesType,
	placement *api.Geo,
	conditions string,
	transportType []uint64,
	fuelType []string,
	brand []uint64,
	costFrom uint64,
	costTill uint64,
	limit uint32,
	offset uint32,
) ([]Transport, int64) {
	tx := c.db.Limit(10) //.Session(&gorm.Session{QueryFields: true})
	if salesType != 0 {
		tx.Where("sales_type = ?", salesType.Number())
	}
	if len(transportType) > 0 {
		tx.Where("transport_type_id in ?", transportType)
	}
	if placement != nil {
		switch placement.GetType() {
		case api.Geo_city:
			tx.Where("city_id = ?", placement.GetGeoId())
			break
		case api.Geo_region:
			tx.Where("region_id = ?", placement.GetGeoId())
			break
		case api.Geo_country:
			tx.Where("country_id = ?", placement.GetGeoId())
			break
		}
	}
	if conditions != "" {
		tx.Where("transports.`condition` = ?", conditions)
	}
	if len(fuelType) > 0 {
		tx.Where("fuel_type in  ?", fuelType)
	}
	if len(brand) > 0 {
		tx.Where("brand_id = ?", brand)
	}

	if costFrom != 0 {
		tx.Where("cost >= ?", costFrom)
	}
	if costTill != 0 {
		tx.Where("cost <= ?", costTill)
	}
	var count int64
	tx.Model(Transport{}).Distinct("id").Count(&count)
	log.Printf("limit: %d, offset %d", limit, offset)
	if limit > 0 {
		tx.Limit(int(limit))
		tx.Offset(int(offset))
	}

	var transports []Transport
	tx.Preload(clause.Associations).Distinct("transports.*").Order("created_at desc").Find(&transports)
	return transports, count
}
