package main

import (
	"github.com/resotto/goilerplate/internal/app/adapter"
	"github.com/spf13/viper"
)

func main() {
	r := adapter.Router()
	viper.SetDefault("BITBANK_HOST", "https://public.bitbank.cc")
	// for local development
	viper.SetDefault("PGHOST", "0.0.0.0")
	viper.SetDefault("PGUSER", "postgres")
	viper.SetDefault("PGPASSWORD", "postgres")

	viper.BindEnv("PGHOST", "PGHOST")
	viper.BindEnv("PGUSER", "PGUSER")
	viper.BindEnv("PGPASSWORD", "PGPASSWORD")
	r.Run(":8080")
}
