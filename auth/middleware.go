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
		token = strings.TrimPrefix(token, "Bearer")

		_, err := ValidateJWT(token)
		
		if err != nil {
			utils.ResponseWithError(w, http.StatusUnauthorized, "Invalid token")
			return 
		}
		next.ServeHTTP(w,r)
	})
}