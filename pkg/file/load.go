package file

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/c-m-hunt/go-concept2/pkg/data"
)

func LoadWorkouts(path string) []data.Workout {
	in, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Cannot open file at %v", path)
	}

	r := csv.NewReader(strings.NewReader(string(in)))
	records, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	wos := []data.Workout{}

	for i, record := range records {

		if i == 0 {
			continue
		}

		fmt.Println(record)

		workoutTime, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			fmt.Printf("Could not load workout time for ID %v\n", record[0])
		}

		date, err := time.Parse("2006-01-02 15:04:05", record[1])

		if err != nil {
			log.Fatalf("Could not load date for ID %v\n", record[0])
		}
		d, err := strconv.Atoi(record[7])
		if err != nil {
			fmt.Printf("Could not load distance for ID %v\n", record[0])
		}
		sr, err := strconv.Atoi(record[9])
		if err != nil {
			fmt.Printf("Could not load stroke rate for ID %v\n", record[0])
		}
		sc, err := strconv.Atoi(record[10])
		if err != nil {
			fmt.Printf("Could not load stroke count for ID %v\n", record[0])
		}
		df, err := strconv.Atoi(record[16])
		if err != nil {
			fmt.Printf("Could not load drag factor for ID %v\n", record[0])
		}

		wo := data.Workout{
			ID:          record[0],
			Date:        date,
			Duration:    time.Duration(workoutTime*1000) * time.Millisecond,
			Distance:    d,
			StrokeRate:  sr,
			StrokeCount: sc,
			DragFactor:  df,
		}
		wos = append(wos, wo)
	}
	return wos
}
