package cloudstorage

import (
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/washington-shoji/gin-api/config"
)

func CloudinaryConnection() (*cloudinary.Cloudinary, error) {

	cld, err := cloudinary.NewFromParams(config.EnvConfig("CLOUDINARY_CLOUD_NAME"), config.EnvConfig("CLOUDINARY_API_KEY"), config.EnvConfig("CLOUDINARY_API_SECRET"))
	if err != nil {
		fmt.Println("setup err", err)
		return nil, err
	}

	return cld, nil
}
