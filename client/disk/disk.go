package disk

import (
	"google.golang.org/api/drive/v3"
)

type MyDisk struct {
	service *drive.Service
}

func New(service *drive.Service) *MyDisk {
	return &MyDisk{service}
}
