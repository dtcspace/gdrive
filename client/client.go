package client

import (
	"google.golang.org/api/drive/v3"
	"net/http"
)

type DriveClient struct {
	service *drive.Service
}

func New(client *http.Client) (*DriveClient, error) {
	service, err := drive.New(client)
	if err != nil {
		return nil, err
	}

	return &DriveClient{service}, nil
}
