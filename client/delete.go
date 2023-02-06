package client

import (
	"fmt"
	"io"
)

type DeleteArgs struct {
	Out       io.Writer
	Id        string
	Recursive bool
}

func (client *DriveClient) Delete(args DeleteArgs) error {
	f, err := client.service.Files.Get(args.Id).Fields("name", "mimeType").Do()
	if err != nil {
		return fmt.Errorf("Failed to get file: %s", err)
	}

	if isDir(f) && !args.Recursive {
		return fmt.Errorf("'%s' is a directory, use the 'recursive' flag to delete directories", f.Name)
	}

	err = client.service.Files.Delete(args.Id).Do()
	if err != nil {
		return fmt.Errorf("Failed to delete file: %s", err)
	}

	fmt.Fprintf(args.Out, "Deleted '%s'\n", f.Name)
	return nil
}

func (client *DriveClient) deleteFile(fileId string) error {
	err := client.service.Files.Delete(fileId).Do()
	if err != nil {
		return fmt.Errorf("Failed to delete file: %s", err)
	}
	return nil
}
