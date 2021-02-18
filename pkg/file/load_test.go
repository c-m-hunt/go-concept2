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
		idx          int
		ID           string
		Duration     time.Duration
		Distance     int
		StrokeRate   int
		StrokeCount  int
		DragFactor   int
		TotalCalores int
	}

	tests := []ex{
		{0, "51640663", time.Duration(1801400 * time.Millisecond), 7096, 22, 670, 122, 441},
		{273, "44085045", time.Duration(2544100 * time.Millisecond), 10000, 19, 818, 126, 624},
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
		if wos[ex.idx].StrokeCount != ex.StrokeCount {
			t.Errorf("Stroke count not loaded correctly")
		}
		if wos[ex.idx].StrokeRate != ex.StrokeRate {
			t.Errorf("Stroke rate not loaded correctly")
		}
		if wos[ex.idx].DragFactor != ex.DragFactor {
			t.Errorf("Workout drag factor not loaded correctly")
		}
		if wos[ex.idx].TotalCalores != ex.TotalCalores {
			t.Errorf("Total calories not loaded correctly")
		}
	}
	if wos[150].IsInterval == false {
		t.Error("Not correctly loading is interval flag")
	}
}
