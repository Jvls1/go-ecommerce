package helpers

import (
	"fmt"
	"log"
	"net/http"
)

func HandlerError(w http.ResponseWriter, err error, message string, statusCode int) {
	log.Println("Error:", err)

	w.WriteHeader(statusCode)
	w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, message)))
}
