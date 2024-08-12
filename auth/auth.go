package auth

import (
	"internal/intern"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey =  []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	Email string   `json:"email"`
	jwt.StandardClaims
}
//GenerateJWT generates jwt token

func GenerateJWT(email string)(string, error){
	expirtationTime := time.Now().Add(24 * time.Hour)

	claims:= &Claims {
		Email:email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirtationTime.Unix(),
		},

	}
	token:= jwt.NewWithClaims(jwt.SigningMethodES256,claims)
	return token.SignedString(jwtKey)

}

//ValidateJWT validates the JWT tokens
func ValidateJWT (tokenStr string) (*Claims, error) {
claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func( token *jwt.Token) (interface{}, error) {
		return jwtKey,nil 
	})
	if err != nil {
		return nil, err 
	}
	if !token.Valid {
			return nil, jwt.NewValidationError("Token is invalid", jwt.ValidationErrorSignatureInvalid)
	}
	return claims, nil 
}
