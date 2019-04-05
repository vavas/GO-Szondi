package main

import (
	"os"
	"strconv"
	"strings"
)

var (
	gEnvironment *Environment
)

// Environment configuration for aggregation app
type Environment struct {
	TestMode         bool
	CorsOrigins      []string
	PlaylistDSN      string
}

func setConfig(env Environment) Environment {
	gEnvironment = &env
	return *gEnvironment
}

func readConfig() *Environment {
	tm, _ := strconv.ParseBool(os.Getenv("ENABLE_TEST_MODE"))

	return &Environment{
		TestMode:         tm,
		CorsOrigins:      strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ","),
		PlaylistDSN:      os.Getenv("PLAYLIST_RPC_DSN"),
	}
}

func getConfig() Environment {
	if gEnvironment == nil {
		gEnvironment = readConfig()
	}

	return *gEnvironment
}
