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

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
