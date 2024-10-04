package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	res := folder.GetAllFolders()

	// example usage
	folderDriver := folder.NewDriver(res)
	orgFolder := folderDriver.GetFoldersByOrgID(orgID)

	folder.PrettyPrint(res)
	fmt.Printf("\n Folders for orgID: %s", orgID)
	folder.PrettyPrint(orgFolder)
}

// package main

// import (
// 	"github.com/georgechieng-sc/interns-2022/folders"
// )

// func main() {
// 	res := folders.GenerateData()

// 	folders.PrettyPrint(res)

// 	folders.WriteSampleData(res)
// }
