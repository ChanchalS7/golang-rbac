package auth

import (
	"net/http"
	"strings"

	"github.com/ChanchalS7/golang-rbac/utils"
)

//AuthMiddleware check for a valid JWT token in the request

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		token := r.Header.Get("Authorization")

		if token == ""{
			utils.ResponseWithError(w, http.StatusUnauthorized, "Authorization token missing")
			return
		}
		// token = strings.TrimPrefix(token, "Bearer")

		//Extract the token form the header
		tokenParts := strings.Split(token, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			utils.ResponseWithError(w, http.StatusUnauthorized, "Malformed token")
			return 
		}

		// _, err := ValidateJWT(token)
		_, err := ValidateJWT(tokenParts[1])
		
		if err != nil {
			utils.ResponseWithError(w, http.StatusUnauthorized, "Invalid token")
			return 
		}
		// Proceed to the next handler if the token is valid
		next.ServeHTTP(w,r)
	})
}