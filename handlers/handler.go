package handlers

const (
	// S3Bucket is the bucket we will use
	S3Bucket = "go-aws-s3"
	photoApi = "https://api.unsplash.com/photos/random"
	videoApi = "https://api.pexels.com/videos"
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

func getPhoto() {

}

func getPhotos() {

}

func addPhotoS3() {

}

func getVideo() {

}
