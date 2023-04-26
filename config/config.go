package config

import (
	"bytes"
	"encoding/json"
	"github.com/huantt/go-graphql-sample/pkg/graphql"
	"github.com/huantt/go-graphql-sample/pkg/log"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type Config struct {
	Log            log.Config     `json:"log" mapstructure:"log"`
	Port           uint16         `json:"port" mapstructure:"port"`
	AllowedOrigins []string       `json:"allowed_origins" mapstructure:"allowed_origins"`
	Graphql        graphql.Config `json:"graphql" mapstructure:"graphql"`
}

var cfg *Config

func Load() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}
	// You should set default config value here
	c := &Config{
		Log: log.Config{
			Level:  "debug",
			Format: "text",
		},
		Port: 9000,
	}

	// --- hacking to load reflect structure config into env ----//
	viper.SetConfigType("json")
	configBuffer, err := json.Marshal(c)

	if err != nil {
		return nil, err
	}

	if err := viper.ReadConfig(bytes.NewBuffer(configBuffer)); err != nil {
		panic(err)
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if os.Getenv("GENERATE_ENV_TEMPLATE") == "true" {
		err = viper.WriteConfigAs(".init.env")
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to write env template file")
		}
	}

	// -- end of hacking --//
	viper.AutomaticEnv()
	err = viper.Unmarshal(c)
	if err != nil {
		return nil, err
	}
	cfg = c
	return c, nil
}
