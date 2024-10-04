package folder_test

import (
	"errors"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_GetAllFolders(t *testing.T) {

	sampleFolders := folder.GetSampleData()

	allFolders := folder.GetAllFolders()

	// GetAllFolders should return list of all sample folders
	assert.Equal(t, sampleFolders, allFolders)
}

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {

	const DefaultOrgIDString = "c1556e17-b7c0-45a3-a6ae-9546248fb17a"
	var DefaultOrgID = uuid.FromStringOrNil(DefaultOrgIDString)
	var otherOrgID = uuid.Must(uuid.NewV4())
	var OrgID_2 = uuid.Must(uuid.NewV4())
	emptyList := make([]folder.Folder, 0)

	driverValid := []folder.Folder{

		{
			Name:  "alpha",
			OrgId: DefaultOrgID,
			Paths: "alpha",
		},
		{
			Name:  "bravo",
			OrgId: DefaultOrgID,
			Paths: "alpha.bravo",
		},
		{
			Name:  "charlie",
			OrgId: DefaultOrgID,
			Paths: "alpha.bravo.charlie",
		},
		{
			Name:  "delta",
			OrgId: DefaultOrgID,
			Paths: "alpha.delta",
		},
		{
			Name:  "echo",
			OrgId: DefaultOrgID,
			Paths: "echo",
		},
		{
			Name:  "foxtrot",
			OrgId: OrgID_2,
			Paths: "foxtrot",
		},
	}

	foldersByDefaultOrgID := []folder.Folder{

		{
			Name:  "alpha",
			OrgId: DefaultOrgID,
			Paths: "alpha",
		},
		{
			Name:  "bravo",
			OrgId: DefaultOrgID,
			Paths: "alpha.bravo",
		},
		{
			Name:  "charlie",
			OrgId: DefaultOrgID,
			Paths: "alpha.bravo.charlie",
		},
		{
			Name:  "delta",
			OrgId: DefaultOrgID,
			Paths: "alpha.delta",
		},
		{
			Name:  "echo",
			OrgId: DefaultOrgID,
			Paths: "echo",
		},
	}

	t.Parallel()
	tests := [...]struct {
		testName string
		orgID    uuid.UUID
		folders  []folder.Folder
		want     []folder.Folder
	}{
		// not including any error checking tests
		// since I am not asked to change this function
		// only to test for it, and function doesn't have any error handling
		{
			testName: "Success case: checks if given a driver that contains folders with given OrgID, correctly returns a list of those folders",
			orgID:    DefaultOrgID,
			folders:  driverValid,
			want:     foldersByDefaultOrgID,
		},
		{
			testName: "Success case: checks if given a driver that does not contain folders with given OrgID, correctly returns an empty list",
			orgID:    otherOrgID,
			folders:  driverValid,
			want:     emptyList,
		},
		{
			testName: "Success case: checks if given an empty driver, correctly returns an empty list",
			orgID:    DefaultOrgID,
			folders:  emptyList,
			want:     emptyList,
		},
	}
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			f := folder.NewDriver(test.folders)
			get := f.GetFoldersByOrgID(test.orgID)

			assert.Equal(t, test.want, get)
		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {

	var otherOrgID = uuid.Must(uuid.NewV4())
	var invalidOrgID = uuid.Must(uuid.NewV4())
	const DefaultOrgIDString = "c1556e17-b7c0-45a3-a6ae-9546248fb17a"
	var DefaultOrgID = uuid.FromStringOrNil(DefaultOrgIDString)

	// test inputs:

	// using given example in README.md
	driverValid := []folder.Folder{

		{
			Name:  "alpha",
			OrgId: DefaultOrgID,
			Paths: "alpha",
		},
		{
			Name:  "bravo",
			OrgId: DefaultOrgID,
			Paths: "alpha.bravo",
		},
		{
			Name:  "charlie",
			OrgId: DefaultOrgID,
			Paths: "alpha.bravo.charlie",
		},
		{
			Name:  "delta",
			OrgId: DefaultOrgID,
			Paths: "alpha.delta",
		},
		{
			Name:  "echo",
			OrgId: DefaultOrgID,
			Paths: "echo",
		},
		{
			Name:  "foxtrot",
			OrgId: otherOrgID,
			Paths: "foxtrot",
		},
	}

	alphaChildFolders := []folder.Folder{
		{
			Name:  "bravo",
			OrgId: DefaultOrgID,
			Paths: "alpha.bravo",
		},
		{
			Name:  "charlie",
			OrgId: DefaultOrgID,
			Paths: "alpha.bravo.charlie",
		},
		{
			Name:  "delta",
			OrgId: DefaultOrgID,
			Paths: "alpha.delta",
		},
	}

	emptyList := make([]folder.Folder, 0)

	t.Parallel()
	tests := [...]struct {
		testName   string
		folderName string
		orgID      uuid.UUID
		folders    []folder.Folder
		want       []folder.Folder
		wantError  error
	}{
		{
			testName:   "Success case: checks if given valid orgId and name given, correct child folders are returned",
			folderName: "alpha",
			orgID:      DefaultOrgID,
			// maybe rename driverValid to some else
			folders:   driverValid,
			want:      alphaChildFolders,
			wantError: nil,
		},
		{
			testName:   "Success case: checks if given valid orgId and name, and given that folder has no child folders, an empty list is returned",
			folderName: "charlie",
			orgID:      DefaultOrgID,
			folders:    driverValid,
			want:       emptyList,
			wantError:  nil,
		},
		{
			testName:   "Fail case: checks if given invalid orgID, returns appropriate error",
			folderName: "alpha",
			orgID:      invalidOrgID,
			folders:    driverValid,
			want:       nil,
			wantError:  errors.New("error: orgID does not exist"),
		},
		{
			testName:   "Fail case: checks if given no orgID, returns appropriate error",
			folderName: "alpha",
			orgID:      uuid.Nil,
			folders:    driverValid,
			want:       nil,
			wantError:  errors.New("error: orgID does not exist"),
		},
		{
			testName:   "Fail case: checks if given folder that does not exist, returns appropriate error",
			folderName: "non-existent_folder",
			orgID:      DefaultOrgID,
			folders:    driverValid,
			want:       nil,
			wantError:  errors.New("error: Folder does not exist"),
		},
		{
			testName:   "Fail case: checks if given folder that does not exist in specified organization, returns appropriate error",
			folderName: "alpha",
			orgID:      otherOrgID,
			folders:    driverValid,
			want:       nil,
			wantError:  errors.New("error: Folder does not exist in the specified organization"),
		},
		{
			testName:   "Fail case: checks if given an empty driver with no folders, returns appropriate error",
			folderName: "alpha",
			orgID:      otherOrgID,
			folders:    emptyList,
			want:       nil,
			wantError:  errors.New("error: There are no folders in the driver"),
		},
		{
			testName:   "Fail case: checks if given an empty ",
			folderName: "alpha",
			orgID:      uuid.Nil,
			folders:    emptyList,
			want:       nil,
			wantError:  errors.New("error: There are no folders in the driver"),
		},

		// invalid path test case
	}
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {

			f := folder.NewDriver(test.folders)
			childFolders, err := f.GetAllChildFolders(test.orgID, test.folderName)

			if test.wantError == nil {
				assert.Equal(t, test.want, childFolders)
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.wantError.Error())
			}
		})
	}
}
