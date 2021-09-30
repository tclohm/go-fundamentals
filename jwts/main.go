package main

import (
	"fmt"
	"time"
	"net/http"
	"io/ioutil"

	jwt "github.com/dgrijalva/jwt-go"
)

// MARK: -- better: os.Get("JWT_SECRET_TOKEN")
var signingKey = []byte("secretPhrase")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "tclohm"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(signingKey)

	if err != nil {
		fmt.Printf("Something went wrong %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func main() {
	fmt.Println("Startup...")

	tokenString, err := GenerateJWT()
	if err != nil {
		fmt.Println("Error generating token string")
	}

	// w is not here, w is the WriterResponse

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:9000/", nil)
	req.Header.Set("Token", tokenString)
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf(w, err.Error())
	}

	fmt.Fprintf(w, string(body))

	fmt.Printf("token generated \n----\n%s\n----\n", tokenString)
	fmt.Println("Shutting down...")
}