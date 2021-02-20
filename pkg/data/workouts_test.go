package data_test

import (
	"testing"

	"github.com/c-m-hunt/go-concept2/pkg/file"
)

var dp = "../file/testdata"
var fp = "../file/testdata/concept2-season-2021.csv"

func TestItGetsTotalDistance(t *testing.T) {
	wos, _ := file.LoadWorkoutsDir(dp)
	d := wos.GetTotalDistance()
	exp := 3918533
	if d != exp {
		t.Errorf("Not getting correct total distance. Expected %v, got %v", exp, d)
	}
}

func TestItFindsMaxLengthWorkout(t *testing.T) {
	wos, _ := file.LoadWorkouts(fp)
	wo := wos.GetLongestWorkout()
	if wo.Distance != 21098 {
		t.Errorf("Not finding the longest row %v", wo.ID)
	}
}

func TestItFiltersOutShortWorkouts(t *testing.T) {
	wos, _ := file.LoadWorkouts(fp)
	wosFiltered := wos.FilterShortWorkouts(200)
	if len(wosFiltered) != 271 {
		t.Errorf("Is not correctly filtering short workouts, got %v", len(wosFiltered))
	}
}

func TestItFiltersOutWorkoutsByLength(t *testing.T) {
	wos, _ := file.LoadWorkouts(fp)
	wosFiltered := wos.FilterWorkoutsByDistance(5000, 5050)
	if len(wosFiltered) != 59 {
		t.Errorf("Is not correctly filtering workouts by length, got %v", len(wosFiltered))
	}
}
