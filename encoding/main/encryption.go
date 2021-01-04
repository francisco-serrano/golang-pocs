package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type myClaims struct {
	jwt.StandardClaims
	Email string
}

const myKey = "my-key"

func Run() {
	msg := "hello world"

	passwd := "ilovedogs"
	bs, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.MinCost)
	if err != nil {
		log.Fatal("could not bcrypt password", err)
	}

	bs = bs[:16]

	w := &bytes.Buffer{}
	encWriter, err := encryptWriter(w, bs)
	if err != nil {
		log.Fatal("could not encrypt writer", err)
	}

	if _, err := io.WriteString(encWriter, msg); err != nil {
		log.Fatal("could not write string", err)
	}

	encrypted := w.String()
	fmt.Println(encrypted)

	fmt.Println("before base64", encrypted)

	result2, err := enDecode(bs, encrypted)
	if err != nil {
		log.Fatal("could not enDecode", err)
	}

	fmt.Println(string(result2))
}

func RunWithSHA() {
	f, err := os.Open("sample-file.txt")
	if err != nil {
		log.Fatal("could not open file", err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal("could not close file", err)
		}
	}()

	h := sha256.New()

	if _, err := io.Copy(h, f); err != nil {
		log.Fatal("could not copy content", err)
	}

	xb := h.Sum(nil)

	fmt.Printf("%T\n", h)
	fmt.Printf("%v\n", h)
	fmt.Printf("%T\n", xb)
	fmt.Printf("%x\n\n", xb)

	xb = h.Sum(nil)

	fmt.Printf("%T\n", xb)
	fmt.Printf("%x\n\n", xb)

	xb = h.Sum(xb)

	fmt.Printf("%T\n", xb)
	fmt.Printf("%x\n\n", xb)
}

func RunWithHMAC() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/submit", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		c = &http.Cookie{}
	}

	signedToken := c.Value
	parsedToken, err := jwt.ParseWithClaims(signedToken, &myClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() == jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("someone tried to hack: changed sining method")
		}

		return []byte(myKey), nil
	})

	isEqual := err == nil && parsedToken.Valid

	message := "not logged in"
	if isEqual {
		message = "logged in"

		claims := parsedToken.Claims.(*myClaims)
		fmt.Println(claims.Email)
		fmt.Println(claims.ExpiresAt)
	}

	html := `
		<!DOCTYPE html>
		<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<meta http-equiv="X-UA-Compatible" content="ie=edge">
				<title>HMAC Example</title>
			</head>
			<body>
				<p>Cookie value: ` + c.Value + `</p>
				<p>` + message + `</p>
				<form action="/submit" method="post">
					<input type="email" name="email" />
					<input type="submit" />
				</form>
			</body>
		</html>`

	io.WriteString(w, html)
}

func bar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	email := r.FormValue("email")
	if email == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	token, err := getJWT(email)
	if err != nil {
		http.Error(w, fmt.Errorf("could not get JWT %w", err).Error(), http.StatusInternalServerError)
		return
	}

	// hash / message digest / digest / hash value | "what we stored"
	c := &http.Cookie{
		Name:  "session",
		Value: token,
	}

	http.SetCookie(w, c)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func getJWT(msg string) (string, error) {
	claims := &myClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
		Email: msg,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(myKey))
	if err != nil {
		return "", fmt.Errorf("error while signing JWT %w", err)
	}

	return signedToken, nil
}

func enDecode(key []byte, input string) ([]byte, error) {
	b, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("could not newCipher %w", err)
	}

	// initialization vector
	iv := make([]byte, aes.BlockSize)

	// create cipher
	s := cipher.NewCTR(b, iv)

	buff := &bytes.Buffer{}
	sw := cipher.StreamWriter{
		S: s,
		W: buff,
	}

	if _, err := sw.Write([]byte(input)); err != nil {
		return nil, fmt.Errorf("could not sw.Write to streamWriter %w", err)
	}

	return buff.Bytes(), nil
}

func encryptWriter(w io.Writer, key []byte) (io.Writer, error) {
	b, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("could not newCipher %w", err)
	}

	// initialization vector
	iv := make([]byte, aes.BlockSize)

	// create cipher
	s := cipher.NewCTR(b, iv)

	return cipher.StreamWriter{
		S: s,
		W: w,
	}, nil
}

func main() {
	RunWithHMAC()
}