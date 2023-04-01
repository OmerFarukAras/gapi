package util

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"gapi/models"
	"gapi/util/db"
	"io"
	"log"
	"net/http"
	"strings"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

func Write(w Writer, s string) {
	_, err := io.WriteString(w, s)
	if err != nil {
		return
	}
}
func ShaHash(input string) string {
	plainText := []byte(input)
	sha256Hash := sha256.Sum256(plainText)
	return hex.EncodeToString(sha256Hash[:])
}
func GetUser(req *http.Request) *models.User {
	tokenString := req.Header.Get("Authorization")

	if len(tokenString) > 1 {
		tokenString = strings.Split(tokenString, " ")[1]
		if len(tokenString) > 1 {
			token, ok := ParseToken(tokenString)
			if ok {
				user, ok := db.FindUserByEmail(token.Email)
				if ok {
					return user
				} else {
					return nil
				}
			} else {
				return nil
			}
		}
	}
	return nil
}

func PostsArrayToString(items []models.Post) (string, error) {
	var buffer bytes.Buffer
	var err error
	var b []byte

	for _, item := range items {
		b, err = json.Marshal(item)
		if err != nil {
			return "", err
		}
		// use space to separate each json string in the array
		buffer.WriteString(string(b) + ",")
	}

	s := strings.TrimSpace("[" + buffer.String() + "]")

	return s, nil
}

func JsonWrite(res Writer, data map[string]string) {
	jsonResp, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	_, err = res.Write(jsonResp)
	if err != nil {
		return
	}
}
