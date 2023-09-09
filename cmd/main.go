package main

import (
	"log"
	"net/http"

	"github.com/diegolopezcode/Go-ApiImage-AwsS3/configs"
)

type Client struct {
	Token     string
	hc        http.Client
	Remaining int
}

// NewClient returns a new Client for the given token.
func NewClient(token string) *Client {
	c := http.Client{}
	return &Client{
		Token: token,
		hc:    c,
	}
}

func main() {
	var TOKEN, err = configs.GetConfig("TOKEN")
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

}
