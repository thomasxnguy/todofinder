package app

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/go-ozzo/ozzo-validation"
	"fmt"
)

type Configuration struct {
	// Network value for used by fastHttp.
	Network string `yaml:"network"`
	// Server listening port.
	ListenOn string `yaml:"listenOn"`
	// If set to true, server will run in TSL mode, providing secure http traffic.
	EnableTls bool `yaml:"enableTls"`
	// Path to SLL certificate private key
	KeyFile string `yaml:"keyFile"`
	// Path to SLL signed certificate
	CertFile string `yaml:"certFile"`
	// Define logger output.
	LogOutput string `yaml:"logOutput"`
	// Location of the log file.
	LogFile string `yaml:"logFile"`
	//Logger level output.
	LogLevel string `yaml:"logLevel"`
}

// Validate validates the mandatory values in the configuration file.
func (config Configuration) Validate() error {
	var err error
	if config.EnableTls {
		err = validation.ValidateStruct(&config,
			validation.Field(&config.KeyFile, validation.Required),
			validation.Field(&config.CertFile, validation.Required),
		)
		if err != nil {
			return err
		}
	}
	if config.LogOutput == "file" {
		err = validation.ValidateStruct(&config,
			validation.Field(&config.LogFile, validation.Required),
		)
		if err != nil {
			return err
		}
	}
	return nil
}

// LoadConfiguration load server configuration yaml file determined by filepath.
// Viper also have the ability to read from env variables.
// Environment variables with the prefix "TODOFINDER_" in their names are also read automatically.
func LoadConfiguration(filepath *string) (*Configuration, error) {
	v := viper.New()
	v.SetConfigFile(*filepath)
	v.SetConfigType("yaml")
	v.SetEnvPrefix("TODOFINDER")
	v.AutomaticEnv()
	logrus.Debugf("Loading configuration from %v", v.ConfigFileUsed())

	v.SetDefault("network", "tcp4")
	v.SetDefault("listenOn", ":8080")
	v.SetDefault("enableTls", "true")
	v.SetDefault("logOutput", "discard")
	v.SetDefault("logLevel", "info")

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read the configuration file: %s", err)
	}
	config := &Configuration{}
	if err := v.Unmarshal(config); err != nil {
		return nil, err
	}
	return config, config.Validate()
}
