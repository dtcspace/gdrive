package disk

import (
	"fmt"
	"log"
)

func (client *MyDisk) List() (err error) {
	d,err := client.service.Drives.List().Do()
	if err != nil {
		log.Fatalf("Unable to retrieve drives: %v", err)
	}
	for _, i := range d.Drives {
		fmt.Printf("%s (%s)\n", i.Name, i.Id)
	}
	return
}
