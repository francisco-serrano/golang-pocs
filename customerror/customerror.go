package customerror

import (
	"errors"
	"fmt"
	"os"
)

type CustomError struct {
	msg string
}

func NewCustomError(msg string) *CustomError {
	return &CustomError{msg: msg}
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("CUSTOM: %s", c.msg)
}

func Hola(param string) error {
	if param == "hola" {
		return NewCustomError("puto")
	}

	return nil
}

func Chau(param string) error {
	if param == "hola" {
		return errors.New("puto")
	}

	return nil
}

func Foo() error {
	var err *os.PathError = nil
	return err
}

func Run() {
	err := Foo()
	fmt.Println(err) // <nil>
}
