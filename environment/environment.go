package environment

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type MyConfig struct {
	SampleVar  int    `env:"SAMPLE_VAR" env-default:"000" env-description:"sample var description" env-upd:""`
	AnotherVar string `env:"ANOTHER_VAR" env-default:"999" env-description:"another var description" env-upd:""`
}

var cfg MyConfig

func Run() {
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("error while reading variables: %v", err)
	}

	fmt.Println(cfg)

	desc, err := cleanenv.GetDescription(&cfg, nil)
	if err != nil {
		log.Fatalf("errow while reading env description: %v", err)
	}

	fmt.Println(desc)
}
