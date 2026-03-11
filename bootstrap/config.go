package bootstrap

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Env struct {
	AppEnv         string `yaml:"APP_ENV"`
	ServerAddress  string `yaml:"SERVER_ADDRESS"`
	ContextTimeout int    `yaml:"CONTEXT_TIMEOUT"`
	APISecret      string `yaml:"API_SECRET"`
}

func NewConfig() *Env {
	env := Env{}

	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("Can't find the file config.yaml : ", err)
	}

	err = yaml.Unmarshal(data, &env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.APISecret == "" {
		log.Fatal("API_SECRET must be set in config.yaml")
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
