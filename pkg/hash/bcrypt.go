package hash

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

const (
	numbers   = "12345678901234567890"
	alphanums = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234567890"
)

func ToBcrypt(plainText string) string {
	password := []byte(plainText)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, 18)
	if err != nil {
		return ""
	}

	return string(hashedPassword)
}

func CompareBcrypt(hashedString, plainString string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedString), []byte(plainString))
	if err != nil {
		return false
	}

	return true
}

func Generate16ByteHash(s string) (salt string) {
	// Create a SHA-256 hash of the input string
	hash := sha256.Sum256([]byte(s))

	// Take the first 16 bytes of the hash
	fixed16Bytes := hash[:16]

	// Convert the 16 bytes to a Base64-encoded string
	salt = base64.StdEncoding.EncodeToString(fixed16Bytes)

	return
}

// RandomStringNumber generates a random string of length n
func RandomAlphaNumString(n int) string {
	var sb strings.Builder
	k := len(alphanums)

	for i := 0; i < n; i++ {
		c := alphanums[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomNumberString generates a random string of length n
func RandomNumberString(n int) string {
	var sb strings.Builder
	k := len(numbers)

	for i := 0; i < n; i++ {
		c := numbers[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// GenerateOtpNumber generates a random positive int number
func GenerateOtpNumber() string {
	return RandomNumberString(6)
}

// Helper function to check if the slice contains a string
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateRandString ..
func GenerateRandString(lg int) string {
	var letterRunes = []rune("123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, lg)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// Generate code
func GenerateCode(code string, year, counter int) string {
	counterStr := fmt.Sprintf("%04d", counter)
	return fmt.Sprintf("%s-%d%s%04s", code, year, "000", counterStr)
}

func GenerateAPIKey() (string, error) {
	// Buat buffer yang cukup besar untuk menyimpan nilai random
	randomBytes := make([]byte, 32)

	// Baca nilai random ke dalam buffer
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Hash nilai random menggunakan SHA256
	hash := sha256.Sum256(randomBytes)

	// Encode hasil hash ke dalam bentuk base64 untuk mendapatkan string yang dapat dibaca
	apiKey := base64.StdEncoding.EncodeToString(hash[:])

	return apiKey, nil
}
