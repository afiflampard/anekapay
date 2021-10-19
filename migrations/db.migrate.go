package migrations

import (
	"anekakios/models"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	tableExist := (db.Migrator().HasTable(&models.User{}) && db.Migrator().HasTable(&models.Role{}) && db.Migrator().HasTable(&models.Product{}) && db.Migrator().HasTable(&models.Order{}) && db.Migrator().HasTable(&models.OrderState{}) && db.Migrator().HasTable(&models.OrderDetail{}))
	if !tableExist {
		dbMigrate := db.Debug().Migrator().DropTable(&models.User{}, &models.Role{}, &models.Product{}, &models.Order{}, &models.OrderState{}, &models.OrderDetail{})
		if dbMigrate != nil {
			log.Fatal("Cannot drop Table")
		}
		db.Debug().AutoMigrate(&models.User{}, &models.Role{}, &models.Product{}, &models.Order{}, &models.OrderState{}, &models.OrderDetail{})

		var Roles = []models.Role{
			{
				Role: "User",
			},
			{
				Role: "Distributor",
			},
			{
				Role: "Pembeli",
			},
		}
		pass, err := bcrypt.GenerateFromPassword([]byte("1234"), bcrypt.MinCost)
		if err != nil {
			log.Println(err)
		}

		var Users = []models.User{
			{
				Email:    "afiflampard32@gmail.com",
				Name:     "Afif",
				Mobile:   "08576543434",
				Address:  "Jabon",
				Password: string(pass),
				RoleID:   1,
			},
			{
				Email:    "afiflampard132@gmail.com",
				Name:     "Fifa",
				Mobile:   "08576543434",
				Address:  "Jabon",
				Password: string(pass),
				RoleID:   2,
			},
			{
				Email:    "afiflampard09@gmail.com",
				Name:     "Fif",
				Mobile:   "08576543434",
				Address:  "Jabon",
				Password: string(pass),
				RoleID:   3,
			},
		}
		var orderStates = []models.OrderState{
			{
				Name: "Pesanan",
			},
			{
				Name: "Pembelian",
			},
		}

		for _, role := range Roles {
			err := db.Debug().Create(&role).Error
			if err != nil {
				log.Fatalf("failed to create Role")
			}
		}
		for _, user := range Users {
			if err := db.Debug().Create(&user).Error; err != nil {
				log.Fatal("Failed to create User")
			}
		}
		for _, orderState := range orderStates {
			if err := db.Debug().Create(&orderState).Error; err != nil {
				log.Fatalf("Failed to create Order State")
			}
		}

	}
}
