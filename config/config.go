package config

// Config stores the application configurations.
type Config struct {
	Host        string
	Port        string
	JWTSecret   string
	BasePathURL string
	DB          map[string]DbConfig
	Redis       RedisConfig
	CacheConfig CacheConfig
	KafkaConfig KafkaConfig
	MongoConfig MongoConfig
	JwtRedis    JwtRedisConfig
	ImagePath   string
	VideoPath   string
	AdsLink     string
}

type DbConfig struct {
	User       string
	Host       string
	Password   string
	DB         string
	DBLogLevel string
}

type CacheConfig struct {
	EnableCache         bool
	CachePrefix         string
	CacheDefaultTimeOut int
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type JwtRedisConfig struct {
	Enable   bool
	Host     string
	Port     string
	Password string
	DB       int
	Key      string
}

type KafkaConfig struct {
	Brokers string
}

type MongoConfig struct {
	Host string
	Port string
	DB   string
	User string
	Pass string
}

var AppConfig *Config
