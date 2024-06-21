package main

import (
	"june/botmgr/internal/api"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func initLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func initConig() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AddConfigPath(home)
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func init() {
	initConig()
	initLogger()

	log.Info().Msg("manager initialized...")
}

func main() {
	log.Info().Msg("starting API...")
	if err := api.New().Start(); err != nil {
		panic(err)
	}
}
