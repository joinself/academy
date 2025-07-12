// Package logging provides structured logging configuration for the Self authentication system.
package logging

import (
	"context"
	"log/slog"
	"os"
	"strings"
	"time"
)

// LogLevel represents the logging level
type LogLevel string

const (
	LevelDebug LogLevel = "debug"
	LevelInfo  LogLevel = "info"
	LevelWarn  LogLevel = "warn"
	LevelError LogLevel = "error"
)

// LogFormat represents the logging format
type LogFormat string

const (
	FormatJSON LogFormat = "json"
	FormatText LogFormat = "text"
)

// Config holds logging configuration
type Config struct {
	Level      LogLevel
	Format     LogFormat
	Service    string
	Version    string
	AddSource  bool
	TimeFormat string
}

// DefaultConfig returns a default logging configuration
func DefaultConfig() *Config {
	return &Config{
		Level:      LevelInfo,
		Format:     FormatJSON,
		Service:    "self-auth-system",
		Version:    "1.0.0",
		AddSource:  false,
		TimeFormat: time.RFC3339,
	}
}

// DevelopmentConfig returns logging configuration optimized for development
func DevelopmentConfig() *Config {
	return &Config{
		Level:      LevelDebug,
		Format:     FormatText,
		Service:    "self-auth-system",
		Version:    "1.0.0-dev",
		AddSource:  true,
		TimeFormat: "15:04:05",
	}
}

// ProductionConfig returns logging configuration optimized for production
func ProductionConfig() *Config {
	return &Config{
		Level:      LevelInfo,
		Format:     FormatJSON,
		Service:    "self-auth-system",
		Version:    "1.0.0",
		AddSource:  false,
		TimeFormat: time.RFC3339,
	}
}

// Setup initializes the global logger with the provided configuration
func Setup(config *Config) *slog.Logger {
	// Parse log level
	var level slog.Level
	switch strings.ToLower(string(config.Level)) {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	// Create handler options
	opts := &slog.HandlerOptions{
		Level:     level,
		AddSource: config.AddSource,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Customize time format
			if a.Key == slog.TimeKey {
				return slog.Time(slog.TimeKey, a.Value.Time().UTC())
			}
			return a
		},
	}

	// Create handler based on format
	var handler slog.Handler
	switch config.Format {
	case FormatJSON:
		handler = slog.NewJSONHandler(os.Stdout, opts)
	case FormatText:
		handler = slog.NewTextHandler(os.Stdout, opts)
	default:
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	// Create logger with service context
	logger := slog.New(handler).With(
		slog.String("service", config.Service),
		slog.String("version", config.Version),
	)

	// Set as default logger
	slog.SetDefault(logger)

	return logger
}

// FromEnv creates a logging configuration from environment variables
func FromEnv() *Config {
	config := DefaultConfig()

	// Check for development environment
	if env := os.Getenv("APP_ENV"); env == "development" || env == "dev" {
		config = DevelopmentConfig()
	}

	// Override with specific environment variables
	if level := os.Getenv("LOG_LEVEL"); level != "" {
		config.Level = LogLevel(strings.ToLower(level))
	}

	if format := os.Getenv("LOG_FORMAT"); format != "" {
		config.Format = LogFormat(strings.ToLower(format))
	}

	if service := os.Getenv("SERVICE_NAME"); service != "" {
		config.Service = service
	}

	if version := os.Getenv("SERVICE_VERSION"); version != "" {
		config.Version = version
	}

	return config
}

// Logger wraps slog.Logger with additional context methods
type Logger struct {
	*slog.Logger
}

// New creates a new Logger instance
func New(logger *slog.Logger) *Logger {
	return &Logger{Logger: logger}
}

// WithContext adds context fields to the logger
func (l *Logger) WithContext(ctx context.Context) *Logger {
	// Extract relevant context values
	logger := l.Logger

	if userDID := ctx.Value("user_did"); userDID != nil {
		logger = logger.With(slog.String("user_did", userDID.(string)))
	}

	if sessionID := ctx.Value("session_id"); sessionID != nil {
		logger = logger.With(slog.String("session_id", sessionID.(string)))
	}

	if requestID := ctx.Value("request_id"); requestID != nil {
		logger = logger.With(slog.String("request_id", requestID.(string)))
	}

	return &Logger{Logger: logger}
}

// WithComponent adds component information to the logger
func (l *Logger) WithComponent(component string) *Logger {
	return &Logger{Logger: l.Logger.With(slog.String("component", component))}
}

// WithUser adds user information to the logger
func (l *Logger) WithUser(userDID string) *Logger {
	return &Logger{Logger: l.Logger.With(slog.String("user_did", userDID))}
}

// WithRequest adds HTTP request information to the logger
func (l *Logger) WithRequest(method, path string, statusCode int, duration time.Duration) *Logger {
	return &Logger{Logger: l.Logger.With(
		slog.String("method", method),
		slog.String("path", path),
		slog.Int("status_code", statusCode),
		slog.Duration("duration", duration),
	)}
}

// LogAuthRequest logs authentication request events
func (l *Logger) LogAuthRequest(requestID, contentID string, requiredClaims []string) {
	l.Logger.Info("Authentication request generated",
		slog.String("request_id", requestID),
		slog.String("content_id", contentID),
		slog.Any("required_claims", requiredClaims),
	)
}

// LogConnection logs connection events
func (l *Logger) LogConnection(event string, connectionID, userDID, contentID string) {
	l.Logger.Info("Connection event",
		slog.String("event", event),
		slog.String("connection_id", connectionID),
		slog.String("user_did", userDID),
		slog.String("content_id", contentID),
	)
}

// LogAuthentication logs authentication completion events
func (l *Logger) LogAuthentication(success bool, userDID, sessionID string, claims map[string]interface{}) {
	if success {
		l.Logger.Info("User authentication successful",
			slog.String("user_did", userDID),
			slog.String("session_id", sessionID),
			slog.Any("claims", claims),
		)
	} else {
		l.Logger.Warn("User authentication failed",
			slog.String("user_did", userDID),
		)
	}
}

// LogSessionEvent logs session-related events
func (l *Logger) LogSessionEvent(event, sessionID, userDID string) {
	l.Logger.Info("Session event",
		slog.String("event", event),
		slog.String("session_id", sessionID),
		slog.String("user_did", userDID),
	)
}

// LogError logs error events with structured context
func (l *Logger) LogError(message string, err error, fields ...slog.Attr) {
	attrs := []slog.Attr{slog.String("error", err.Error())}
	attrs = append(attrs, fields...)
	l.Logger.LogAttrs(context.Background(), slog.LevelError, message, attrs...)
}

// LogServiceEvent logs service lifecycle events
func (l *Logger) LogServiceEvent(event, component string, details map[string]interface{}) {
	l.Logger.Info("Service event",
		slog.String("event", event),
		slog.String("component", component),
		slog.Any("details", details),
	)
}
