package model

import (
	"gorm.io/gorm"
	"html"
	"tsl_server/api"
)

type Review struct {
	gorm.Model
	positive   bool
	text       string
	entityId   uint64
	entityType string `gorm:"type:VARCHAR(100)"`
	ipAddress  string `gorm:"type:VARCHAR(25)"`
}

func (r *Review) ToAPI() *api.Review {
	return &api.Review{
		Positive:   r.positive,
		Text:       html.EscapeString(r.text),
		EntityId:   r.entityId,
		EntityType: r.entityType,
	}
}

func (rs *ReferenceStorage) CreateReview(r *api.Review, ip string) *Review {
	review := Review{
		positive:   r.Positive,
		entityId:   r.EntityId,
		entityType: r.EntityType,
		text:       html.EscapeString(r.Text),
		ipAddress:  ip,
	}
	result := rs.db.Create(&review)
	if result.Error != nil {
		return nil
	}
	return &review
}

func (rs *ReferenceStorage) ListReview(r *api.ReviewRequest) *api.ReviewResponse {
	var reviews []Review
	rs.db.Where("entity_type = ? and entity_id = ?", r.EntityType, r.EntityId).Find(&reviews)
	var result api.ReviewResponse
	for _, rw := range reviews {
		result.Reviews = append(result.Reviews, rw.ToAPI())
	}
	return &result
}

type Like struct {
	gorm.Model
	Positive   bool
	EntityId   uint64
	EntityType string `gorm:"type:VARCHAR(100)"`
	IpAddress  string `gorm:"type:VARCHAR(25)"`
}

func (l *Like) ToAPI() *api.Like {
	return &api.Like{
		Positive:   l.Positive,
		EntityId:   l.EntityId,
		EntityType: l.EntityType,
	}
}

func (rs *ReferenceStorage) AddLike(l *api.Like, ip string) *Like {
	like := Like{
		Positive:   l.Positive,
		EntityId:   l.EntityId,
		EntityType: l.EntityType,
		IpAddress:  ip,
	}
	result := rs.db.Create(&like)
	if result.Error != nil {
		return nil
	}
	return &like
}

func (rs *ReferenceStorage) GetLikes(l *api.LikeRequest) *api.LikeResponse {
	like := api.LikeResponse{}
	rs.db.Model(Like{}).
		Where("entity_id = ? and entity_type = ? and positive = true", l.EntityId, l.EntityType).
		Count(&like.Positive)
	rs.db.Model(Like{}).
		Where("entity_id = ? and entity_type = ? and positive = false", l.EntityId, l.EntityType).
		Count(&like.Negative)
	return &like
}

type Brand struct {
	gorm.Model
	Name string    `gorm:"type:VARCHAR(100)"`
	Type api.RType `gorm:"type:INT"`
}

func (b *Brand) ToAPI() *api.Reference {
	var brandType api.RType
	switch b.Type {
	case 0:
		brandType = api.RType_autoBrand
		break
	case 1:
		brandType = api.RType_busBrand
		break
	case 2:
		brandType = api.RType_specialBrand
		break
	}
	return &api.Reference{
		Id:   uint64(b.ID),
		Name: b.Name,
		Type: brandType,
	}
}

type TransportType struct {
	gorm.Model
	Name string    `gorm:"type:VARCHAR(100)"`
	Type api.RType `gorm:"type:INT"`
}

func (t *TransportType) ToAPI() *api.Reference {
	var tt api.RType
	switch t.Type {
	case 0:
		tt = api.RType_transportType
		break
	case 1:
		tt = api.RType_busTransportType
		break
	case 2:
		tt = api.RType_specialTransportType
		break
	}
	return &api.Reference{
		Id:   uint64(t.ID),
		Name: html.EscapeString(t.Name),
		Type: tt,
	}
}

type CargoType struct {
	gorm.Model
	Name string    `gorm:"type:VARCHAR(100)"`
	Type api.RType `gorm:"type:INT"`
}

func (t *CargoType) ToAPI() *api.Reference {
	return &api.Reference{
		Id:   uint64(t.ID),
		Name: html.EscapeString(t.Name),
		Type: api.RType_cargoType,
	}
}

type LoadingType struct {
	gorm.Model
	Name string `gorm:"type:VARCHAR(100)"`
}

func (t *LoadingType) ToAPI() *api.Reference {
	return &api.Reference{
		Id:   uint64(t.ID),
		Name: html.EscapeString(t.Name),
		Type: api.RType_loadingType,
	}
}

type Addition struct {
	gorm.Model
	Name string `gorm:"type:VARCHAR(100)"`
}

func (t *Addition) ToAPI() *api.Reference {
	return &api.Reference{
		Id:   uint64(t.ID),
		Name: html.EscapeString(t.Name),
		Type: api.RType_addition,
	}
}

type ServiceCategory struct {
	gorm.Model
	Name string `gorm:"type:VARCHAR(100)"`
}

func (t *ServiceCategory) ToAPI() *api.Reference {
	return &api.Reference{
		Id:   uint64(t.ID),
		Name: html.EscapeString(t.Name),
		Type: api.RType_serviceCategory,
	}
}

type ReferenceStorage struct {
	db *gorm.DB
}

var rs *ReferenceStorage

func NewReferenceStorage() *ReferenceStorage {
	if rs == nil {
		rs = &ReferenceStorage{
			db: GetDB(),
		}
	}
	return rs
}

func (rs *ReferenceStorage) GetBrand(name string, rType uint32) *Brand {
	var b *Brand
	result := rs.db.Where("name = ? and type = ?", name, rType).First(&b)
	if result.Error != nil {
		return nil
	}
	return b
}

func (rs *ReferenceStorage) GetCargoType(name string) *CargoType {
	var b *CargoType
	result := rs.db.Where("name = ?", name).First(&b)
	if result.Error != nil {
		return nil
	}
	return b
}

func (rs *ReferenceStorage) GetCargoTypeById(id uint64) *CargoType {
	var b *CargoType
	result := rs.db.First(&b, id)
	if result.Error != nil {
		return nil
	}
	return b
}

func (rs *ReferenceStorage) GetBrandById(id uint64) *Brand {
	var b *Brand
	result := rs.db.First(&b, id)
	if result.Error != nil {
		return nil
	}
	return b
}

func (rs *ReferenceStorage) GetTransportType(name string, rType uint32) *TransportType {
	var tt *TransportType
	result := rs.db.Where("name = ? and type = ?", name, rType).First(&tt)
	if result.Error != nil {
		return nil
	}
	return tt
}

func (rs *ReferenceStorage) GetTransportTypeById(id uint64) *TransportType {
	var tt *TransportType
	result := rs.db.First(&tt, id)
	if result.Error != nil {
		return nil
	}
	return tt
}

func (rs *ReferenceStorage) GetLoadingType(name string) *LoadingType {
	var lt *LoadingType
	result := rs.db.Where("name = ?", name).First(&lt)
	if result.Error != nil {
		return nil
	}
	return lt
}

func (rs *ReferenceStorage) GetAddition(name string) *Addition {
	var a *Addition
	result := rs.db.Where("name = ?", name).First(&a)
	if result.Error != nil {
		return nil
	}
	return a
}

func (rs *ReferenceStorage) GetServiceCategory(Id uint64) *ServiceCategory {
	var sc *ServiceCategory
	result := rs.db.First(&sc, Id)
	if result.Error != nil {
		return nil
	}
	return sc
}

func (rs *ReferenceStorage) GetServiceCategoryByName(name string) *ServiceCategory {
	var sc *ServiceCategory
	result := rs.db.Where("name = ?", name).First(&sc)
	if result.Error != nil {
		return nil
	}
	return sc
}

func (rs *ReferenceStorage) GetBrandList(order string, rType uint32) []Brand {
	if order == "r_name" {
		order = "name"
	}
	var b []Brand
	result := rs.db.Where("type = ?", rType).Order(order).Find(&b)
	if result.Error != nil {
		return nil
	}
	return b
}

func (rs *ReferenceStorage) GetCargoTypeList(order string) []CargoType {
	if order == "r_name" {
		order = "name"
	}
	var b []CargoType
	result := rs.db.Order("type, " + order).Find(&b).Find(&b)
	if result.Error != nil {
		return nil
	}
	return b
}

func (rs *ReferenceStorage) GetTransportTypeList(order string, rType uint32) []TransportType {
	if order == "r_name" {
		order = "name"
	}
	var tt []TransportType
	result := rs.db.Where("type = ?", rType).Order(order).Find(&tt)
	if result.Error != nil {
		return nil
	}
	return tt
}

func (rs *ReferenceStorage) GetLoadingTypeList(order string) []LoadingType {
	if order == "r_name" {
		order = "name"
	}
	var lt []LoadingType
	result := rs.db.Order(order).Find(&lt)
	if result.Error != nil {
		return nil
	}
	return lt
}

func (rs *ReferenceStorage) GetAdditionList(order string) []Addition {
	if order == "r_name" {
		order = "name"
	}
	var a []Addition
	result := rs.db.Order(order).Find(&a)
	if result.Error != nil {
		return nil
	}
	return a
}

func (rs *ReferenceStorage) GetServiceCategoryList(order string) []ServiceCategory {
	if order == "r_name" {
		order = "name"
	}
	var sc []ServiceCategory
	result := rs.db.Order(order).Find(&sc)
	if result.Error != nil {
		return nil
	}
	return sc
}
