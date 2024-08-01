package encrypt

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTEncoder interface {
	EncodeToJWT(data interface{}) (string, error)
	DecodeFromJWT(token string, v interface{}) error
}

type JWTService struct {
	secretKey string
}

func NewJWTService(secretKey string) *JWTService {
	return &JWTService{secretKey: secretKey}
}

func (j *JWTService) EncodeToJWT(data interface{}) (string, error) {
	claims := jwt.MapClaims{
		"data": data,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTService) DecodeFromJWT(tokenString string, v interface{}) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if data, ok := claims["data"]; ok {
			return mapToStruct(data, v)
		}
		return errors.New("data field not found in token claims")
	}
	return errors.New("invalid token")
}

func mapToStruct(data interface{}, v interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonData, v)
}

func HashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}
