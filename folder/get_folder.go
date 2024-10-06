package folder

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for i := range folders {

		// storing pointers to f.folders folders with orgID
		// for GetAllChildFolders to use
		folder := &f.folders[i]

		if folder.OrgId == orgID {
			res = append(res, *folder)

			f.orgIDFoldersPointers = append(f.orgIDFoldersPointers, folder)
		}
	}

	return res
}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) (childFolders []Folder, err error) {

	orgIDs := map[uuid.UUID]bool{}
	folderExists := false
	folderInOrgExists := false

	allFolders := f.folders

	if len(allFolders) == 0 {
		var err error = errors.New("error: There are no folders in the driver")
		return nil, err
	}

	for _, folder := range allFolders {

		orgIDs[folder.OrgId] = true

		if folder.Name == name {
			folderExists = true

			if folder.OrgId == orgID {
				folderInOrgExists = true
				break
			}
		}
	}

	if !orgIDs[orgID] {
		var err error = errors.New("error: orgID does not exist")
		return nil, err
	}

	if !folderExists {
		var err error = errors.New("error: Folder does not exist")
		return nil, err
	}

	if !folderInOrgExists {
		var err error = errors.New("error: Folder does not exist in the specified organization")
		return nil, err
	}

	// we are not using the output of GetFoldersByOrgID,
	// we are using the field it initializes in f.orgIDFoldersPointers below
	f.GetFoldersByOrgID(orgID)

	childFolders = make([]Folder, 0)

	// reset pointers from previous function call!
	f.childFoldersPointers = nil

	for _, folder := range f.orgIDFoldersPointers {

		// Example:
		// name+"." = bravo.
		// folder.Paths = alpha.bravo.charlie
		// --> charlie is a child of bravo

		if strings.Contains((*folder).Paths, name+".") {

			childFolders = append(childFolders, *folder)

			// used in MoveFolder function
			f.childFoldersPointers = append(f.childFoldersPointers, folder)
		}
	}

	return childFolders, nil
}
