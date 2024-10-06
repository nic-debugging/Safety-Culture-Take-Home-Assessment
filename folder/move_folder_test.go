package folder_test

import (
	"errors"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {

	var otherOrgID = uuid.Must(uuid.NewV4())
	const DefaultOrgIDString = "c1556e17-b7c0-45a3-a6ae-9546248fb17a"
	var DefaultOrgID = uuid.FromStringOrNil(DefaultOrgIDString)
	emptyList := make([]folder.Folder, 0)

	driver := []folder.Folder{
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
			Paths: "alpha.delta.echo",
		},
		{
			Name:  "foxtrot",
			OrgId: otherOrgID,
			Paths: "foxtrot",
		},
		{
			Name:  "golf",
			OrgId: DefaultOrgID,
			Paths: "golf",
		},
	}

	driverAfterMoveBravoToDelta := []folder.Folder{
		{
			Name:  "alpha",
			OrgId: DefaultOrgID,
			Paths: "alpha",
		},
		{
			Name:  "bravo",
			OrgId: DefaultOrgID,
			Paths: "alpha.delta.bravo",
		},
		{
			Name:  "charlie",
			OrgId: DefaultOrgID,
			Paths: "alpha.delta.bravo.charlie",
		},
		{
			Name:  "delta",
			OrgId: DefaultOrgID,
			Paths: "alpha.delta",
		},
		{
			Name:  "echo",
			OrgId: DefaultOrgID,
			Paths: "alpha.delta.echo",
		},
		{
			Name:  "foxtrot",
			OrgId: otherOrgID,
			Paths: "foxtrot",
		},
		{
			Name:  "golf",
			OrgId: DefaultOrgID,
			Paths: "golf",
		},
	}

	driverAfterMoveBravoToGolf := []folder.Folder{
		{
			Name:  "alpha",
			OrgId: DefaultOrgID,
			Paths: "alpha",
		},
		{
			Name:  "bravo",
			OrgId: DefaultOrgID,
			Paths: "golf.bravo",
		},
		{
			Name:  "charlie",
			OrgId: DefaultOrgID,
			Paths: "golf.bravo.charlie",
		},
		{
			Name:  "delta",
			OrgId: DefaultOrgID,
			Paths: "alpha.delta",
		},
		{
			Name:  "echo",
			OrgId: DefaultOrgID,
			Paths: "alpha.delta.echo",
		},
		{
			Name:  "foxtrot",
			OrgId: otherOrgID,
			Paths: "foxtrot",
		},
		{
			Name:  "golf",
			OrgId: DefaultOrgID,
			Paths: "golf",
		},
	}

	driver_2 := []folder.Folder{
		{
			Name:  "adam",
			OrgId: DefaultOrgID,
			Paths: "adam",
		},
		{
			Name:  "bill",
			OrgId: DefaultOrgID,
			Paths: "adam.bill",
		},
		{
			Name:  "cook",
			OrgId: DefaultOrgID,
			Paths: "cook",
		},
		{
			Name:  "cook",
			OrgId: otherOrgID,
			Paths: "cook",
		},
		{
			Name:  "dan",
			OrgId: DefaultOrgID,
			Paths: "dan",
		},
	}

	const SampleOrgIdString = "38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"
	var SampleOrgId = uuid.FromStringOrNil(SampleOrgIdString)

	driverFromSample := []folder.Folder{
		{
			Name:  "gentle-tempest",
			OrgId: SampleOrgId,
			Paths: "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf.gentle-tempest",
		},
		{
			Name:  "famous-rescue",
			OrgId: SampleOrgId,
			Paths: "creative-scalphunter.clear-arclight.topical-micromax.famous-rescue",
		},
		{
			Name:  "crucial-mister-sinister",
			OrgId: SampleOrgId,
			Paths: "creative-scalphunter.clear-arclight.topical-micromax.famous-rescue.crucial-mister-sinister",
		},
		{
			Name:  "flexible-iron-man",
			OrgId: SampleOrgId,
			Paths: "creative-scalphunter.clear-arclight.topical-micromax.famous-rescue.flexible-iron-man",
		},
	}

	driverFromSampleAnswer := []folder.Folder{
		{
			Name:  "gentle-tempest",
			OrgId: SampleOrgId,
			Paths: "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf.gentle-tempest",
		},
		{
			Name:  "famous-rescue",
			OrgId: SampleOrgId,
			Paths: "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf.gentle-tempest.famous-rescue",
		},
		{
			Name:  "crucial-mister-sinister",
			OrgId: SampleOrgId,
			Paths: "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf.gentle-tempest.famous-rescue.crucial-mister-sinister",
		},
		{
			Name:  "flexible-iron-man",
			OrgId: SampleOrgId,
			Paths: "creative-scalphunter.clear-arclight.topical-micromax.patient-red-wolf.gentle-tempest.famous-rescue.flexible-iron-man",
		},
	}

	// each test case calls MoveFolder once as persistence is not required
	tests := [...]struct {
		testName        string
		sourceName      string
		destinationName string
		folders         []folder.Folder
		want            []folder.Folder
		wantError       error
	}{
		{
			testName:        "Success case: checks if driver is correctly updated, given valid source folder bravo and valid destination delta",
			sourceName:      "bravo",
			destinationName: "delta",
			folders:         driver,
			want:            driverAfterMoveBravoToDelta,
			wantError:       nil,
		},
		{
			testName:        "Success case: checks if driver is correctly updated, given valid source folder bravo and valid destination golf",
			sourceName:      "bravo",
			destinationName: "golf",
			folders:         driver,
			want:            driverAfterMoveBravoToGolf,
			wantError:       nil,
		},
		{
			testName:        "Success case: checks if driver is correctly updated, given valid source and destination folder with long paths",
			sourceName:      "famous-rescue",
			destinationName: "gentle-tempest",
			folders:         driverFromSample,
			want:            driverFromSampleAnswer,
			wantError:       nil,
		},
		{
			testName:        "Fail case: checks if given a source folder that does not exist, returns appropriate error",
			sourceName:      "non-existent-folder",
			destinationName: "golf",
			folders:         driver,
			want:            nil,
			wantError:       errors.New("error: Source folder does not exist"),
		},
		{
			testName:        "Fail case: checks if given a destination folder that does not exist, returns appropriate error",
			sourceName:      "bravo",
			destinationName: "non-existent-folder",
			folders:         driver,
			want:            nil,
			wantError:       errors.New("error: Destination folder does not exist"),
		},
		{
			testName:        "Fail case: checks if given a source folder that does not belong to destination's organization, returns appropriate error",
			sourceName:      "bravo",
			destinationName: "foxtrot",
			folders:         driver,
			want:            nil,
			wantError:       errors.New("error: Cannot move source folder to a different organization"),
		},
		{
			testName:        "Fail case: checks if given a source folder that does not belong to destination's organization, and has same name as destination folder, returns appropriate error",
			sourceName:      "cook",
			destinationName: "cook",
			folders:         driver_2,
			want:            nil,
			wantError:       errors.New("error: Cannot move source folder to a different organization"),
		},
		{
			testName:        "Fail case: checks if given a source folder that is the same as destination folder, returns appropriate error",
			sourceName:      "bravo",
			destinationName: "bravo",
			folders:         driver,
			want:            nil,
			wantError:       errors.New("error: Cannot move a folder to itself"),
		},
		{
			testName:        "Fail case: checks if given a destination folder that is a child of the source folder, returns appropriate error",
			sourceName:      "bravo",
			destinationName: "charlie",
			folders:         driver,
			want:            nil,
			wantError:       errors.New("error: Cannot move a folder to a child of itself"),
		},
		{
			testName:        "Fail case: checks if given a destination folder that is a grandchild of the source folder, returns appropriate error",
			sourceName:      "alpha",
			destinationName: "charlie",
			folders:         driver,
			want:            nil,
			wantError:       errors.New("error: Cannot move a folder to a child of itself"),
		},
		{
			testName:        "Fail case: checks if given an empty folder, returns appropriate error",
			sourceName:      "bravo",
			destinationName: "golf",
			folders:         emptyList,
			want:            nil,
			wantError:       errors.New("error: Source folder does not exist"),
		},
	}
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {

			// make a copy so that changes to the drivers do not persist
			// as my MoveFolder function will permanently change the drivers
			foldersCopy := append([]folder.Folder{}, test.folders...)

			f := folder.NewDriver(foldersCopy)
			updatedDriver, err := f.MoveFolder(test.sourceName, test.destinationName)

			if test.wantError == nil {
				assert.Equal(t, test.want, updatedDriver)
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.wantError.Error())
			}
		})
	}
}
