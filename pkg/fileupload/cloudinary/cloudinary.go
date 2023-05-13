package fileupload

import (
	"context"
	"errors"
	"mime/multipart"
	"os"

	"online-course/pkg/utils"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

type FileUpload interface {
	Upload(file multipart.FileHeader) (*string, error)
	Delete(fileName string) (*string, error)
}

type FileUploadImpl struct {
}

// Delete implements File
func (fileUploadImpl *FileUploadImpl) Delete(fileName string) (*string, error) {
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

// Upload implements File
func (fileUploadImpl *FileUploadImpl) Upload(file multipart.FileHeader) (*string, error) {
	// Connect to cloudinary
	// cld, err := cloudinary.NewFromURL("cloudinary://<api_key>:<api_secret>@<cloud_name>")

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

func NewFileUpload() FileUpload {
	return &FileUploadImpl{}
}
