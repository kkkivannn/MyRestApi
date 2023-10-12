package handler

import (
	"crypto/sha1"
	"fmt"
	"time"
)

const (
	Salt       = "dslkj932q90jdqos0219jd3fjreasokcmnurn4875678f"
	SigningKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	TokenTTL   = 1 * time.Hour
)

func PasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(Salt)))
}
