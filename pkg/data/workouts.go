package data

import (
	"time"
)

// Workouts A slice of workouts
type Workouts []Workout

// Workout Struct of a workout
type Workout struct {
	ID           string
	Date         time.Time
	Description  string
	IsInterval   bool
	Duration     time.Duration
	Distance     int
	StrokeRate   int
	StrokeCount  int
	DragFactor   int
	TotalCalores int
}

// GetLongestWorkout - Returns the longest workout over the workout slice
func (wos Workouts) GetLongestWorkout() *Workout {
	var longest *Workout
	for i, w := range wos {
		if longest == nil || longest.Distance < w.Distance {
			longest = &wos[i]
		}
	}
	return longest
}

// FilterShortWorkouts - Returns workouts which are greater than a certain distance
func (wos Workouts) FilterShortWorkouts(minDist int) Workouts {
	return wos.FilterWorkoutsByDistance(minDist, 99999999)
}

// FilterWorkoutsByDistance - Returns the workouts between two distances
func (wos Workouts) FilterWorkoutsByDistance(minDist, maxDist int) Workouts {
	wosOut := Workouts{}
	for _, w := range wos {
		if w.Distance >= minDist && w.Distance < maxDist {
			wosOut = append(wosOut, w)
		}
	}
	return wosOut
}
