package file

import (
	"encoding/csv"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/c-m-hunt/go-concept2/pkg/data"
	log "github.com/sirupsen/logrus"
)

// LoadWorkoutsDir Loads all the files in directory
func LoadWorkoutsDir(path string) (data.Workouts, error) {
	wos := data.Workouts{}

	filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {

		if d.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".csv" {
			newWos, err := LoadWorkouts(path)
			if err != nil {
				panic(err)
			}
			wos = append(wos, newWos...)
		}
		return nil
	})

	return wos, nil
}

// CreateWorkoutsFromCSV Creates workouts struct from CSV reader
func CreateWorkoutsFromCSV(csvData io.Reader) (data.Workouts, error) {

	r := csv.NewReader(csvData)
	records, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	wos := []data.Workout{}

	for i, record := range records {

		// Skip the header record
		if i == 0 {
			continue
		}

		workoutTime, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			log.Debugf("Could not load workout time for ID %v\n", record[0])
		}
		date, err := time.Parse("2006-01-02 15:04:05", record[1])
		if err != nil {
			log.Debugf("Could not load date for ID %v\n", record[0])
		}
		d, err := strconv.Atoi(record[7])
		if err != nil {
			log.Debugf("Could not load distance for ID %v\n", record[0])
		}
		sr, err := strconv.Atoi(record[9])
		if err != nil {
			log.Debugf("Could not load stroke rate for ID %v\n", record[0])
		}
		sc, err := strconv.Atoi(record[10])
		if err != nil {
			log.Debugf("Could not load stroke count for ID %v\n", record[0])
		}
		df, err := strconv.Atoi(record[16])
		if err != nil {
			log.Debugf("Could not load drag factor for ID %v\n", record[0])
		}
		tc, err := strconv.Atoi(record[14])
		if err != nil {
			log.Debugf("Could not load drag factor for ID %v\n", record[0])
		}
		isInt := false
		if len(record[5]) > 0 {
			isInt = true
		}

		wo := data.Workout{
			ID:           record[0],
			Date:         date,
			IsInterval:   isInt,
			Duration:     time.Duration(workoutTime*1000) * time.Millisecond,
			Distance:     d,
			StrokeRate:   sr,
			StrokeCount:  sc,
			DragFactor:   df,
			TotalCalores: tc,
		}
		wos = append(wos, wo)
	}
	return wos, nil
}

// LoadWorkouts Loads workouts from a CSV file which has been downloaded from Concept2
func LoadWorkouts(path string) (data.Workouts, error) {
	in, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Cannot open file at %v", path)
	}
	sr := strings.NewReader(string(in))
	return CreateWorkoutsFromCSV(sr)
}
