package encryption

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
)

var key []byte

func init() {
	for i := 0; i < 64; i++ {
		key = append(key, byte(i))
	}
}

func Run() {
	hmacFunc()
}

func hmacFunc() {
	h := sha256.New()
	if _, err := io.WriteString(h, "hello world"); err != nil {
		log.Fatal(err)
	}
	s := h.Sum(nil)
	fmt.Println(hex.EncodeToString(s))

	h1 := sha256.New()
	if _, err := io.WriteString(h1, "hello world"); err != nil {
		log.Fatal(err)
	}
	s1 := h1.Sum(nil)
	fmt.Println(hex.EncodeToString(s1))

	h2 := sha256.New()
	h2.Write([]byte("hello world"))
	s2 := h2.Sum([]byte{})
	fmt.Println(s2)
	fmt.Printf("%x\n", s2)
}

func baseEncoding() {
	v := "hello world"

	encoded := base64.StdEncoding.EncodeToString([]byte(v))

	fmt.Println(encoded)

	decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		log.Fatal(err)
	}

	decoded := string(decodedBytes)

	fmt.Println(decoded)
}

func hash() {
	pass := "123456789"

	hashedPass, err := hashPassword(pass)
	if err != nil {
		log.Fatal(err)
	}

	if err := comparePassword(pass, hashedPass); err != nil {
		log.Fatal("not logged in")
	}

	log.Println("logged in!")
}

func hashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error while generating bcrypt hash from password: %w", err)
	}

	return bs, nil
}

func comparePassword(password string, hashedPass []byte) error {
	if err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password)); err != nil {
		return fmt.Errorf("invalid password: %w", err)
	}

	return nil
}
