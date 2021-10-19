package services

import (
	"anekakios/helpers"
	"anekakios/models"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateAccount(db *gorm.DB, user *models.User) (helpers.MessageResponse, *models.User) {
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	addPerson := models.User{
		Email:    user.Email,
		Name:     user.Name,
		Mobile:   user.Mobile,
		Address:  user.Address,
		Password: string(pass),
		RoleID:   user.RoleID,
	}
	if err := db.Create(&addPerson).Error; err != nil {
		return *helpers.MessageResponses(false, http.StatusBadRequest, "Cannot add User"), nil
	}
	return *helpers.MessageResponses(true, 200, "Succesfully"), &addPerson
}

func Login(db *gorm.DB, user *models.Login) (*models.LoginResponse, error) {
	var userTemp models.User
	email := strings.ToLower(user.Email)
	if err := db.Where("email = ?", email).Find(&userTemp).Error; err != nil {
		return nil, err
	}

	result := bcrypt.CompareHashAndPassword([]byte(userTemp.Password), []byte(user.Password))

	if result != nil {
		return nil, result
	}
	token, err := GenerateToken(&userTemp)

	if err != nil {
		return nil, err
	}
	return &models.LoginResponse{
		ID:    uint64(userTemp.ID),
		Email: userTemp.Email,
		Token: token,
	}, nil
}

func GetUserByID(db *gorm.DB, id int) (*models.User, error) {
	var user *models.User
	if err := db.Preload("Role").First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetAllUser(db *gorm.DB) (user []models.User, err error) {
	if err := db.Preload("Role").Find(&user).Error; err != nil {
		helpers.MessageResponses(false, http.StatusBadRequest, "User Not Foud")
	}
	return user, nil
}

func UpdateUser(db *gorm.DB, id int, user *models.User) (*models.User, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	var updateUser models.User
	if err := db.Where("id = ?", id).First(&updateUser).Error; err != nil {
		return nil, err
	}
	if user.Email != "" {
		updateUser.Email = user.Email
		updateUser.Password = string(pass)
		updateUser.Mobile = user.Mobile
		updateUser.Address = user.Address
		updateUser.RoleID = user.RoleID
		db.Save(&updateUser)
	}
	return &updateUser, nil
}

func DeleteUser(db *gorm.DB, id int) (map[string]string, error) {
	var user models.User
	if err := db.Delete(&user, id).Error; err != nil {
		return map[string]string{
			"message": "User tidak ada",
		}, err
	}
	return map[string]string{
		"message": "User telah terhapus",
	}, nil
}
