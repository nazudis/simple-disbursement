package config

const (
	ModeDev   = "dev"
	ModeProd  = "prod"
	ModeLocal = "local"
)

// Environment variables
const (
	Mode      = "MODE"
	LogFile   = "LOG_FILE"
	JWTSecret = "JWT_SECRET"

	// Database
	DBHost     = "DB_HOST"
	DBPort     = "DB_PORT"
	DBUser     = "DB_USER"
	DBPass     = "DB_PASS"
	DBName     = "DB_NAME"
	DBTimeZone = "DB_TIMEZONE"
)
