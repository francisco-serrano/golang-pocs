package custom_error

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type ErrFileNotFound struct {
	Filename string
	When     time.Time
}

func (e ErrFileNotFound) Error() string {
	return fmt.Sprintf("file %s was not found at %v", e.Filename, e.When)
}

func (e ErrFileNotFound) Is(other error) bool {
	_, ok := other.(ErrFileNotFound)
	return ok
}

var ErrNotExist = fmt.Errorf("file does not exist")
var ErrUserNotExist = errors.New("user does not exist")

type ErrFile struct {
	Filename string
	Base error
}

func (e ErrFile) Error() string {
	return fmt.Sprintf("file %s: %v", e.Filename, e.Base)
}

func (e ErrFile) Unwrap() error {
	return e.Base
}

func openFile(filename string) (string, error) {
	return "", ErrNotExist
}

func openFile2(filename string) (string, error) {
	return "", ErrFile{
		Filename: filename,
		Base:     ErrNotExist,
	}
}

func processFile(filename string) error {
	if _, err := openFile2(filename); err != nil {
		return fmt.Errorf("error while opening file: %w", err)
	}

	return nil
}

func assertErrors() {
	f, err := os.Open("test.txt")

	var pErr *os.PathError
	switch {
	case errors.Is(err, os.ErrPermission) && errors.As(err, &pErr):
		err = fmt.Errorf("you do not have permission to open the file: %w and the path is %s", err, pErr.Path)
		log.Println(err)
	case errors.Is(err, os.ErrNotExist) && errors.As(err, &pErr):
		err = fmt.Errorf("the file does not exist: %w and the path is %s", err, pErr.Path)
		log.Println(err)
	case errors.As(err, &pErr):
		err = fmt.Errorf("here is the original error %w and the path is %s", err, pErr.Path)
		log.Println(err)
	case err != nil:
		log.Println(err)
	}

	defer f.Close()
}

type writeFile struct {
	f *os.File
	err error
}

func (w *writeFile) WriteString(text string) {
	if w.err != nil {
		return
	}

	if _, err := io.WriteString(w.f, text); err != nil {
		w.err = err
	}
}

func (w *writeFile) Close() {
	if w.err != nil {
		return
	}

	if err := w.f.Close(); err != nil {
		w.err = err
	}
}

func (w *writeFile) Err() error {
	return w.err
}

func newWriteFile(filename string) *writeFile {
	f, err := os.Create(filename)

	return &writeFile{
		f:   f,
		err: err,
	}
}

func directFileExample() {
	f := newWriteFile("file.txt")
	f.WriteString("hello world")
	f.WriteString("more text")
	f.Close()

	if err := f.Err(); err != nil {
		panic(err)
	}
}

func foo() error {
	return fmt.Errorf("this error is from FOO - %w", bar())
}

func bar() error {
	return errors.New("this error is from BAR")
}

func Run() {
	err := foo()
	err2 := errors.Unwrap(err)

	fmt.Println(err)
	fmt.Println(err2)

	//directFileExample()

	//assertErrors()

	//if err := processFile("test.txt"); err != nil {
	//	var fErr ErrFile
	//	if errors.As(err, &fErr) {
	//		fmt.Printf("was unable to do something with file %s\n", fErr.Filename)
	//		fmt.Println("this is an ErrNotExist")
	//	}
	//
	//	fmt.Println(err)
	//}


	//_, err := openFile("test.txt")
	//if err != nil {
	//	wrappedErr := fmt.Errorf("unable to open file %v: %w", "test.txt", err)
	//	if errors.Is(wrappedErr, ErrNotExist) {
	//		fmt.Println("this is an ErrNotExist")
	//	}
	//
	//	fmt.Println(wrappedErr)
	//}
	//
	//_, err = openFile2("test.txt")
	//if err != nil {
	//	if errors.Is(err, ErrNotExist) {
	//		fmt.Println("this is an ErrNotExist")
	//	}
	//
	//	fmt.Println(err)
	//}
}
