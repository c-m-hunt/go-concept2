package file_test

import (
	"strings"
	"testing"
	"time"

	"github.com/c-m-hunt/go-concept2/pkg/file"
)

func TestItLoadsWorksoutsFromDir(t *testing.T) {
	fp := "./testdata"
	wos, _ := file.LoadWorkoutsDir(fp)
	if len(wos) != 1282 {
		t.Errorf("Got wrong amount of workouts. Expected %v, got %v", 1282, len(wos))
	}
}

func TestItLoadsWorksoutsFromString(t *testing.T) {
	csv := `"ID","Date","Description","Work Time (Formatted)","Work Time (Seconds)","Rest Time (Formatted)","Rest Time (Seconds)","Work Distance","Rest Distance","Stroke Rate/Cadence","Stroke Count","Pace","Avg Watts","Cal/Hour","Total Cal","Avg Heart Rate","Drag Factor","Age","Weight","Type","Ranked","Comments"
"44023457","2020-04-30 12:42:00","5000m row","""21:11.3""","1271.3","","","5000","","21","442","2:07.1","170","886","311","","126","42","Hwt","Indoor Rower","No",""
"44023458","2020-04-30 12:21:00","1:07 row","""1:07.5""","67.5","","","242","","20","21","2:19.4","129","744","14","","126","42","Hwt","Indoor Rower","No",""`
	csvReader := strings.NewReader(csv)
	wos, _ := file.CreateWorkoutsFromCSV(csvReader)
	if len(wos) != 2 {
		t.Errorf("Didn't load correct amount of records")
	}
}

func TestItLoadsWorkouts(t *testing.T) {
	fp := "./testdata/concept2-season-2021.csv"
	wos, _ := file.LoadWorkouts(fp)
	expectedLength := 277
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
		{3, "51640663", time.Duration(1801400 * time.Millisecond), 7096, 22, 670, 122, 441},
		{276, "44085045", time.Duration(2544100 * time.Millisecond), 10000, 19, 818, 126, 624},
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
	if wos[153].IsInterval == false {
		t.Error("Not correctly loading is interval flag")
	}
}
