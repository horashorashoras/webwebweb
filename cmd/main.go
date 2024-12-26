package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	dir := getEnv("PUBLIC_DIR", "public")
	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)

	port := getEnv("PORT", ":8080")
	fmt.Println("> server is running on http://localhost" + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
