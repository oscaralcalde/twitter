package middlew

import (
	"net/http"

	"github.com/oscaralcalde/twitter/bd"
)

func BDCheck(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ConnectionCheck() == 0 {
			http.Error(w, " BD connection lost", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
