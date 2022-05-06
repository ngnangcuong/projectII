package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"fmt"
	"sync"
	"log"
	"projectII/models"
)

var (
	lock &sync.Mutex{};
	singletonPostgres *gorm.DB
)

func InitPostGres() *gorm.DB {
	if singletonPostgres == nil {
		lock.Lock()
		defer lock.Unlock()

		if singletonPostgres == nil {
			dbname := "projectII"
			db := "postgres"
			dbpassword := "Cuongnguyen2001"
			dburl := "postgres://postgres:" + dbpassword + "@/localhost/" + dbname + "?sslmode=disable"

			connection, err := gorm.Open(db, dburl)
			if err != nil {
				log.Fatalln("Wrong database url")
			}

			sqldb := connection.DB()
			err = sqldb.Ping()
			if err != nil {
				log.Fatalln("Database is not connected.")
			}

			fmt.Println("Database is connected")

			return connection
		}

		fmt.Println("Database is connected")
	}

	fmt.Println("Database is connected")

	return singletonPostgres
}

func InitMigration() {
	connection := InitPostGres()

	connection.AutoMigrate(models.EducationUser{})
	connection.AutoMigrate(models.EducationClass{})
	connection.AutoMigrate(models.EducationAssignment{})
	connection.AutoMigrate(models.EducationSubmission{})
	connection.AutoMigrate(models.EducationStudent{})
	connection.AutoMigrate(models.EducationTeacher{})
	connection.AutoMigrate(models.EducationTerm{})
}

func ClosePostgres() {
	sqldb := singletonPostgres.DB()
	sqldb.Close() 
}

