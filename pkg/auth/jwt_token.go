package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"time"
	"tomaxut/config"
)

type JwtCustomClaims struct {
	ID    uint
	Guard string
	jwt.StandardClaims
}

type Token struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

type JwtToken struct {
}

func New() *JwtToken {
	return &JwtToken{}
}

func (j *JwtToken) CreateUserToken(ID uint, guard string) (*Token, error) {
	expiresAt := time.Now().Add(time.Hour * 24 * time.Duration(config.Get().JwtExpireDay)).Unix()
	claims := &JwtCustomClaims{
		ID,
		guard,
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Get().JwtSecret))

	return &Token{Token: tokenString, ExpiresAt: expiresAt}, err
}

func (j *JwtToken) ParseToken(token string) (*JwtCustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Get().JwtSecret), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*JwtCustomClaims); ok && tokenClaims.Valid {
			logrus.WithFields(logrus.Fields{"claims": claims}).Debug(tokenClaims.Valid)
			return claims, nil
		}
	}
	return nil, err
}

func (j *JwtToken) JwtClaims(c *gin.Context) (*JwtCustomClaims, error) {
	token := c.GetHeader("Authorization")
	token = token[7:]
	claims, err := j.ParseToken(token)
	return claims, err
}

func (j *JwtToken) JwtUserId(c *gin.Context) uint {
	claims, _ := j.JwtClaims(c)
	return claims.ID
}

func (j *JwtToken) JwtUserGuard(c *gin.Context) string {
	claims, _ := j.JwtClaims(c)
	return claims.Guard
}
