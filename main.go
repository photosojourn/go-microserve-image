package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/gorilla/handlers"
)

func handler(w http.ResponseWriter, r *http.Request) {
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<title>{{ .Title}}</title>
	</head> 
	<body>
	<img src="/images/image.png">
	</body
</html>`

	t, err := template.New("webpage").Parse(tpl)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to render template")
	}

	data := struct {
		Title string
		Image string
	}{
		Title: os.Getenv("GO_MS_MESSAGE"),
		Image: "image.png",
	}

	err = t.Execute(w, data)

	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to display template")
	}

}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func loadImage() {
	sess := session.Must(session.NewSession())
	downloader := s3manager.NewDownloader(sess)

	os.Mkdir("images", 0755)
	f, err := os.Create("images/image.png")
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to create file image.png, %v", err)
		return
	}

	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String("imageA.png"),
	})

	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to download file %v, size %v", err, n)
		return
	}
}

func main() {
	loadImage()
	fmt.Fprintf(os.Stdout, "Starting listening on Port 8080\n")
	http.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handler)))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images/"))))
	http.HandleFunc("/health", healthcheckHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
