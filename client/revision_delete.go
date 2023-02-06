package client

import (
	"fmt"
	"io"
)

type DeleteRevisionArgs struct {
	Out        io.Writer
	FileId     string
	RevisionId string
}

func (client *DriveClient) DeleteRevision(args DeleteRevisionArgs) (err error) {
	rev, err := client.service.Revisions.Get(args.FileId, args.RevisionId).Fields("originalFilename").Do()
	if err != nil {
		return fmt.Errorf("Failed to get revision: %s", err)
	}

	if rev.OriginalFilename == "" {
		return fmt.Errorf("Deleting revisions for this file type is not supported")
	}

	err = client.service.Revisions.Delete(args.FileId, args.RevisionId).Do()
	if err != nil {
		return fmt.Errorf("Failed to delete revision", err)
	}

	fmt.Fprintf(args.Out, "Deleted revision '%s'\n", args.RevisionId)
	return
}
