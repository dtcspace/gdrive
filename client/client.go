package client

import (
	"golang.org/x/net/context"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"net/http"
)

type DriveClient struct {
	service *drive.Service
}

// New
// Deprecated
func New(client *http.Client) (*DriveClient, error) {
	service, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	return &DriveClient{service}, nil
}

func NewDriveService(client *http.Client) (*drive.Service, error) {
	return drive.NewService(context.Background(), option.WithHTTPClient(client))
}
