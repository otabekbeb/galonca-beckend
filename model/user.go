package model

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"html"
	"log"
	"os"
	"strconv"
	"tsl_server/api"
)

type TslUser struct {
	gorm.Model
	Name      string `gorm:"type:VARCHAR(100);NOT NULL"`
	Password  string `gorm:"type:VARCHAR(256);NOT NULL"`
	Email     string `gorm:"type:VARCHAR(60);NOT NULL;uniqueIndex"`
	Phone     string `gorm:"type:VARCHAR(20);NOT NULL;index"`
	Role      string `gorm:"type:VARCHAR(20);NOT NULL;default:'user'"`
	CompanyID *uint
}

func (user *TslUser) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func (user *TslUser) ToAPI() *api.TslUser {
	return &api.TslUser{
		Id:        int64(user.ID),
		Name:      html.EscapeString(user.Name),
		Email:     html.EscapeString(user.Email),
		Phone:     html.EscapeString(user.Phone),
		Role:      html.EscapeString(user.Role),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

type UserStorage struct {
	tableName string
	db        *gorm.DB
}

var us *UserStorage

func NewUserStorage() *UserStorage {
	if us == nil {
		us = &UserStorage{
			tableName: "tsl_users",
			db:        GetDB(),
		}
	}
	return us
}

func (s *UserStorage) Create(id uint64, name string, password string, email string, phone string, role string) (*TslUser, error) {

	user := TslUser{
		Name:      html.EscapeString(name),
		Email:     html.EscapeString(email),
		Phone:     html.EscapeString(phone),
		Role:      html.EscapeString(role),
		CompanyID: nil,
	}
	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return nil, fmt.Errorf("cannot hash password: %w", err)
		}
		user.Password = string(hashedPassword)
	}

	var result *gorm.DB
	if id != 0 {
		user.ID = uint(id)
		log.Printf("User  %v", user)
		result = s.db.Omit("created_at", "password", "role", "company_id").Save(&user)
	} else {
		result = s.db.Create(&user)
		_ = os.Mkdir("/var/www/tsl/uploads/image/user"+strconv.Itoa(int(user.ID)), 0775)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (s *UserStorage) GetUser(id uint64) *TslUser {
	u := TslUser{}
	result := s.db.First(&u, id)
	if result.Error != nil {
		return nil
	}
	return &u
}

func (s *UserStorage) Delete(id uint64) *api.CommonUserResponse {
	u := TslUser{}
	result := s.db.First(&u, id)
	if result.Error != nil {
		return &api.CommonUserResponse{
			Success: false,
			Error:   "User not found",
		}
	}
	s.db.Where("owner_id=?", id).Delete(&Cargo{})
	s.db.Where("owner_id=?", id).Delete(&Truck{})
	s.db.Where("owner_id=?", id).Delete(&ServiceStation{})
	s.db.Where("owner_id=?", id).Delete(&SparePart{})
	s.db.Where("owner_id=?", id).Delete(&RoadsideService{})
	s.db.Where("owner_id=?", id).Delete(&Transport{})
	if u.Role == "admin" && u.CompanyID != nil {
		s.db.Where("company_id=?", u.CompanyID).Delete(&TslUser{})
		s.db.Delete(&Company{}, u.CompanyID)
	}
	s.db.Delete(&u)
	return &api.CommonUserResponse{
		Success: true,
		Error:   "",
	}
}

func (s *UserStorage) GetUserByEmail(email string) *TslUser {
	u := TslUser{}
	result := s.db.Where("email = ?", email).First(&u)
	if result.Error != nil {
		return nil
	}
	return &u
}

func (s UserStorage) ChangePassword(req *api.ChangePasswordRequest) *api.CommonUserResponse {
	user := s.GetUser(uint64(req.GetId()))
	if user == nil {
		return &api.CommonUserResponse{
			Success: false,
			Error:   "User not found",
		}
	}
	if req.GetCheckOld() {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.GetOldPassword()))
		if err != nil {
			return &api.CommonUserResponse{
				Success: false,
				Error:   "Old password is incorrect",
			}
		}
	}
	if req.GetChkPassword() != "" && req.GetNewPassword() != req.GetChkPassword() {
		return &api.CommonUserResponse{
			Success: false,
			Error:   "New passwords not match",
		}
	}
	newPass, err := bcrypt.GenerateFromPassword([]byte(req.GetNewPassword()), bcrypt.DefaultCost)
	if err != nil {
		return &api.CommonUserResponse{
			Success: false,
			Error:   err.Error(),
		}
	}
	user.Password = string(newPass)
	s.db.Omit("created_at").Save(&user)
	return &api.CommonUserResponse{
		Success: true,
		Error:   "",
	}
}
