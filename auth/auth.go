package auth

import (
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
	expirationTime := time.Now().Add(24 * time.Hour)

	claims:= &Claims {
		Email:email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},

	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token.SignedString(jwtKey)

}

//ValidateJWT validates the JWT tokens
func ValidateJWT (tokenStr string) (*Claims, error) {
claims := &Claims{}
	
token, err := jwt.ParseWithClaims(tokenStr, claims,func(token *jwt.Token)(interface{},error){
	//Ensure the tokens signing method is what you expect
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, jwt.NewValidationError("Unexpected signing method",jwt.ValidationErrorSignatureInvalid)
	}
	return jwtKey, nil
})
if err != nil {
	return nil, err 
}
if !token.Valid {
	return nil, jwt.NewValidationError("Token is invalid", jwt.ValidationErrorSignatureInvalid)
}
return claims, nil 
}
