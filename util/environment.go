package util

import (
	"fmt"
	"os"
	"strconv"
)

type envKey struct {
	ListenPort string
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	DBSSLMode  string
}

// Env is a helper object for accessing environment variables.
var Env = &envKey{
	ListenPort: "REST_PORT",
	DBHost:     "DB_HOST",
	DBPort:     "DB_PORT",
	DBName:     "DB_DBNAME",
	DBUser:     "DB_USER",
	DBPassword: "DB_PASSWORD",
	DBSSLMode:  "DB_SSL_MODE",
}

func (g *envKey) GetListenPort() int {
	portStr := os.Getenv(g.ListenPort)
	if portStr == "" {
		msg := fmt.Sprintf("%s is not defined", g.ListenPort)
		panic(msg)
	}
	portNum, err := strconv.Atoi(portStr)
	if err != nil {
		msg := fmt.Sprintf("Invalid REST API port %s", portStr)
		panic(msg)
	}
	return portNum
}

func (g *envKey) GetDBHost() string {
	host := os.Getenv(g.DBHost)
	if host == "" {
		msg := fmt.Sprintf("%s is not defined", g.DBHost)
		panic(msg)
	}
	return host
}

func (g *envKey) GetDBPort() int {
	portStr := os.Getenv(g.DBPort)
	if portStr == "" {
		msg := fmt.Sprintf("%s is not defined", g.DBPort)
		panic(msg)
	}
	portNum, err := strconv.Atoi(portStr)
	if err != nil {
		msg := fmt.Sprintf("Invalid Postgres port %s", portStr)
		panic(msg)
	}
	return portNum
}

func (g *envKey) GetDBUser() string {
	v := os.Getenv(g.DBUser)
	if v == "" {
		msg := fmt.Sprintf("%s is not defined", g.DBUser)
		panic(msg)
	}
	return v
}

func (g *envKey) GetDBPassword() string {
	v := os.Getenv(g.DBPassword)
	if v == "" {
		msg := fmt.Sprintf("%s is not defined", g.DBPassword)
		panic(msg)
	}
	return v
}

func (g *envKey) GetDBName() string {
	v := os.Getenv(g.DBName)
	if v == "" {
		msg := fmt.Sprintf("%s is not defined", g.DBName)
		panic(msg)
	}
	return v
}

func (g *envKey) GetDBSSLMode() string {
	v := os.Getenv(g.DBSSLMode)
	if v == "" {
		defaultVal := "disable"
		return defaultVal
	}
	return v
}
