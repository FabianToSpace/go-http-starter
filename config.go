package main

import (
	"os"
	"strconv"
)

var (
	listenHostName   string = GetEnv("LISTENHOSTNAME", "127.0.0.1")
	hostPort         string = GetEnv("PORT", "8887")
	postgresUsername string = GetEnv("POSTGRES_USER", "postgres")
	postgresPassword string = GetEnv("POSTGRES_PASSWORD", "postgres")
	postgresServer   string = GetEnv("POSTGRES_SERVER", "localhost")
	postgresPort     string = GetEnv("POSTGRES_PORT", "5432")
	postgresDatabase string = GetEnv("POSTGRES_DATABASE", "postgres")
	postgresSslMode  string = GetEnv("POSTGRES_SSLMODE", "disable")
)

func GetEnv(key, fallback string) string {
	v := os.Getenv(key)

	if len(v) == 0 {
		return fallback
	}

	return v
}

func GetEnvBool(key string, fallback bool) bool {
	v := GetEnv(key, strconv.FormatBool(fallback))

	if len(v) == 0 {
		return fallback
	}

	b, err := strconv.ParseBool(v)

	if err != nil {
		panic(err)
	}

	return b
}
