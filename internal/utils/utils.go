package utils

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"gocommerce/internal/config"
	"log"
	"net/http"
	"os"

	"github.com/imagekit-developer/imagekit-go/api/uploader"
)

func JSONResponse(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func UploadImageImageKit(file string) string {
	imageBytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	base64Image := "data:image/png;base64," + base64.StdEncoding.EncodeToString(imageBytes)

	ik, err := config.Ikit.Uploader.Upload(context.Background(), base64Image, uploader.UploadParam{
		FileName: file,
	})
	if err != nil {
		log.Fatal(err)
	}
	return ik.Data.Url
}