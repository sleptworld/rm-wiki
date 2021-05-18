package Middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sleptworld/test/Config"
	"net/http"
)

var (
	Malformed   = "10001"
	Expired     = "10002"
	NotValidYet = "10003"
	INVALID     = "10000"
)

type JWT struct {
	SigningKey []byte
}

type CustomClaims struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParserToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New(Malformed)
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New(Expired)
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New(NotValidYet)
			} else {
				return nil, errors.New(INVALID)
			}
		}
		return nil, errors.New(INVALID)
	} else {
		claims, ok := token.Claims.(*CustomClaims)
		if ok && token.Valid {
			return claims, nil
		}
		return nil, errors.New(INVALID)
	}
}

func NewJWT(k string) *JWT {
	return &JWT{
		[]byte(k),
	}
}

func Jwt() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("token")
		if token == "" {
			context.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "no token",
				"data":   nil,
			})
			context.Abort()
			return
		}

		j := NewJWT(Config.JWTKey)

		claims, err := j.ParserToken(token)

		if err != nil {
			switch err.Error() {
			case Expired:
				context.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    "This token has expired.",
					"data":   nil,
				})
				context.Abort()
				return
			default:
				context.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    err.Error(),
					"data":   nil,
				})
				context.Abort()
				return
			}
		}
		context.Set("claims",claims)
	}

}
