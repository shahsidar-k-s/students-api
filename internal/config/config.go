package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address" env-required:"true"`
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {

	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("setEnv", "", "Sets the env variables")
		flag.Parse()
		configPath = *flags
		if configPath == "" {
			log.Fatal("No envs provided please provide the configPath")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {

		log.Fatal("NOt config path exist please provide the config path")
	}
	var cfg Config
	error := cleanenv.ReadConfig(configPath, &cfg)
	if error != nil {
		log.Fatal("Faild to read the env file please try again  ")
	}
	return &cfg

}
