package folder

import (
	"errors"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {

	sameFolderName := false

	if name == dst {
		// if source has same name as destination
		// there are 2 cases, both ending in an error;
		// case 1: source and dst are from same organization; they are same file
		// case 2: not from same organization; they are different files
		sameFolderName = true
	}
	numOrgsWithSourceFileName := 0

	allFolders := f.folders

	var sourceFolder *Folder = nil
	var destFolder *Folder = nil

	for i := range allFolders {

		folder := &allFolders[i]

		if sameFolderName {
			// we will check for duplicates. we will check if
			// there are multiple files with same name, or just 1
			// if there is only 1, source and dst are same file from same org
			// otherwise, source and dst are diff files from diff orgs
			if folder.Name == name {
				sourceFolder = folder

				numOrgsWithSourceFileName++
			}
			continue
		}

		if folder.Name == name {
			sourceFolder = folder
		} else if folder.Name == dst {
			destFolder = folder
		}
	}

	if sourceFolder == nil {
		var err error = errors.New("error: Source folder does not exist")
		return nil, err
	}

	if sameFolderName {
		// numOrgsWithSourceFileName = 0 is taken care of above

		var err error

		if numOrgsWithSourceFileName == 1 {
			// case 1
			err = errors.New("error: Cannot move a folder to itself")
		} else {
			// case 2
			err = errors.New("error: Cannot move source folder to a different organization")
		}
		return nil, err
	}

	if destFolder == nil {
		var err error = errors.New("error: Destination folder does not exist")
		return nil, err
	}

	if sourceFolder.OrgId != destFolder.OrgId {
		var err error = errors.New("error: Cannot move source folder to a different organization")
		return nil, err
	}

	// we are not using the output of getAllChildFolders,
	// we are using the field it initializes in f.childFoldersPointers below
	// in order to actually change child folders of source folder
	// in f.folders and return f.folders
	_, err := f.GetAllChildFolders(sourceFolder.OrgId, name)

	if err != nil {
		return nil, err
	}

	// EXPLANATION OF STORING POINTERS IN DRIVER STRUCT:
	// this allows me to more efficiently loop only for child folders
	// of the source, instead of looping for the entire driver's list of folders
	// and having to check if each folder is a child folder and is in the correct organization
	// , allowing me to reuse my previous function efficiently
	// it also helps the changes persist, although I understand it is not required

	// EXAMPLE:
	// destFolderPath = golf
	destFolderPath := destFolder.Paths

	for _, folder := range f.childFoldersPointers {

		if folder.Name == dst {
			var err error = errors.New("error: Cannot move a folder to a child of itself")
			return nil, err
		}

		// sourceFolder.Paths = alpha.bravo
		// sourceFolder.name = bravo
		// sourceFolderParentPath = alpha
		sourceFolderParentPath := strings.TrimSuffix(sourceFolder.Paths, "."+sourceFolder.Name)

		// folder.Paths = alpha.bravo.charlie
		// SourceToFolderPath = bravo.charlie
		sourceToFolderPath := strings.TrimPrefix(folder.Paths, sourceFolderParentPath+".")

		// newFolderPath = golf.bravo.charlie
		newFolderPath := destFolderPath + "." + sourceToFolderPath

		folder.Paths = newFolderPath
	}

	// changing path of source folder too
	sourceFolder.Paths = destFolderPath + "." + sourceFolder.Name

	return f.folders, nil
}
