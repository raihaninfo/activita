package helper

import (
	"crypto/rand"
	"log"

	"golang.org/x/crypto/bcrypt"
)

var Salt string = "someSecretPassActivita"

func GenerateHash(password, confirmPassword string) (string, error) {
	var hashP string
	if password == confirmPassword {
		HashPassword, err := bcrypt.GenerateFromPassword([]byte(password+Salt), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}
		hashP = string(HashPassword)
	}
	return hashP, nil
}

func ComparePassword(hash, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+Salt))
	if err != nil {
		return false, err
	}
	return true, nil
}

const otpChars = "1234567890"

func GenerateOTP(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}
