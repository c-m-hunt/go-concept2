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

func TestItFiltersOutShortWorkouts(t *testing.T) {
	fp := "../file/testdata/c2data.csv"
	wos := file.LoadWorkouts(fp)
	wosFiltered := wos.FilterShortWorkouts(200)
	if len(wosFiltered) != 268 {
		t.Errorf("Is not correctly filtering short workouts, got %v", len(wosFiltered))
	}
}

func TestItFiltersOutWorkoutsByLength(t *testing.T) {
	fp := "../file/testdata/c2data.csv"
	wos := file.LoadWorkouts(fp)
	wosFiltered := wos.FilterWorkoutsByDistance(5000, 5050)
	if len(wosFiltered) != 58 {
		t.Errorf("Is not correctly filtering workouts by length, got %v", len(wosFiltered))
	}
}
