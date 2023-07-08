package interceptor

import (
	"log"
	"net/http"
)

func GlobalInterceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Por exemplo, você pode fazer log da requisição
		log.Printf("Requisição recebida: %s %s", r.Method, r.URL.Path)

		// Chame o próximo handler na cadeia
		next.ServeHTTP(w, r)
	})
}
