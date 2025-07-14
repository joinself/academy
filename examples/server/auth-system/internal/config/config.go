// Package config provides configuration management for the Self authentication server.
package config

import (
	"encoding/base64"
	"encoding/hex"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joinself/academy/examples/server/auth-system/internal/auth"
	"github.com/joinself/academy/examples/server/auth-system/internal/server"
	"github.com/joinself/self-go-sdk/account"
)

// LoadAuthConfigFromEnv loads authentication configuration from environment variables
func LoadAuthConfigFromEnv() *auth.Config {
	// Start with default config
	config := auth.DefaultConfig()

	// StoragePath: SELF_AUTH_STORAGE_PATH
	if storagePath := os.Getenv("SELF_AUTH_STORAGE_PATH"); storagePath != "" {
		config.StoragePath = storagePath
	}

	// StorageKey: SELF_AUTH_STORAGE_KEY (hex or base64 encoded)
	if storageKeyEncoded := os.Getenv("SELF_AUTH_STORAGE_KEY"); storageKeyEncoded != "" {
		if storageKey, err := base64.StdEncoding.DecodeString(storageKeyEncoded); err == nil {
			config.StorageKey = storageKey
		}
	}

	// Environment: SELF_AUTH_ENVIRONMENT (sandbox only for now)
	if env := os.Getenv("SELF_AUTH_ENVIRONMENT"); env != "" {
		switch strings.ToLower(env) {
		case "sandbox":
			config.Environment = *account.TargetSandbox
		}
	}

	// LogLevel: SELF_AUTH_LOG_LEVEL (debug, info, warn, error)
	if logLevel := os.Getenv("SELF_AUTH_LOG_LEVEL"); logLevel != "" {
		switch strings.ToLower(logLevel) {
		case "debug":
			config.LogLevel = account.LogDebug
		case "info":
			config.LogLevel = account.LogInfo
		case "warn":
			config.LogLevel = account.LogWarn
		case "error":
			config.LogLevel = account.LogError
		}
	}

	// SessionTimeout: SELF_AUTH_SESSION_TIMEOUT (duration string, e.g., "1h", "30m")
	if sessionTimeout := os.Getenv("SELF_AUTH_SESSION_TIMEOUT"); sessionTimeout != "" {
		if duration, err := time.ParseDuration(sessionTimeout); err == nil {
			config.SessionTimeout = duration
		}
	}

	// QRCodeExpiration: SELF_AUTH_QR_EXPIRATION (duration string, e.g., "5m", "10m")
	if qrExpiration := os.Getenv("SELF_AUTH_QR_EXPIRATION"); qrExpiration != "" {
		if duration, err := time.ParseDuration(qrExpiration); err == nil {
			config.QRCodeExpiration = duration
		}
	}

	// RequiredClaims: SELF_AUTH_REQUIRED_CLAIMS (comma-separated, e.g., "liveness,email")
	if requiredClaims := os.Getenv("SELF_AUTH_REQUIRED_CLAIMS"); requiredClaims != "" {
		claims := strings.Split(requiredClaims, ",")
		for i, claim := range claims {
			claims[i] = strings.TrimSpace(claim)
		}
		config.RequiredClaims = claims
	}

	return config
}

// LoadServerConfigFromEnv loads server configuration from environment variables
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

	// SessionKey: SELF_SERVER_SESSION_KEY (hex or base64 encoded)
	if sessionKeyEncoded := os.Getenv("SELF_SERVER_SESSION_KEY"); sessionKeyEncoded != "" {
		// Try to decode as hex first
		if sessionKey, err := hex.DecodeString(sessionKeyEncoded); err == nil {
			config.SessionKey = sessionKey
		} else {
			// If hex fails, try base64
			if sessionKey, err := base64.StdEncoding.DecodeString(sessionKeyEncoded); err == nil {
				config.SessionKey = sessionKey
			}
			// If both fail, keep the default key
		}
	}

	// EnableTLS: SELF_SERVER_ENABLE_TLS
	if enableTLS := os.Getenv("SELF_SERVER_ENABLE_TLS"); enableTLS != "" {
		config.EnableTLS = strings.ToLower(enableTLS) == "true"
	}

	// TLSCertFile: SELF_SERVER_TLS_CERT_FILE
	if tlsCertFile := os.Getenv("SELF_SERVER_TLS_CERT_FILE"); tlsCertFile != "" {
		config.TLSCertFile = tlsCertFile
	}

	// TLSKeyFile: SELF_SERVER_TLS_KEY_FILE
	if tlsKeyFile := os.Getenv("SELF_SERVER_TLS_KEY_FILE"); tlsKeyFile != "" {
		config.TLSKeyFile = tlsKeyFile
	}

	// RequestTimeout: SELF_SERVER_REQUEST_TIMEOUT (duration string)
	if requestTimeout := os.Getenv("SELF_SERVER_REQUEST_TIMEOUT"); requestTimeout != "" {
		if duration, err := time.ParseDuration(requestTimeout); err == nil {
			config.RequestTimeout = duration
		}
	}

	// ShutdownTimeout: SELF_SERVER_SHUTDOWN_TIMEOUT (duration string)
	if shutdownTimeout := os.Getenv("SELF_SERVER_SHUTDOWN_TIMEOUT"); shutdownTimeout != "" {
		if duration, err := time.ParseDuration(shutdownTimeout); err == nil {
			config.ShutdownTimeout = duration
		}
	}

	return config
}
