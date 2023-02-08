package disk

import "google.golang.org/api/drive/v3"

func SwitchFilesListDrive(call *drive.FilesListCall) *drive.FilesListCall {
	call.Corpora("drive")
	call.IncludeItemsFromAllDrives(true)
	call.TeamDriveId("0AG5icJJpNXAKUk9PVA")
	call.SupportsAllDrives(true)
	return call
}

func SwitchFilesCreateDrive(call *drive.FilesCreateCall) *drive.FilesCreateCall {
	call.SupportsAllDrives(true)
	return call
}

func SwitchFilesUpdateDrive(call *drive.FilesUpdateCall) *drive.FilesUpdateCall {
	call.SupportsAllDrives(true)
	return call
}

func SwitchFilesGetDrive(call *drive.FilesGetCall) *drive.FilesGetCall {
	call.SupportsAllDrives(true)
	return call
}
