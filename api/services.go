package api_user

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"time"
)

func HashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func ToString(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func FormatDateToString(date *time.Time) string {
	if date == nil {
		return ""
	}
	return date.Format("2000-01-01")
}

func FormatStringToDate(date *string) *time.Time {
	if date == nil || *date == "" {
		return nil
	}
	parsedDate, err := time.Parse("2000-01-01", *date)
	if err != nil {
		log.Println("Error parsing date:", err)
		return nil
	}
	return &parsedDate
}
