package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	//orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	res := folder.GetAllFolders()

	org1 := "38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"

	// example usage
	folderDriver := folder.NewDriver(res)
	orgFolder := folderDriver.GetFoldersByOrgID(uuid.FromStringOrNil(org1))

	// NEW OWN
	childFolders, _ := folderDriver.GetAllChildFolders(uuid.FromStringOrNil(org1), "topical-micromax")
	//updatedDriver, _ := folderDriver.MoveFolder("famous-rescue", "gentle-tempest")

	//////folder.PrettyPrint(res)
	fmt.Printf("\n 111111111111111Folders for orgID: %s", org1)
	folder.PrettyPrint(orgFolder)

	fmt.Printf("\n 222222222222222children for folder topical-micromax")
	folder.PrettyPrint(childFolders)

	//fmt.Printf("\n 333333333333333 updated driver ")
	//folder.PrettyPrint(updatedDriver)
}

// package main

// import (
// 	"fmt"

// 	"github.com/georgechieng-sc/interns-2022/folder"
// 	"github.com/georgechieng-sc/interns-2022/folders"
// 	"github.com/gofrs/uuid"
// )

// func main() {
// 	res := folders.GenerateData()

// 	folders.PrettyPrint(res)

// 	folders.WriteSampleData(res)

// 	folderDriver := folder.NewDriver(res)

// 	orgFolder := folderDriver.GetFoldersByOrgID(uuid.FromStringOrNil(org1))

// 	childFolders, _ := folderDriver.GetAllChildFolders(uuid.FromStringOrNil(org1), "topical-micromax")
// 	//updatedDriver, _ := folderDriver.MoveFolder("famous-rescue", "gentle-tempest")

// 	//////folder.PrettyPrint(res)
// 	fmt.Printf("\n 111111111111111Folders for orgID: %s", org1)
// 	folder.PrettyPrint(orgFolder)

// 	fmt.Printf("\n 222222222222222children for folder topical-micromax")
// 	folder.PrettyPrint(childFolders)
// }
