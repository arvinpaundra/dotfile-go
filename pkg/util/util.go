package util

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func LoadVersion() string {
	content, err := os.ReadFile("version.txt")

	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSuffix(string(content), "\n")
}

func GetUuid() string {
	id, err := uuid.NewV7()
	if err != nil {
		return ""
	}

	return id.String()
}

func InArrayNumber(slice []int, n int) bool {
	for _, s := range slice {
		if s == n {
			return true
		}
	}
	return false
}

func NumberSliceToString(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}

	s := make([]string, len(slice))

	for i, v := range slice {
		s[i] = strconv.Itoa(v)
	}

	return strings.Join(s, sep)
}

func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(b)
}

func StringToSlices(str, sep string) []string {
	splitted := strings.Split(str, sep)

	var sanitizedStr []string

	for _, s := range splitted {
		sanitizedStr = append(sanitizedStr, strings.TrimSpace(s))
	}

	return sanitizedStr
}
