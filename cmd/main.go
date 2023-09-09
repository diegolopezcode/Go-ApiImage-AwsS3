package main

import (
	"fmt"
	"net/http"

	"github.com/diegolopezcode/Go-ApiImage-AwsS3/handlers"
)

func main() {
	http.HandleFunc("/api/v1/photo", handlers.GetPhoto)
	port := "8080"
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	/* c := handlers.NewClient(TOKEN)
	result, err := handlers.GetPhoto(c, "cat") */

}
