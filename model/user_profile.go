package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"html"
	"log"
	"strings"
	"tsl_server/api"
)

type Favorite struct {
	gorm.Model
	EntityType string `gorm:"type:VARCHAR(100)"`
	EntityId   uint64
	OwnerID    uint64
}

func (f *Favorite) ToAPI() *api.Favorite {
	return &api.Favorite{
		EntityType: f.EntityType,
		Id:         f.EntityId,
	}
}

type FavoriteStorage struct {
	tableName string
	db        *gorm.DB
}

var favS *FavoriteStorage

func NewFavoriteStorage() *FavoriteStorage {
	if favS == nil {
		favS = &FavoriteStorage{
			tableName: "favorites",
			db:        GetDB(),
		}
	}
	return favS
}

func (f *FavoriteStorage) ListMy(model interface{}, limit uint32, offset uint32, userId uint) (interface{}, int) {
	tx := f.db.Where("owner_id = ?", userId)
	if limit != 0 {
		tx.Limit(int(limit)).Offset(int(offset))
	}

	switch model.(type) {
	case Cargo:
		var c []Cargo
		tx.Preload(clause.Associations).Find(&c)
		return c, len(c)
	case Transport:
		var r []Transport
		tx.Preload(clause.Associations).Find(&r)
		return r, len(r)
	case RoadsideService:
		var r []RoadsideService
		tx.Preload(clause.Associations).Find(&r)
		return r, len(r)
	case ServiceStation:
		var r []ServiceStation
		tx.Preload(clause.Associations).Find(&r)
		return r, len(r)
	case SparePart:
		var r []SparePart
		tx.Preload(clause.Associations).Find(&r)
		return r, len(r)
	case Truck:
		var r []Truck
		tx.Preload(clause.Associations).Find(&r)
		return r, len(r)
	}
	return nil, -1
}

func (f *FavoriteStorage) AddFavorite(fav *api.Favorite, userId uint64) *Favorite {
	favorite := Favorite{
		EntityType: fav.EntityType,
		EntityId:   fav.GetId(),
		OwnerID:    userId,
	}
	result := f.db.Create(&favorite)
	if result.Error != nil {
		log.Printf("Can not ctreate favorite: %v", result.Error)
		return nil
	}
	return &favorite
}

func (f *FavoriteStorage) DeleteFavorite(fav *api.Favorite, userId uint64) *Favorite {
	favorite := Favorite{
		EntityType: fav.EntityType,
		EntityId:   fav.GetId(),
		OwnerID:    userId,
	}
	result := f.db.Where("entity_type=?", favorite.EntityType).
		Where("entity_id=?", favorite.EntityId).
		Where("owner_id=?", favorite.OwnerID).
		Delete(&favorite)
	if result.Error != nil {
		log.Printf("Can not delete favorite: %v", result.Error)
		return nil
	}
	return &favorite
}

func (f *FavoriteStorage) ListFavorites(req *api.FavoritesRequest, userId uint64) *api.FavoritesResponse {
	var favorites []Favorite
	tx := f.db.Where("owner_id = ? and entity_type = ?", userId, req.EntityType)
	if req.Limit != 0 {
		tx.Limit(int(req.Limit))
	}
	if req.Offset != 0 {
		tx.Offset(int(req.Offset))
	}
	tx.Find(&favorites)
	var result api.FavoritesResponse
	for _, fav := range favorites {
		result.Favorite = append(result.Favorite, fav.ToAPI())
	}
	return &result
}

type Company struct {
	gorm.Model
	Name         string `gorm:"type:VARCHAR(100);NOT NULL"`
	BIN          string `gorm:"type:VARCHAR(60);NOT NULL;uniqueIndex"`
	Phone        string `gorm:"type:VARCHAR(20);NOT NULL;index"`
	Email        string `gorm:"type:VARCHAR(255);NOT NULL;index"`
	ActivityType string `gorm:"type:VARCHAR(255);NOT NULL"`
	CityId       uint64
	Address      string `gorm:"type:VARCHAR(255);NOT NULL"`
	Documents    string
	Images       string
	OwnerID      uint64
	Owner        *TslUser `gorm:"foreignKey:OwnerID"`
	Employees    []TslUser
}

func (c *Company) ToAPI() api.Company {
	return api.Company{
		Id:           int64(c.ID),
		Name:         html.EscapeString(c.Name),
		Bin:          html.EscapeString(c.BIN),
		Email:        html.EscapeString(c.Email),
		Phone:        html.EscapeString(c.Phone),
		CityId:       c.CityId,
		Address:      html.EscapeString(c.Address),
		ActivityType: c.ActivityType,
		Documents:    strings.Split(c.Documents, ";"),
		Images:       strings.Split(c.Images, ";"),
	}
}

type CompanyStorage struct {
	tableName string
	db        *gorm.DB
}

var compS *CompanyStorage

func NewCompanyStorage() *CompanyStorage {
	if compS == nil {
		compS = &CompanyStorage{
			tableName: "company",
			db:        GetDB(),
		}
	}
	return compS
}

func (s *CompanyStorage) Create(c *api.Company, userid uint64) (*Company, error) {
	user := NewUserStorage().GetUser(userid)
	company := Company{
		Name:         html.EscapeString(c.Name),
		Email:        html.EscapeString(c.Email),
		Phone:        html.EscapeString(c.Phone),
		BIN:          html.EscapeString(c.Bin),
		ActivityType: c.ActivityType,
		CityId:       c.CityId,
		Address:      c.Address,
		Documents:    strings.Join(c.Documents, ";"),
		Images:       strings.Join(c.Images, ";"),
		OwnerID:      uint64(user.ID),
	}
	var result *gorm.DB
	if c.Id != 0 {
		company.ID = uint(c.Id)
		result = s.db.Omit("created_at").Save(&company)
	} else {
		result = s.db.Create(&company)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &company, nil
}

func (s *CompanyStorage) Delete(id uint64) uint64 {
	result := s.db.Delete(&Company{}, id)
	if result.Error != nil {
		return 0
	}
	return id
}

func (s *CompanyStorage) GetCompany(id uint64) *Company {
	c := Company{}
	result := s.db.Preload(clause.Associations).First(&c, id)
	if result.Error != nil {
		return nil
	}
	return &c
}

func (s *CompanyStorage) GetCompanyByUser(email string) *Company {
	user := NewUserStorage().GetUserByEmail(email)
	c := new(Company)
	result := s.db.Preload(clause.Associations).Where("owner_id = ?", user.ID).First(c)
	if result.Error != nil {
		return nil
	}
	return c
}

func (s *CompanyStorage) AddEmployee(employee *api.CreateUserRequest, userid uint64) *TslUser {
	owner := NewUserStorage().GetUser(userid)
	if owner == nil {
		log.Printf("Invalid user!")
		return nil
	}
	company := NewCompanyStorage().GetCompanyByUser(owner.Email)
	if company == nil {
		log.Printf("Invalid company!")
		return nil
	}
	us, err := NewUserStorage().Create(0, employee.Name, employee.Password, employee.Email, employee.Phone, "user")
	if err != nil {
		log.Printf("Can't create user %v", err)
		return nil
	}
	us.CompanyID = new(uint)
	*us.CompanyID = company.ID
	s.db.Save(us)
	return us
}

func (s *CompanyStorage) ListEmployees(company *api.GetCompanyRequest) *api.EmployeesResponse {
	var emp []TslUser
	s.db.Where("company_id = ?", company.Id).Find(&emp)
	var result api.EmployeesResponse
	for _, e := range emp {
		result.Employees = append(result.Employees, e.ToAPI())
	}
	return &result
}

func (s *CompanyStorage) RemoveEmployees(employee *api.DeleteUserRequest) *api.CommonUserResponse {
	return NewUserStorage().Delete(employee.Id)
}
