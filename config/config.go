package config

// all global environment variable read from ENV
var (
	Environment   = GetString("ENVIRONMENT", "development")                                                                       // environment
	Port          = GetInt64("APP_PORT", 80)                                                                                      // app port
	DbConnStr     = GetString("DB_CONN_STR", "host=postgres port=5432 user=elotus password=elotus dbname=elotus sslmode=disable") // postgres connection string
	DbMaxConn     = GetInt64("DB_MAX_CONN", 10)                                                                                   // max connection to db
	DbMaxIdleConn = GetInt64("DB_MAX_IDLE_CONN", 2)                                                                               // max idle connection to db
	DBLogLevel    = GetInt64("DB_LOG_LEVEL", 4)                                                                                   // db log level
	JWTKey        = GetString("JWT_KEY", "development-key")
	TokenLifeTime = GetInt64("TOKEN_LIFE_TIME", 2)        // jwt token life time
	StoragePath   = GetString("STORAGE_PATH", "/upload")  // upload path
	RedisHost     = GetString("REDIS_HOST", "redis")      // redis host
	RedisPort     = GetString("REDIS_PORT", "6379")       // redis host
	RedisPassword = GetString("REDIS_PASSWORD", "elotus") // redis pw
	LogLevel      = GetInt64("LOG_LEVEL", -1)
)
