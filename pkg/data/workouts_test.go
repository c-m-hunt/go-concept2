package data_test

import (
	"testing"

	"github.com/c-m-hunt/go-concept2/pkg/file"
)

func TestItFindsMaxLengthWorkout(t *testing.T) {
	fp := "../file/testdata/c2data.csv"
	wos := file.LoadWorkouts(fp)
	wo := wos.GetLongestWorkout()
	if wo.Distance != 21098 {
		t.Errorf("Not finding the longest row %v", wo.ID)
	}
}
