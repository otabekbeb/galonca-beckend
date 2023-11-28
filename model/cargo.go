package model

import (
	"gorm.io/gorm"
	"html"
	"log"
	"strings"
	"time"
	"tsl_server/api"
)

type Cargo struct {
	gorm.Model
	LoadingPoints   []*City `gorm:"many2many:cargo_loading_points;ForeignKey:id;References:id"`
	UnloadingPoints []*City `gorm:"many2many:cargo_unloading_points"`
	From            time.Time
	Till            time.Time
	Weight          uint
	Volume          uint
	Cost            uint
	Currency        string
	Transport       []*TransportType `gorm:"many2many:cargo_transport"`
	Loading         []*LoadingType   `gorm:"many2many:cargo_loading"`
	Addition        []*Addition      `gorm:"many2many:cargo_addition"`
	TypeId          uint
	Type            *CargoType
	AdditionalPhone string `gorm:"type:text"`
	AdditionalEmail string `gorm:"type:text"`
	OwnerID         uint
	Owner           *TslUser `gorm:"foreignKey:OwnerID"`
}

type CargoLoadingPoints struct {
	CargoId   uint64 `gorm:"primaryKey"`
	CityId    uint64 `gorm:"primaryKey"`
	SortOrder uint   `gorm:"not null,DEFAULT 0"`
}

type CargoUnloadingPoints struct {
	CargoId   uint64 `gorm:"primaryKey"`
	CityId    uint64 `gorm:"primaryKey"`
	SortOrder uint   `gorm:"not null,DEFAULT 0"`
}

type CargoStorage struct {
	db *gorm.DB
}

var cs *CargoStorage

func NewCargoStorage() *CargoStorage {
	if cs == nil {
		cs = &CargoStorage{
			db: GetDB(),
		}
		//_ = cs.db.SetupJoinTable(&Cargo{}, "LoadingPoints", &CargoLoadingPoints{})
	}
	return cs
}

func (c *CargoStorage) loadingPoints(cargo *Cargo) error {
	var lPoints []CargoLoadingPoints
	err := c.db.Model(&CargoLoadingPoints{}).Order("sort_order asc").Find(&lPoints, "cargo_id = ?", cargo.ID)
	if err.Error != nil {
		return err.Error
	}
	for _, p := range lPoints {
		var city City
		err = c.db.Model(&City{}).Find(&city, p.CityId)
		if err.Error == nil {
			city.SortOrder = p.SortOrder
			cargo.LoadingPoints = append(cargo.LoadingPoints, &city)
		}
	}
	return nil
}

func (c *CargoStorage) unloadingPoints(cargo *Cargo) error {
	var lPoints []CargoUnloadingPoints
	err := c.db.Model(&CargoUnloadingPoints{}).Order("sort_order asc").Find(&lPoints, "cargo_id = ?", cargo.ID)
	if err.Error != nil {
		return err.Error
	}
	for _, p := range lPoints {
		var city City
		err = c.db.Model(&City{}).Find(&city, p.CityId)
		if err.Error == nil {
			city.SortOrder = p.SortOrder
			cargo.UnloadingPoints = append(cargo.UnloadingPoints, &city)
		}
	}
	return nil
}

func (c *CargoStorage) saveLoadingPoints(cargoId uint, cities []uint64) {
	for index, cityId := range cities {
		geo := NewGeoStorage()
		city := geo.GetCity(uint(cityId))
		if city != nil {
			clp := CargoLoadingPoints{
				CargoId:   uint64(cargoId),
				CityId:    cityId,
				SortOrder: uint(index),
			}
			c.db.Save(&clp)
		} else {
			log.Printf("can not find city: %d", cityId)
		}
	}
}

func (c *CargoStorage) saveUnloadingPoints(cargoId uint, cities []uint64) {
	for index, cityId := range cities {
		geo := NewGeoStorage()
		city := geo.GetCity(uint(cityId))
		if city != nil {
			clp := CargoUnloadingPoints{
				CargoId:   uint64(cargoId),
				CityId:    cityId,
				SortOrder: uint(index),
			}
			c.db.Save(&clp)
		} else {
			log.Printf("can not find city: %d", cityId)
		}
	}
}

func (cs *Cargo) ToAPI() *api.FindCargoResult {

	cargo := api.FindCargoResult{
		Id:        uint64(cs.ID),
		CreatedAt: cs.CreatedAt.Format("02.01.2006"),
		UpdatedAt: cs.UpdatedAt.Format("02.01.2006"),
		From:      cs.From.Format("02.01.2006"),
		Till:      cs.Till.Format("02.01.2006"),
		Weight:    uint32(cs.Weight),
		Volume:    uint32(cs.Volume),
		Cost:      uint32(cs.Cost),
		Currency:  html.EscapeString(cs.Currency),
	}
	if cs.Type != nil {
		cargo.CargoType = cs.Type.Name
	}
	if cs.Owner != nil {
		owner := api.FindCargoResult_Owner{
			Id:   uint64(cs.Owner.ID),
			Name: cs.Owner.Name,
		}
		cargo.Owner = &owner
	}

	for _, lp := range cs.LoadingPoints {
		point := api.FindCargoResult_ShortGeo{
			Id:   uint64(lp.ID),
			Name: html.EscapeString(lp.Name),
			Type: 0,
		}
		cargo.LoadingPoints = append(cargo.LoadingPoints, &point)
	}
	for _, up := range cs.UnloadingPoints {
		point := api.FindCargoResult_ShortGeo{
			Id:   uint64(up.ID),
			Name: html.EscapeString(up.Name),
			Type: 0,
		}
		cargo.UnloadingPoints = append(cargo.UnloadingPoints, &point)
	}
	for _, tt := range cs.Transport {
		cargo.TransportType = append(cargo.TransportType, html.EscapeString(tt.Name))
	}
	for _, tt := range cs.Loading {
		cargo.LoadingType = append(cargo.LoadingType, html.EscapeString(tt.Name))
	}
	for _, tt := range cs.Addition {
		cargo.Addition = append(cargo.Addition, html.EscapeString(tt.Name))
	}
	cargo.AdditionalPhones = strings.Split(cs.AdditionalPhone, ";")
	cargo.AdditionalMails = strings.Split(cs.AdditionalEmail, ";")
	return &cargo
}

func (c *CargoStorage) Create(cargo *api.Cargo, userid uint64) (*Cargo, error) {
	user := NewUserStorage().GetUser(userid)
	if user == nil {
		log.Printf("User %d not found", userid)
		return nil, nil
	}
	newCargo := Cargo{
		From:            cargo.From.AsTime(),
		Till:            cargo.Till.AsTime(),
		Weight:          uint(cargo.Weight),
		Volume:          uint(cargo.Volume),
		Cost:            uint(cargo.Cost),
		Currency:        cargo.Currency,
		AdditionalPhone: strings.Join(cargo.AdditionalPhones, ";"),
		AdditionalEmail: strings.Join(cargo.AdditionalMails, ";"),
		OwnerID:         user.ID,
	}
	//for _, cityId := range cargo.LoadingPoints {
	//	geo := NewGeoStorage()
	//	city := geo.GetCity(uint(cityId))
	//	if city != nil {
	//		newCargo.LoadingPoints = append(newCargo.LoadingPoints, city)
	//	} else {
	//		log.Printf("can not find city: %d", cityId)
	//	}
	//}
	//for _, cityId := range cargo.UnloadingPoints {
	//	geo := NewGeoStorage()
	//	city := geo.GetCity(uint(cityId))
	//	if city != nil {
	//		newCargo.UnloadingPoints = append(newCargo.UnloadingPoints, city)
	//	} else {
	//		log.Printf("can not find city: %d", cityId)
	//	}
	//}
	for _, transport := range cargo.TransportType {
		storage := NewReferenceStorage()
		transportType := storage.GetTransportType(transport, 0)
		if transportType != nil {
			newCargo.Transport = append(newCargo.Transport, transportType)
		} else {
			log.Printf("can not find transport: %s", transport)
		}
	}

	for _, loading := range cargo.LoadingType {
		storage := NewReferenceStorage()
		loadingType := storage.GetLoadingType(loading)
		if loadingType != nil {
			newCargo.Loading = append(newCargo.Loading, loadingType)
		} else {
			log.Printf("can not find transport: %s", loading)
		}
	}
	if cargo.CargoType != 0 {
		storage := NewReferenceStorage()
		cargoType := storage.GetCargoTypeById(cargo.CargoType)
		if cargoType != nil {
			newCargo.Type = cargoType
			newCargo.TypeId = cargoType.ID
		} else {
			log.Printf("can not find cargotype: %d", cargo.CargoType)
		}
	}

	for _, add := range cargo.Addition {
		storage := NewReferenceStorage()
		addition := storage.GetAddition(add)
		if addition != nil {
			newCargo.Addition = append(newCargo.Addition, addition)
		} else {
			log.Printf("can not find transport: %s", add)
		}
	}
	log.Printf("newCargo: %v", newCargo)
	result := db.Omit("LoadingPoints").Omit("UnloadingPoints").Create(&newCargo)
	c.saveLoadingPoints(newCargo.ID, cargo.LoadingPoints)
	c.saveUnloadingPoints(newCargo.ID, cargo.UnloadingPoints)
	if result.Error != nil {
		log.Printf("Can not ctreate Cargo %v", result.Error.Error())
		return nil, result.Error
	}
	return &newCargo, nil
}

func (c *CargoStorage) Update(cargo *api.Cargo, userid uint64) (*Cargo, error) {
	var oldCargo *Cargo
	result := c.db.First(&oldCargo, cargo.GetId())
	if result.Error != nil {
		log.Printf("Cargo %d not found", cargo.GetId())
		return nil, result.Error
	}
	user := NewUserStorage().GetUser(userid)
	if user == nil {
		log.Printf("User %d not found", userid)
		return nil, nil
	}
	newCargo := &Cargo{
		Model:           gorm.Model{ID: oldCargo.ID, CreatedAt: oldCargo.CreatedAt},
		From:            cargo.From.AsTime(),
		Till:            cargo.Till.AsTime(),
		Weight:          uint(cargo.Weight),
		Volume:          uint(cargo.Volume),
		Cost:            uint(cargo.Cost),
		Currency:        cargo.Currency,
		AdditionalPhone: strings.Join(cargo.AdditionalPhones, ";"),
		AdditionalEmail: strings.Join(cargo.AdditionalMails, ";"),
		Owner:           user,
	}
	_ = db.Model(&newCargo).Association("LoadingPoints").Clear()
	//for _, cityId := range cargo.LoadingPoints {
	//	geo := NewGeoStorage()
	//	city := geo.GetCity(uint(cityId))
	//	if city != nil {
	//		newCargo.LoadingPoints = append(newCargo.LoadingPoints, city)
	//	} else {
	//		log.Printf("can not find city: %d", cityId)
	//	}
	//}
	_ = db.Model(&newCargo).Association("UnloadingPoints").Clear()
	c.saveLoadingPoints(newCargo.ID, cargo.LoadingPoints)
	c.saveUnloadingPoints(newCargo.ID, cargo.UnloadingPoints)

	//for _, cityId := range cargo.UnloadingPoints {
	//	geo := NewGeoStorage()
	//	city := geo.GetCity(uint(cityId))
	//	if city != nil {
	//		newCargo.UnloadingPoints = append(newCargo.UnloadingPoints, city)
	//	} else {
	//		log.Printf("can not find city: %d", cityId)
	//	}
	//}
	_ = db.Model(&newCargo).Association("Transport").Clear()
	for _, transport := range cargo.TransportType {
		storage := NewReferenceStorage()
		transportType := storage.GetTransportType(transport, 0)
		if transportType != nil {
			newCargo.Transport = append(newCargo.Transport, transportType)
		} else {
			log.Printf("can not find transport: %s", transport)
		}
	}
	_ = db.Model(&newCargo).Association("Type").Clear()
	if cargo.CargoType != 0 {
		storage := NewReferenceStorage()
		cargoType := storage.GetCargoTypeById(cargo.CargoType)
		if cargoType != nil {
			newCargo.Type = cargoType
			newCargo.TypeId = cargoType.ID
		} else {
			log.Printf("can not find cargotype: %d", cargo.CargoType)
		}
	}

	_ = db.Model(&newCargo).Association("Loading").Clear()
	for _, loading := range cargo.LoadingType {
		storage := NewReferenceStorage()
		loadingType := storage.GetLoadingType(loading)
		if loadingType != nil {
			newCargo.Loading = append(newCargo.Loading, loadingType)
		} else {
			log.Printf("can not find transport: %s", loading)
		}
	}

	_ = db.Model(&newCargo).Association("Addition").Clear()
	for _, add := range cargo.Addition {
		storage := NewReferenceStorage()
		addition := storage.GetAddition(add)
		if addition != nil {
			newCargo.Addition = append(newCargo.Addition, addition)
		} else {
			log.Printf("can not find transport: %s", add)
		}
	}
	log.Printf("newCargo: %v", newCargo.LoadingPoints)

	result = db.Omit("LoadingPoints").Omit("UnloadingPoints").Save(&newCargo)
	if result.Error != nil {
		log.Printf("Can not ctreate Cargo %v", result.Error.Error())
		return nil, result.Error
	}
	return newCargo, nil
}

func (c *CargoStorage) Delete(id uint64) uint64 {
	result := c.db.Delete(&Cargo{}, id)
	if result.Error != nil {
		return 0
	}
	return id
}

func (c *CargoStorage) Get(id uint64) *Cargo {
	var cargo Cargo
	result := db.
		Preload("Transport").
		Preload("Loading").
		Preload("Addition").
		Preload("Owner").
		Preload("Type").
		First(&cargo, id)
	if result.Error != nil {
		return nil
	}
	_ = c.loadingPoints(&cargo)
	_ = c.unloadingPoints(&cargo)
	return &cargo
}

func (c *CargoStorage) Find(
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

) ([]*Cargo, int64) {
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
	tx := c.db.Joins("left join cargo_loading_points on cargo_loading_points.cargo_id = cargos.id").
		Joins("left join cargo_unloading_points on cargo_unloading_points.cargo_id = cargos.id").
		Joins("left join cargo_loading on cargo_loading.cargo_id = cargos.id")
	if len(loadingPoints) > 0 {
		tx.Where("cargo_loading_points.city_id in ?", loadingPoints)
	}
	if len(unloadingPoints) > 0 {
		tx.Or("cargo_unloading_points.city_id in ?", unloadingPoints)
	}
	if len(typeLoading) > 0 {
		tx.Where("cargo_loading.loading_type_id in ?", typeLoading)
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
		tx.Where("cargos.`from` >= ?", strFrom)
	}
	if to.Unix() != 0 {
		strTo := to.Format("2006-01-02")
		tx.Where("cargos.`from` <= ?", strTo)
	}
	if timeFilter.Unix() != 0 {
		strFilter := from.Format("2006-01-02")
		tx.Where("created_at >= ?", strFilter)
	}
	var count int64
	tx.Model(Cargo{}).Distinct("id").Count(&count)
	log.Printf("count %d", count)

	if limit != 0 {
		tx.Limit(int(limit))
	}
	tx.Offset(int(offset))
	var cargos []*Cargo
	tx.
		Preload("Transport").
		Preload("Loading").
		Preload("Addition").
		Preload("Owner").
		Preload("Type").
		Distinct("cargos.*").Order("created_at desc").Find(&cargos)
	for _, car := range cargos {
		_ = c.loadingPoints(car)
		_ = c.unloadingPoints(car)
	}
	return cargos, count
}
