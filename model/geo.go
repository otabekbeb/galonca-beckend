package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"html"
	"tsl_server/api"
)

type City struct {
	ID        uint   `gorm:"primary key"`
	Name      string `gorm:"type:VARCHAR(128);NOT NULL;index"`
	CountryID uint
	RegionID  uint
	Country   Country
	Region    Region
	SortOrder uint `gorm:"-"`
}

type Country struct {
	ID      uint   `gorm:"primary key"`
	Name    string `gorm:"type:VARCHAR(128);NOT NULL;index"`
	Regions []Region
}

type Region struct {
	ID        uint   `gorm:"primary key"`
	Name      string `gorm:"type:VARCHAR(128);NOT NULL;index"`
	CountryID uint
	Country   Country
	Cities    []City
}

func (c *City) ToAPI() *api.City {
	return &api.City{
		Id:      uint64(c.ID),
		Name:    html.EscapeString(c.Name),
		Region:  c.Region.ToAPI(),
		Country: c.Country.ToAPI(),
	}
}

func (c *Country) ToAPI() *api.Country {
	return &api.Country{
		Id:   uint64(c.ID),
		Name: html.EscapeString(c.Name),
	}
}

func (r *Region) ToAPI() *api.Region {
	return &api.Region{
		Id:   uint64(r.ID),
		Name: html.EscapeString(r.Name),
	}
}

type GeoStorage struct {
	db *gorm.DB
}

var gs *GeoStorage

func NewGeoStorage() *GeoStorage {
	if gs == nil {
		gs = &GeoStorage{
			db: GetDB(),
		}
	}
	return gs
}

func (g *GeoStorage) ListCountry(name string) []Country {
	var Countries []Country
	g.db.Where("name like ?", "%"+name+"%").Find(&Countries)
	return Countries
}

func (g *GeoStorage) ListCity(name string) []City {
	var Cities []City
	g.db.Where("name like ?", "%"+name+"%").
		Preload("Region").
		Preload("Country").
		Clauses(clause.OrderBy{
			Expression: clause.Expr{
				SQL:                "FIELD(country_id,?) desc, name",
				Vars:               []interface{}{[]int{9705, 81, 245, 248, 2303, 2788, 9575, 9787, 9908, 3159, 1894}},
				WithoutParentheses: true,
			},
		}).
		Find(&Cities)
	return Cities
}

func (g *GeoStorage) GetCity(id uint) *City {
	var city *City
	result := g.db.Preload("Region").
		Preload("Country").First(&city, id)
	if result.Error != nil {
		return nil
	}
	return city
}

func (g *GeoStorage) GetCountry(id uint) *Country {
	var country *Country
	result := g.db.First(&country, id)
	if result.Error != nil {
		return nil
	}
	return country
}

func (g *GeoStorage) GetRegion(id uint) *Region {
	var region *Region
	result := g.db.Preload("Country").Find(&region, id)
	if result.Error != nil {
		return nil
	}
	return region
}

func (g *GeoStorage) ListRegion(name string) []Region {
	var Regions []Region
	g.db.Where("name like ?", "%"+name+"%").Find(&Regions)
	return Regions
}
