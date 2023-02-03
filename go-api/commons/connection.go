package commons

import (
	"go-api/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
	db, error := gorm.Open("mysql", "lautaro:Aced*2020@/usuarios?parseTime=true")

	if error != nil {
		log.Fatal(error)
	}
	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()

	log.Println("migrando...")
	db.AutoMigrate(&models.Persona{})
}
