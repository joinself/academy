// Package config provides configuration management for the Self authentication server.
package config

import (
	"encoding/base64"
	"os"
	"strconv"

	"github.com/joinself/academy/examples/server/auth-system/internal/auth"
	"github.com/joinself/academy/examples/server/auth-system/internal/server"
)

// LoadAuthConfigFromEnv loads essential authentication configuration from environment variables
func LoadAuthConfigFromEnv() *auth.Config {
	// Start with default config
	config := auth.DefaultConfig()

	// StoragePath: SELF_AUTH_STORAGE_PATH
	if storagePath := os.Getenv("SELF_AUTH_STORAGE_PATH"); storagePath != "" {
		config.StoragePath = storagePath
	}

	// StorageKey: SELF_AUTH_STORAGE_KEY (base64 encoded)
	if storageKeyEncoded := os.Getenv("SELF_AUTH_STORAGE_KEY"); storageKeyEncoded != "" {
		if storageKey, err := base64.StdEncoding.DecodeString(storageKeyEncoded); err == nil {
			config.StorageKey = storageKey
		}
	}

	return config
}

// LoadServerConfigFromEnv loads essential server configuration from environment variables
func LoadServerConfigFromEnv() *server.Config {
	// Start with default config
	config := server.DefaultServerConfig()

	// Address: SELF_SERVER_ADDRESS
	if address := os.Getenv("SELF_SERVER_ADDRESS"); address != "" {
		config.Address = address
	}

	// Port: SELF_SERVER_PORT
	if portStr := os.Getenv("SELF_SERVER_PORT"); portStr != "" {
		if port, err := strconv.Atoi(portStr); err == nil {
			config.Port = port
		}
	}

	return config
}
