package main

// Key generated with :
// openssl genrsa -out key.pem 2048
// openssl rsa -in key.pem -outform PEM -pubout -out public.pem

// Valid token
// eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.jFC5vYe50WBFsCyx5syiVHWbS1cqmt1PPW8NQuYJTR9VVxKxeT8hZ7lDWDgjyvKTL3h3ZL9w8xLi6onfea64uZOz5Z53GlP6aS2Atevu-AMbtL5Jw17V_kcPjBqBqCg6G2qQfEm3astUbKMXfrbakNMUuHj0P5z8WjXBP7pp8VQin9ixH5zP7wZLxO_4E1bgXWY8ggeGIn5coaiQh8KJ0SwAN8cnh4PfOw8O5McIEYou6UH5w1UjJzKzA_TuZJReJ8RvcrKp4b6MQqAP8ob2INXn5rZ_qZmgI7OAlUuaDIvvLQtOWEzEfAU12dun81qyNlA8RdAkWQnG4vr0iFv3LA

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

var rsaPKey *rsa.PublicKey

func authentication(w http.ResponseWriter, req *http.Request) {
	bearer := req.Header.Get("Authorization")
	if !strings.Contains(bearer, "Bearer ") {
		w.WriteHeader(401)
		return
	}

	token := strings.ReplaceAll(bearer, "Bearer ", "")

	claim := &CustomClaim{}
	validatedToken, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "RS256" {
			return nil, fmt.Errorf("wrong signing method : %s", t.Method.Alg())
		}

		return rsaPKey, nil
	})

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(401)
		return
	}
	if !validatedToken.Valid {
		fmt.Println("Token not valid")
		w.WriteHeader(401)
		return
	}

	payload, err := json.Marshal(claim)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(payload)
}

func main() {
	data, err := os.ReadFile("./public.pem")
	if err != nil {
		panic(err)
	}

	rsaPKey, err = jwt.ParseRSAPublicKeyFromPEM(data)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/authentication", authentication)
	fmt.Println("listening on 8090")
	http.ListenAndServe(":8090", nil)
}

type CustomClaim struct {
	jwt.RegisteredClaims
	Name string `json:"name"`
}
