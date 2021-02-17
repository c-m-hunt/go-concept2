package data

import "time"

type Workout struct {
	ID          string
	Date        time.Time
	Description string
	Duration    time.Duration
	Distance    int
	StrokeRate  int
	StrokeCount int
	DragFactor  int
}
