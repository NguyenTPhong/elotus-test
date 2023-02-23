package config

// all global environment variable read from ENV
var (
	Environment   = GetString("ENVIRONMENT", "development")                                                                       // environment
	Port          = GetInt64("APP_PORT", 80)                                                                                      // app port
	DbConnStr     = GetString("DB_CONN_STR", "host=postgres port=5432 user=elotus password=elotus dbname=elotus sslmode=disable") // postgres connection string
	JWTKey        = GetString("JWT_KEY", "development-key")
	TokenLifeTime = GetInt64("TOKEN_LIFE_TIME", 2)        // jwt token life time
	StoragePath   = GetString("STORAGE_PATH", "/upload")  // upload path
	RedisHost     = GetString("REDIS_HOST", "redis")      // redis host
	RedisPort     = GetString("REDIS_PORT", "6379")       // redis host
	RedisPassword = GetString("REDIS_PASSWORD", "elotus") // redis pw
	LogLevel      = GetInt64("LOG_LEVEL", -1)
)
