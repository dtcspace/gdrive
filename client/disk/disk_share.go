package disk

import "google.golang.org/api/drive/v3"

func SwitchFilesListDrive(call *drive.FilesListCall) *drive.FilesListCall {
	call.Corpora("drive")
	call.IncludeItemsFromAllDrives(true)
	call.TeamDriveId("0AG5icJJpNXAKUk9PVA")
	return call.SupportsAllDrives(true)

}

func SwitchFilesCreateDrive(call *drive.FilesCreateCall) *drive.FilesCreateCall {
	return call.SupportsAllDrives(true)
}

func SwitchFilesGetDrive(call *drive.FilesGetCall) *drive.FilesGetCall {
	return call.SupportsAllDrives(true)
}
