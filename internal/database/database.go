package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"iam-service/internal/model"
)

var once sync.Once
var singletonDB *gorm.DB

// GetDatabase returns a pointer to a `gorm.DB` ojbect.
func GetDatabase() *gorm.DB {
	// ensures that the code block inside the function is executed exactly once
	once.Do(func() {
		dsn := fmt.Sprintf(
			"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
			os.Getenv("POSTGRES_HOST"),
			"5432",
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err != nil {
			log.Fatalln("Cannot connect to database with DSN: ", dsn)
		}
		singletonDB = db.Debug() // logs SQL statements and additional information
	})

	return singletonDB
}

// AutoMigrate initializes and performs automatic migrations on the provided `gorm.DB` object.
// It ensures that the database schema is synchronized with the Gorm model definitions.
func AutoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(
		&model.Application{},
		&model.Tenant{},
		&model.Subject{},
		&model.TenantSubject{},
		&model.TenantRole{},
		&model.TenantRoleSubject{},
		&model.TenantResource{},
		&model.TenantResourceAction{},
		&model.TenantRoleAction{},
	); err != nil {
		log.Fatalln("Database migrate failed with error: ", err)
	}
}
