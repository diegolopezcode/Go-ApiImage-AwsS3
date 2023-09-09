package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/diegolopezcode/Go-ApiImage-AwsS3/configs"
)

const (
	// S3Bucket is the bucket we will use
	photoApi   = "https://api.pexels.com/v1/photos"
	videoApi   = "https://api.pexels.com/videos"
	awsRegion  = "your-aws-region"
	bucketName = "your-s3-bucket-name"
	objectKey  = "images/image.jpg"
)

type SearchPhoto struct {
	TotalResults int     `json:"total_results"`
	Page         int     `json:"page"`
	PerPage      int     `json:"per_page"`
	Photos       []Photo `json:"photos"`
	NextPage     string  `json:"next_page"`
}
type Photo struct {
	ID              int      `json:"id"`
	Width           int      `json:"width"`
	Height          int      `json:"height"`
	URL             string   `json:"url"`
	Photographer    string   `json:"photographer"`
	PhotographerURL string   `json:"photographer_url"`
	PhotographerID  int      `json:"photographer_id"`
	AvgColor        string   `json:"avg_color"`
	Src             PhotoSrc `json:"src"`
	Liked           bool     `json:"liked"`
	Alt             string   `json:"alt"`
}

type PhotoSrc struct {
	Original  string `json:"original"`
	Large2X   string `json:"large2x"`
	Large     string `json:"large"`
	Medium    string `json:"medium"`
	Small     string `json:"small"`
	Portrait  string `json:"portrait"`
	Landscape string `json:"landscape"`
	Tiny      string `json:"tiny"`
}

type Client struct {
	Token     string
	Hc        http.Client
	Remaining int
}

var TOKEN, err = configs.GetConfig("API_KEY_PEXEL")

// NewClient returns a new Client for the given token.
func NewClient(token string) *Client {
	c := http.Client{}
	return &Client{
		Token: token,
		Hc:    c,
	}
}

func SearchPhotos(c *Client, query string) (*SearchPhoto, error) {
	return nil, nil
}

func GetPhoto(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	fmt.Print(
		"r.URL.Query(): ", r.URL.Query(),
	)
	// Get the photo
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Id not found", http.StatusBadRequest)
	}
	fmt.Print("id: ", id)
	req, err := http.NewRequest(http.MethodGet, photoApi+"/"+id, nil)
	if err != nil {
		http.Error(w, "Error getting photo", http.StatusInternalServerError)
	}

	req.Header.Set("Authorization", TOKEN)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error getting photo", http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	// Get the photo
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error getting photo", http.StatusInternalServerError)
	}

	pht := Photo{}
	err = json.Unmarshal([]byte(body), &pht)
	if err != nil {
		http.Error(w, "Error with the data, review provider", http.StatusInternalServerError)
	}

	res, err := json.Marshal(pht)
	if err != nil {
		http.Error(w, "Error with the data, review provider", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetPhotos() {

}

func AddPhotoS3(urlImage string) error {
	resp, err := http.Get(urlImage)
	if err != nil {
		return errors.New("Error downloadin photo")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Error downloadin photo with status code: " + string(resp.StatusCode))
	}

	// Create an AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
	})
	if err != nil {
		return fmt.Errorf("Error creating AWS session: %v", err)
	}
	// Create an S3 client
	s3Client := s3.New(sess)

	// Upload the image to S3
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(bucketName),
		Key:           aws.String(objectKey),
		Body:          resp.Body,
		ContentType:   aws.String("image/jpeg"),
		ContentLength: aws.Int64(resp.ContentLength),
	})
	if err != nil {
		return fmt.Errorf("Error uploading image to S3: %v", err)
	}

	fmt.Println("Image uploaded successfully to S3")
	return nil
}

func GetVideo() {

}
