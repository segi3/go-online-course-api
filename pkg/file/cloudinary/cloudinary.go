package image

import (
	"context"
	"errors"
	"mime/multipart"
	"online-course/pkg/utils"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

type File interface {
	Upload(file multipart.FileHeader) (*string, error)
	Delete(file string) (*string, error)
}

type FileImpl struct {
}

// Delete implements Image
func (fileImpl *FileImpl) Delete(fileName string) (*string, error) {
	// Connect to cloudinary
	cld, err := cloudinary.NewFromURL("cloudinary://" + os.Getenv("CLOUDINARY_APIKEY") + ":" + os.Getenv("CLOUDINARY_SECRET") + "@" + os.Getenv("CLOUDINARY_CLOUDNAME"))

	if err != nil {
		return nil, err
	}

	var ctx = context.Background()

	fileName = utils.GetFileName(fileName)

	resp, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: fileName})

	if err != nil {
		return nil, err
	}

	return &resp.Result, nil
}

// Upload implements Image
func (fileImpl *FileImpl) Upload(file multipart.FileHeader) (*string, error) {
	// connect to cloudinary
	cld, err := cloudinary.NewFromURL("cloudinary://" + os.Getenv("CLOUDINARY_APIKEY") + ":" + os.Getenv("CLOUDINARY_SECRET") + "@" + os.Getenv("CLOUDINARY_CLOUDNAME"))

	if err != nil {
		return nil, err
	}

	var ctx = context.Background()

	binary, err := file.Open()

	defer binary.Close()

	if err != nil {
		return nil, err
	}

	if binary != nil {
		uploadResult, err := cld.Upload.Upload(
			ctx,
			binary,
			uploader.UploadParams{
				PublicID: uuid.New().String(),
			},
		)

		if err != nil {
			return nil, err
		}

		return &uploadResult.SecureURL, nil
	}

	return nil, errors.New("format file tidak sesuai")
}

func NewFile() File {
	return &FileImpl{}
}
