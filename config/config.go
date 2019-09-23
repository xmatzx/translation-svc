package config

// should have here default const of system
const (
	DbPort               = "3306"
	DbHost               = "localhost"
	DbName               = "translate"
	DbUser               = "root"
	DbPassword           = "root"
	LogLevel             = "info"
	Environment          = "staging"
	DbMaxOpenConnections = 100
	DbMaxIdleConnections = 0
	DbDebug              = "true"
)
