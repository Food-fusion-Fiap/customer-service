package database

import (
	"fmt"
	"log"
	"os"

	"github.com/CAVAh/api-tech-challenge/src/infra/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB Database
)

// Database is an interface that defines methods for interacting with a database.
type Database interface {
	Create(data interface{}) error
	Where(query interface{}, args ...interface{}) *gorm.DB
	Find(dest interface{}, conds ...interface{}) *gorm.DB
}

type RealDatabase struct {
	db *gorm.DB
}

func (rdb *RealDatabase) Create(data interface{}) error {
	return rdb.db.Create(data).Error
}

func (rdb *RealDatabase) Where(query interface{}, args ...interface{}) *gorm.DB {
	return rdb.db.Where(query, args...)
}

func (rdb *RealDatabase) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	return rdb.db.Find(dest, conds...)
}

func ConnectDB() {
	conectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Fortaleza",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	db, err := gorm.Open(postgres.Open(conectionString))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB = &RealDatabase{
		db: db,
	}

	err = db.AutoMigrate(&models.Customer{})
	if err != nil {
		log.Panic("Erro ao fazer auto migrate")
	}
}
