package postgresql

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connection gets connection of postgresql database
func Connection() (db *gorm.DB) {
	host := viper.Get("PGHOST")
	user := viper.Get("PGUSER")
	pass := viper.Get("PGPASSWORD")
	dsn := fmt.Sprintf("host=%v user=%v password=%v", host, user, pass)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
