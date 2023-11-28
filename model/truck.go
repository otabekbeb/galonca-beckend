package model

import (
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
	"tsl_server/api"
)

type Truck struct {
	gorm.Model
	LoadingPoints   []*City `gorm:"many2many:truck_loading_points"`
	UnloadingPoints []*City `gorm:"many2many:truck_unloading_points"`
	From            time.Time
	Till            time.Time
	Weight          uint
	Volume          uint
	Cost            uint
	Currency        string
	Transport       []*TransportType `gorm:"many2many:truck_transport"`
	Loading         []*LoadingType   `gorm:"many2many:truck_loading"`
	Addition        []*Addition      `gorm:"many2many:truck_addition"`
	TypeId          uint
	Type            *CargoType
	AdditionalPhone string `gorm:"type:text"`
	AdditionalEmail string `gorm:"type:text"`
	OwnerID         uint
	Owner           *TslUser `gorm:"foreignKey:OwnerID"`
}

type TruckLoadingPoints struct {
	TruckId   uint64 `gorm:"primaryKey"`
	CityId    uint64 `gorm:"primaryKey"`
	SortOrder uint   `gorm:"not null,DEFAULT 0"`
}

type TruckUnloadingPoints struct {
	TruckId   uint64 `gorm:"primaryKey"`
	CityId    uint64 `gorm:"primaryKey"`
	SortOrder uint   `gorm:"not null,DEFAULT 0"`
}

type TruckStorage struct {
	db *gorm.DB
}

var ts *TruckStorage

func NewTruckStorage() *TruckStorage {
	if ts == nil {
		ts = &TruckStorage{
			db: GetDB(),
		}
	}
	return ts
}

func (c *TruckStorage) loadingPoints(truck *Truck) error {
	var lPoints []TruckLoadingPoints
	err := c.db.Model(&TruckLoadingPoints{}).Order("sort_order asc").Find(&lPoints, "truck_id = ?", truck.ID)
	if err.Error != nil {
		return err.Error
	}
	for _, p := range lPoints {
		var city City
		err = c.db.Model(&City{}).Find(&city, p.CityId)
		if err.Error == nil {
			city.SortOrder = p.SortOrder
			truck.LoadingPoints = append(truck.LoadingPoints, &city)
		}
	}
	return nil
}

func (c *TruckStorage) unloadingPoints(truck *Truck) error {
	var lPoints []TruckUnloadingPoints
	err := c.db.Model(&TruckUnloadingPoints{}).Order("sort_order asc").Find(&lPoints, "truck_id = ?", truck.ID)
	if err.Error != nil {
		return err.Error
	}
	for _, p := range lPoints {
		var city City
		err = c.db.Model(&City{}).Find(&city, p.CityId)
		if err.Error == nil {
			city.SortOrder = p.SortOrder
			truck.UnloadingPoints = append(truck.UnloadingPoints, &city)
		}
	}
	return nil
}

func (c *TruckStorage) saveLoadingPoints(truckId uint, cities []uint64) {
	for index, cityId := range cities {
		geo := NewGeoStorage()
		city := geo.GetCity(uint(cityId))
		if city != nil {
			clp := TruckLoadingPoints{
				TruckId:   uint64(truckId),
				CityId:    cityId,
				SortOrder: uint(index),
			}
			c.db.Save(&clp)
		} else {
			log.Printf("can not find city: %d", cityId)
		}
	}
}

func (c *TruckStorage) saveUnloadingPoints(truckId uint, cities []uint64) {
	for index, cityId := range cities {
		geo := NewGeoStorage()
		city := geo.GetCity(uint(cityId))
		if city != nil {
			clp := TruckUnloadingPoints{
				TruckId:   uint64(truckId),
				CityId:    cityId,
				SortOrder: uint(index),
			}
			c.db.Save(&clp)
		} else {
			log.Printf("can not find city: %d", cityId)
		}
	}
}

func (cs *Truck) ToAPI() *api.FindTruckResult {

	truck := api.FindTruckResult{
		Id:        uint64(cs.ID),
		CreatedAt: cs.CreatedAt.Format("02.01.2006"),
		UpdatedAt: cs.UpdatedAt.Format("02.01.2006"),
		From:      cs.From.Format("02.01.2006"),
		Till:      cs.Till.Format("02.01.2006"),
		Weight:    uint32(cs.Weight),
		Volume:    uint32(cs.Volume),
		Cost:      uint32(cs.Cost),
		Currency:  cs.Currency,
	}
	if cs.Type != nil {
		truck.CargoType = cs.Type.Name
	}
	if cs.Owner != nil {
		owner := api.FindTruckResult_Owner{
			Id:   uint64(cs.Owner.ID),
			Name: cs.Owner.Name,
		}
		truck.Owner = &owner
	}

	for _, lp := range cs.LoadingPoints {
		point := api.FindTruckResult_ShortGeo{
			Id:   uint64(lp.ID),
			Name: lp.Name,
			Type: 0,
		}
		truck.LoadingPoints = append(truck.LoadingPoints, &point)
	}
	for _, up := range cs.UnloadingPoints {
		point := api.FindTruckResult_ShortGeo{
			Id:   uint64(up.ID),
			Name: up.Name,
			Type: 0,
		}
		truck.UnloadingPoints = append(truck.UnloadingPoints, &point)
	}
	for _, tt := range cs.Transport {
		truck.TransportType = append(truck.TransportType, tt.Name)
	}
	for _, tt := range cs.Loading {
		truck.LoadingType = append(truck.LoadingType, tt.Name)
	}
	for _, tt := range cs.Addition {
		truck.Addition = append(truck.Addition, tt.Name)
	}
	truck.AdditionalPhones = strings.Split(cs.AdditionalPhone, ";")
	truck.AdditionalMails = strings.Split(cs.AdditionalEmail, ";")
	return &truck
}

func (c *TruckStorage) Create(truck *api.Truck, userid uint64) (*Truck, error) {
	user := NewUserStorage().GetUser(userid)
	if user == nil {
		log.Printf("User %d not found", userid)
		return nil, nil
	}
	newTruck := Truck{
		From:            truck.From.AsTime(),
		Till:            truck.Till.AsTime(),
		Weight:          uint(truck.Weight),
		Volume:          uint(truck.Volume),
		Cost:            uint(truck.Cost),
		Currency:        truck.Currency,
		AdditionalPhone: strings.Join(truck.AdditionalPhones, ";"),
		AdditionalEmail: strings.Join(truck.AdditionalMails, ";"),
		OwnerID:         user.ID,
	}
	//for _, cityId := range truck.LoadingPoints {
	//	geo := NewGeoStorage()
	//	city := geo.GetCity(uint(cityId))
	//	if city != nil {
	//		newTruck.LoadingPoints = append(newTruck.LoadingPoints, city)
	//	} else {
	//		log.Printf("can not find city: %d", cityId)
	//	}
	//}
	//for _, cityId := range truck.UnloadingPoints {
	//	geo := NewGeoStorage()
	//	city := geo.GetCity(uint(cityId))
	//	if city != nil {
	//		newTruck.UnloadingPoints = append(newTruck.UnloadingPoints, city)
	//	} else {
	//		log.Printf("can not find city: %d", cityId)
	//	}
	//}
	for _, transport := range truck.TransportType {
		storage := NewReferenceStorage()
		transportType := storage.GetTransportType(transport, 0)
		if transportType != nil {
			newTruck.Transport = append(newTruck.Transport, transportType)
		} else {
			log.Printf("can not find transport: %s", transport)
		}
	}

	for _, loading := range truck.LoadingType {
		storage := NewReferenceStorage()
		loadingType := storage.GetLoadingType(loading)
		if loadingType != nil {
			newTruck.Loading = append(newTruck.Loading, loadingType)
		} else {
			log.Printf("can not find transport: %s", loading)
		}
	}
	if truck.CargoType != 0 {
		storage := NewReferenceStorage()
		cargoType := storage.GetCargoTypeById(truck.CargoType)
		if cargoType != nil {
			newTruck.Type = cargoType
			newTruck.TypeId = cargoType.ID
		} else {
			log.Printf("can not find cargotype: %d", truck.CargoType)
		}
	}

	for _, add := range truck.Addition {
		storage := NewReferenceStorage()
		addition := storage.GetAddition(add)
		if addition != nil {
			newTruck.Addition = append(newTruck.Addition, addition)
		} else {
			log.Printf("can not find transport: %s", add)
		}
	}
	log.Printf("newTruck: %v", newTruck)
	result := db.Omit("LoadingPoints").Omit("UnloadingPoints").Create(&newTruck)
	c.saveLoadingPoints(newTruck.ID, truck.LoadingPoints)
	c.saveUnloadingPoints(newTruck.ID, truck.UnloadingPoints)

	if result.Error != nil {
		log.Printf("Can not ctreate Truck %v", result.Error.Error())
		return nil, result.Error
	}
	return &newTruck, nil
}

func (c *TruckStorage) Update(truck *api.Truck, userid uint64) (*Truck, error) {
	var oldTruck *Truck
	result := c.db.First(&oldTruck, truck.GetId())
	if result.Error != nil {
		log.Printf("Truck %d not found", truck.GetId())
		return nil, result.Error
	}
	user := NewUserStorage().GetUser(userid)
	if user == nil {
		log.Printf("User %d not found", userid)
		return nil, nil
	}
	newTruck := &Truck{
		Model:           gorm.Model{ID: oldTruck.ID, CreatedAt: oldTruck.CreatedAt},
		From:            truck.From.AsTime(),
		Till:            truck.Till.AsTime(),
		Weight:          uint(truck.Weight),
		Volume:          uint(truck.Volume),
		Cost:            uint(truck.Cost),
		Currency:        truck.Currency,
		AdditionalPhone: strings.Join(truck.AdditionalPhones, ";"),
		AdditionalEmail: strings.Join(truck.AdditionalMails, ";"),
		Owner:           user,
	}
	_ = db.Model(&newTruck).Association("LoadingPoints").Clear()
	//for _, cityId := range truck.LoadingPoints {
	//	geo := NewGeoStorage()
	//	city := geo.GetCity(uint(cityId))
	//	if city != nil {
	//		newTruck.LoadingPoints = append(newTruck.LoadingPoints, city)
	//	} else {
	//		log.Printf("can not find city: %d", cityId)
	//	}
	//}
	_ = db.Model(&newTruck).Association("UnloadingPoints").Clear()
	//for _, cityId := range truck.UnloadingPoints {
	//	geo := NewGeoStorage()
	//	city := geo.GetCity(uint(cityId))
	//	if city != nil {
	//		newTruck.UnloadingPoints = append(newTruck.UnloadingPoints, city)
	//	} else {
	//		log.Printf("can not find city: %d", cityId)
	//	}
	//}
	_ = db.Model(&newTruck).Association("Transport").Clear()
	c.saveLoadingPoints(newTruck.ID, truck.LoadingPoints)
	c.saveUnloadingPoints(newTruck.ID, truck.UnloadingPoints)

	for _, transport := range truck.TransportType {
		storage := NewReferenceStorage()
		transportType := storage.GetTransportType(transport, 0)
		if transportType != nil {
			newTruck.Transport = append(newTruck.Transport, transportType)
		} else {
			log.Printf("can not find transport: %s", transport)
		}
	}
	_ = db.Model(&newTruck).Association("Type").Clear()
	if truck.CargoType != 0 {
		storage := NewReferenceStorage()
		cargoType := storage.GetCargoTypeById(truck.CargoType)
		if cargoType != nil {
			newTruck.Type = cargoType
			newTruck.TypeId = cargoType.ID
		} else {
			log.Printf("can not find cargotype: %d", truck.CargoType)
		}
	}
	_ = db.Model(&newTruck).Association("Loading").Clear()
	for _, loading := range truck.LoadingType {
		storage := NewReferenceStorage()
		loadingType := storage.GetLoadingType(loading)
		if loadingType != nil {
			newTruck.Loading = append(newTruck.Loading, loadingType)
		} else {
			log.Printf("can not find transport: %s", loading)
		}
	}
	_ = db.Model(&newTruck).Association("Addition").Clear()
	for _, add := range truck.Addition {
		storage := NewReferenceStorage()
		addition := storage.GetAddition(add)
		if addition != nil {
			newTruck.Addition = append(newTruck.Addition, addition)
		} else {
			log.Printf("can not find transport: %s", add)
		}
	}
	result = db.Save(&newTruck)
	if result.Error != nil {
		log.Printf("Can not ctreate Truck %v", result.Error.Error())
		return nil, result.Error
	}
	return newTruck, nil
}

func (c *TruckStorage) Delete(id uint64) uint64 {
	result := c.db.Delete(&Truck{}, id)
	if result.Error != nil {
		return 0
	}
	return id
}

func (c *TruckStorage) Get(id uint64) *Truck {
	var truck Truck
	result := db.
		Preload("Transport").
		Preload("Loading").
		Preload("Addition").
		Preload("Owner").
		First(&truck, id)
	if result.Error != nil {
		return nil
	}
	_ = c.loadingPoints(&truck)
	_ = c.unloadingPoints(&truck)
	return &truck
}

func (c *TruckStorage) Find(
	geoFrom []*api.Geo,
	geoTo []*api.Geo,
	weightFrom uint32,
	weightTo uint32,
	volumeFrom uint32,
	volumeTo uint32,
	from time.Time,
	to time.Time,
	limit uint32,
	offset uint32,
	typeLoading []uint32,
	timeFilter time.Time,
	cargoType uint64,

) ([]*Truck, int64) {
	var loadingPoints []uint64
	for _, geo := range geoFrom {
		switch geo.GetType() {
		case api.Geo_city:
			var city City
			c.db.First(&city, geo.GetGeoId())
			loadingPoints = append(loadingPoints, uint64(city.ID))
			break
		case api.Geo_region:
			var region Region
			c.db.Preload("Cities").First(&region, geo.GetGeoId())
			for _, r := range region.Cities {
				loadingPoints = append(loadingPoints, uint64(r.ID))
			}
			break
		case api.Geo_country:
			var cities []City
			c.db.Where("country_id = ?", geo.GetGeoId()).Find(&cities, "")
			for _, city := range cities {
				loadingPoints = append(loadingPoints, uint64(city.ID))
			}
			break
		}
	}

	var unloadingPoints []uint64
	for _, geo := range geoTo {
		switch geo.GetType() {
		case api.Geo_city:
			var city City
			c.db.First(&city, geo.GetGeoId())
			unloadingPoints = append(unloadingPoints, uint64(city.ID))
			break
		case api.Geo_region:
			var region Region
			c.db.Preload("Cities").First(&region, geo.GetGeoId())
			for _, r := range region.Cities {
				unloadingPoints = append(unloadingPoints, uint64(r.ID))
			}
			break
		case api.Geo_country:
			var cities []City
			c.db.Where("country_id = ?", geo.GetGeoId()).Find(&cities, "")
			for _, city := range cities {
				unloadingPoints = append(unloadingPoints, uint64(city.ID))
			}
			break
		}
	}
	tx := c.db.Joins("left join truck_loading_points on truck_loading_points.truck_id = trucks.id").
		Joins("left join truck_unloading_points on truck_unloading_points.truck_id = trucks.id").
		Joins("left join truck_loading on truck_loading.truck_id = trucks.id")
	if len(loadingPoints) > 0 {
		tx.Where("truck_loading_points.city_id in ?", loadingPoints)
	}
	if len(unloadingPoints) > 0 {
		tx.Or("truck_unloading_points.city_id in ?", unloadingPoints)
	}
	if len(typeLoading) > 0 {
		tx.Where("truck_loading.loading_type_id in ?", typeLoading)
	}
	if cargoType != 0 {
		tx.Where("type_id = ?", cargoType)
	}
	if weightFrom != 0 {
		tx.Where("weight >= ?", weightFrom)
	}
	if weightTo != 0 {
		tx.Where("weight <= ?", weightTo)
	}
	if volumeFrom != 0 {
		tx.Where("volume >= ?", volumeFrom)
	}
	if volumeTo != 0 {
		tx.Where("volume <= ?", volumeTo)
	}
	if from.Unix() != 0 {
		strFrom := from.Format("2006-01-02")
		tx.Where("trucks.`from`>= ?", strFrom)
	}
	if to.Unix() != 0 {
		strTo := to.Format("2006-01-02")
		tx.Where("trucks.`from` <= ?", strTo)
	}
	if timeFilter.Unix() != 0 {
		strFilter := from.Format("2006-01-02")
		tx.Where("trucks.`created_at` >= ?", strFilter)
	}

	var count int64
	tx.Model(Truck{}).Distinct().Count(&count)
	if limit != 0 {
		tx.Limit(int(limit))
	}
	tx.Offset(int(offset))
	var trucks []*Truck
	tx.Preload("Transport").
		Preload("Loading").
		Preload("Addition").
		Preload("Owner").
		Distinct("trucks.*").Order("created_at desc").Find(&trucks)
	for _, tr := range trucks {
		_ = c.loadingPoints(tr)
		_ = c.unloadingPoints(tr)
	}
	return trucks, count
}
