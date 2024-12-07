package config

import (
	"os"

	"github.com/imagekit-developer/imagekit-go"
	"github.com/joho/godotenv"
)

var Ikit *imagekit.ImageKit

func InitImageKitConfig() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
	ikit := imagekit.NewFromParams(imagekit.NewParams{
		PrivateKey:  os.Getenv("IMAGEKIT_PRIVATE_KEY"),
		PublicKey:   os.Getenv("IMAGEKIT_PUBLIC_KEY"),
		UrlEndpoint: os.Getenv("IMAGEKIT_URL_ENDPOINT"),
	})
	Ikit = ikit
	return
}
