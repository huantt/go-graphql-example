package jwt

type Config struct {
	SecretKey string `json:"secret_key" mapstructure:"secret_key"`
}
