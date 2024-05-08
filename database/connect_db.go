package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/udodinho/daily-standup-app/domain"
	"github.com/udodinho/daily-standup-app/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(&entity.CheckIn{})

	return err
}

func (pgs *Database) init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	configFile := &Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := Connect(configFile)

	if err != nil {
		log.Fatal("Could not connect to db, ", err)
	}

	err = MigrateDB(db)

	if err != nil {
		log.Fatal("Could not migrate database, ", err)
	}

	pgs.Db = db

	fmt.Println("Database connected successfully")

}

// Connect creates connection to database
func Connect(config *Config) (*gorm.DB, error) {
	var dsn string

	databaseUrl := os.Getenv("DATABASE_URL")

	if databaseUrl == "" {
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)
	} else {
		dsn = databaseUrl
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewDatabase() domain.DB {
	db := new(Database)
	db.init()
	return db
}

// CreateUpdate creates update for an employee
func (pg *Database) CreateUpdate(checkin *entity.CheckIn) (*entity.CheckIn, error) {
	err := pg.Db.Create(checkin).Error
	return nil, err
}

// FindAll retrieves all the updates and by a given filter
func (pg *Database) FindAll(weekStart, weekEnd, sprint, day, owner string) ([]entity.CheckIn, error) {
	var updates []entity.CheckIn

	// Construct the base query
	query := pg.Db.Model(&entity.CheckIn{})

	// Apply filter conditions
	if weekStart != "" && weekEnd != "" {
		startDate, endDate, err := getDateRange(weekStart, weekEnd)
		if err != nil {
			log.Println("error while getting date range: ", err)
			return nil, err
		}
		query = query.Where("check_in >= ? AND check_in < ?", startDate, endDate)
	}

	if sprint != "" {
		query = query.Where("sprint LIKE ?", fmt.Sprintf("%%%s%%", sprint))
	}

	if day != "" {
		query = query.Where("date LIKE ?", fmt.Sprintf("%%%s%%", day))
	}

	if owner != "" {
		query = query.Where("employee_name LIKE ?", fmt.Sprintf("%%%s%%", owner))
	}

	// Execute the query
	result := query.Find(&updates)
	if result.Error != nil {
		log.Println("error while querying database: ", result.Error)
	}

	return updates, result.Error
}

// getDateRange returns the start and end dates of a given date
func getDateRange(startDate, endDate string) (time.Time, time.Time, error) {

	startDateStr, err := time.Parse("2006-01-02 00:00:00", startDate)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid start date: %v", err)
	}

	endDateStr, err := time.Parse("2006-01-02 00:00:00", endDate)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid end date: %v", err)
	}

	return startDateStr, endDateStr, nil
}
