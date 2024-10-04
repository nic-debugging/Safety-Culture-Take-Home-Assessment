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
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
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
		// here, the folder exists so it passed the if statement above
		// but is in another organization.
		var err error = errors.New("error: Folder does not exist in the specified organization")
		return nil, err
	}

	orgIDFolders := f.GetFoldersByOrgID(orgID)

	// start with length 0, since there can be 0 valid child folders
	// and we can always make slice longer
	childFolders = make([]Folder, 0)

	for _, folder := range orgIDFolders {

		// name is the prefix of the given folder, and if current
		// orgIDFolder has the prefix and a "." afterwards,
		// it means it is a child of the folder.

		if strings.HasPrefix(folder.Paths, name+".") {
			childFolders = append(childFolders, folder)
		}
	}

	return childFolders, nil
}
