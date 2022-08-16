package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)


type JWTService interface {
	GenerateToken(UserId string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
    UserID string `json:"user_id"`
    jwt.StandardClaims
}

type jwtService struct {
    secretKey string
    issuer string
}

func NewJWTService() JWTService {
    return &jwtService{
        issuer: "sulkifli",
        secretKey: getSecretKey(),
    }
}

func getSecretKey() string {
    secretKey := os.Getenv("JWT_SECRET")
    if secretKey != "" {
        secretKey = "kiflisecret"
    }
    return secretKey
}

func (j *jwtService) GenerateToken(UserID string) string  {
    claims := &jwtCustomClaim{
        UserID,
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 4).Unix(),
            Issuer: j.issuer,
            IssuedAt: time.Now().Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    t, err := token.SignedString([]byte(j.secretKey))
    if err != nil {
        panic(err)
    }

    return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error){
    return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method %v", t.Header["alg"])
        }
        return []byte(j.secretKey), nil
    })
}