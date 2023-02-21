package config

import (
	"log"

	"github.com/Netflix/go-env"
)

type Environment struct {
	Port   int  `env:"PORT,default=8000"`
	Debug  bool `env:"DEBUG,default=false"`
	Extras env.EnvSet
}

var Env Environment

func init() {
	_, err := env.UnmarshalFromEnviron(&Env)
	if err != nil {
		log.Fatal(err)
	}
}

func GetEnv() Environment {
	return Env
}
