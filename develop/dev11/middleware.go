package main

import (
	"log"
	"net/http"
)

// middleware функция обвёртки http.Handler,
// функция логирует http.Request
func middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Вызовите основго обработчика
		next.ServeHTTP(w, r)

		// логирование http.Request
		log.Printf("Method: %s, Url: %s", r.Method, r.URL)
	})
}
