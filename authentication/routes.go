package authentication

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func RegisterRoutes(r chi.Router) {
	r.Get("/refresh", AttemptTokenRefresh)
	r.Get("/login", login)
}

func AttemptTokenRefresh(w http.ResponseWriter, r *http.Request) {

	refreshToken, err := r.Cookie("refresh-token")
	fmt.Printf("%v", refreshToken)
	if err != nil {
		fmt.Println("error line 60")
		http.Error(w, "authentication error", http.StatusUnauthorized)
		return
	}
	token, err := RefreshToken(refreshToken.Value)
	if err != nil {
		http.Error(w, "authentication error", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"access-token": token})
}

func login(w http.ResponseWriter, r *http.Request) {
	pair, err := GenerateTokenPair(44)
	if err != nil {
		http.Error(w, "authentication error", http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(pair)
}
