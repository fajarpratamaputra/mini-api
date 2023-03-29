package module

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"interaction-api/config"
)

// Config stores the application configurations.
type Config struct {
	Host        string `envconfig:"APP_HOST"`
	Port        string `envconfig:"PORT" default:"8080"`
	JWTSecret   string `envconfig:"JWT_SECRET" default:"secret"`
	BasePathURL string `envconfig:"BASE_PATH_URL" default:"/api/"`
	DB          map[string]config.DbConfig
	Redis       config.RedisConfig
	CacheConfig config.CacheConfig
	KafkaConfig config.KafkaConfig
	MongoConfig config.MongoConfig
	JwtRedis    config.JwtRedisConfig
	ImagePath   string `envconfig:"IMAGE_PATH"`
	VideoPath   string `envconfig:"VIDEO_PATH"`
	AdsLink     string `envconfig:"ADS_LINK"`
}

type DB struct {
	DBConfig
}

type DBConfig struct {
	User       string `envconfig:"DB_DATABASE_USER" required:"true"`
	Host       string `envconfig:"DB_DATABASE_HOST" required:"true"`
	Password   string `envconfig:"DB_DATABASE_PASSWORD" required:"true"`
	DB         string `envconfig:"DB_DATABASE_DB" required:"true"`
	DBLogLevel string `envconfig:"DB_LOG_LEVEL" default:"INFO"`
}

type RedisConfig struct {
	Host     string `envconfig:"REDIS_HOST"`
	Port     string `envconfig:"REDIS_PORT" default:"6379"`
	Password string `envconfig:"REDIS_PASSWORD"`
	DB       int    `envconfig:"REDIS_DB"`
}

type JwtRedisConfig struct {
	Enable   bool   `envconfig:"JWT_REDIS_ENABLE" default:"false"`
	Host     string `envconfig:"JWT_REDIS_HOST"`
	Port     string `envconfig:"JWT_REDIS_PORT" default:"6379"`
	Password string `envconfig:"JWT_REDIS_PASSWORD"`
	DB       int    `envconfig:"JWT_REDIS_DB"`
	Key      string `envconfig:"JWT_REDIS_KEY"`
}

type KafkaConfig struct {
	Brokers string `envconfig:"KAFKA_BROKERS"`
}

type MongoConfig struct {
	Host string `envconfig:"MONGO_HOST"`
	Port string `envconfig:"MONGO_PORT"`
	DB   string `envconfig:"MONGO_DB"`
	User string `envconfig:"MONGO_USER"`
	Pass string `envconfig:"MONGO_PASS"`
}

type CacheConfig struct {
	EnableCache         bool   `envconfig:"ENABLE_CACHE" required:"true"`
	CachePrefix         string `envconfig:"CACHE_PREFIX" default:"homepage-api"`
	CacheDefaultTimeOut int    `envconfig:"DEFAULT_CACHE_TIMEOUT" default:"300"`
}

func init() {
	_ = godotenv.Overload()
	var configLoader Config
	var dbConfigLoader DB
	var redisConfigLoader RedisConfig
	var JwtRedisConfigLoader JwtRedisConfig
	var cacheConfigLoader CacheConfig
	var kafkaConfigLoader KafkaConfig
	var mongoConfigLoader MongoConfig

	// load core config
	if err := envconfig.Process("", &configLoader); err != nil {
		panic(err)
	}
	// load database config
	if err := envconfig.Process("", &dbConfigLoader); err != nil {
		panic(err)
	}
	// load cache config
	if err := envconfig.Process("", &cacheConfigLoader); err != nil {
		panic(err)
	}
	// load redis config
	if err := envconfig.Process("", &redisConfigLoader); err != nil {
		panic(err)
	}
	// load redis config
	if err := envconfig.Process("", &JwtRedisConfigLoader); err != nil {
		panic(err)
	}
	if err := envconfig.Process("", &kafkaConfigLoader); err != nil {
		panic(err)
	}
	if err := envconfig.Process("", &mongoConfigLoader); err != nil {
		panic(err)
	}

	cfg := config.Config(configLoader)
	cfg.DB = map[string]config.DbConfig{}
	cfg.DB["db"] = config.DbConfig(dbConfigLoader.DBConfig)
	cfg.CacheConfig = config.CacheConfig(cacheConfigLoader)
	cfg.Redis = config.RedisConfig(redisConfigLoader)
	cfg.KafkaConfig = config.KafkaConfig(kafkaConfigLoader)
	cfg.JwtRedis = config.JwtRedisConfig(JwtRedisConfigLoader)
	cfg.MongoConfig = config.MongoConfig(mongoConfigLoader)

	config.AppConfig = &cfg
}
