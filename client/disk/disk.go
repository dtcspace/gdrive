package disk

import (
	"google.golang.org/api/drive/v3"
)

type MyDiskCall struct {
	*drive.Service
}

func New(service *drive.Service) *MyDiskCall {
	return &MyDiskCall{service}
}
