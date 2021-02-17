package file_test

import (
	"testing"
	"time"

	"github.com/c-m-hunt/go-concept2/pkg/file"
)

func TestItLoadsWorkouts(t *testing.T) {
	fp := "./testdata/c2data.csv"
	wos := file.LoadWorkouts(fp)
	expectedLength := 274
	if len(wos) != expectedLength {
		t.Errorf("Not the right amount of workouts loaded. Wanted %v, got %v", expectedLength, len(wos))
	}

	type ex struct {
		idx      int
		ID       string
		Duration time.Duration
		Distance int
	}

	tests := []ex{
		{0, "51640663", time.Duration(1801400 * time.Millisecond), 7096},
		{273, "44085045", time.Duration(2544100 * time.Millisecond), 10000},
	}

	for _, ex := range tests {
		if wos[ex.idx].ID != ex.ID {
			t.Errorf("Workout ID not loaded correctly")
		}
		if wos[ex.idx].Duration != ex.Duration {
			t.Errorf("Workout distance not loaded correctly. Got %v, expected %v", wos[ex.idx].Duration, ex.Duration)
		}
		if wos[ex.idx].Distance != ex.Distance {
			t.Errorf("Workout distance not loaded correctly")
		}
	}

}
