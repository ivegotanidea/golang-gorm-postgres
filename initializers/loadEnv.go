package initializers

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	S3AccessKey    string `mapstructure:"S3_ACCESS_KEY"`
	S3AccessSecret string `mapstructure:"S3_ACCESS_SECRET"`
	S3Bucket       string `mapstructure:"S3_IMG_UPLOAD_BUCKET"`
	S3Endpoint     string `mapstructure:"S3_RESOURCE_ENDPOINT"`
	S3Region       string `mapstructure:"S3_RESOURCE_REGION"`

	ImgProxyBaseUrl           string `mapstructure:"IMG_PROXY_BASE_URL"`
	ImgProxySigningHexKey     string `mapstructure:"IMG_PROXY_SIGNING_HEX_KEY"`
	ImgProxySigningSaltHex    string `mapstructure:"IMG_PROXY_SIGNING_SALT_HEX"`
	ProcessingGoroutinesCount int    `mapstructure:"PROCESSING_GOROUTINES_COUNT"`

	CasbinModelPath  string `mapstructure:"CASBIN_MODEL_PATH"`
	CasbinPolicyPath string `mapstructure:"CASBIN_POLICY_PATH"`

	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	ServerPort     string `mapstructure:"PORT"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`

	AccessTokenPrivateKey  string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string        `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string        `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn  time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge      int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenMaxAge     int           `mapstructure:"REFRESH_TOKEN_MAXAGE"`

	VerifiedDistanceThreshold int `mapstructure:"VERIFIED_DISTANCE_THRESHOLD"`
	ReviewUpdateLimitHours    int `mapstructure:"REVIEW_UPDATE_LIMIT_HOURS"`

	ParsedBaseUrl string `mapstructure:"PARSED_BASE_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
