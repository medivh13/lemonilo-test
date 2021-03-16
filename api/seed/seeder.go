package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/medivh13/lemonilo-test/api/models"
)

var users = []models.User{
	models.User{
		Nickname: "Jody Almaida",
		Email:    "jody.almaida@gmail.com",
		Password: "jodyjody13",
	},
	models.User{
		Nickname: "John Doe",
		Email:    "jhon.doe@gmail.com",
		Password: "jhonjhon13",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
