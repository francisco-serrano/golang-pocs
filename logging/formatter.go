package logging

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"os"
)



func Run() {
	logger := &logrus.Logger{
		Out:   os.Stdout,
		Level: logrus.InfoLevel,
		Formatter: &easy.Formatter{
			LogFormat: "[%lvl%]: %time% - field_A={%field_A%} - %msg%",
		},
	}

	logger.WithField("field_A", "value A").
		WithField("field_B", "value B").
		Info("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	logger.WithField("field_A", "value A").
		WithField("field_B", "value B").
		Info("aaaaa\n")

	aux := map[string]interface{}{}

	aux["hello"] = 10
	aux["world"] = 49
	aux["how"] = 23
	aux["are"] = 12
	aux["you"] = 71

	fmt.Println(aux)

	bytes, err := json.Marshal(aux)
	if err != nil {
		panic(err)
	}

	fmt.Println(bytes)
}
