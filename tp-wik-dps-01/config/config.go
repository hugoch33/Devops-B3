package config

import "os"

func GetPort() string {
	port := os.Getenv("PING_LISTEN_PORT")

	if port == "" {
		port = "8080"
	}

	return port
}
