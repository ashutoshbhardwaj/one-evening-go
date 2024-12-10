package main

import (
	"fmt"
	"time"
	"errors"
)

type DateRange struct {
	start time.Time
	end   time.Time
}

func (d DateRange) Hours() float64 {
	return d.end.Sub(d.start).Hours()
}

func NewDateRange(start time.Time, end time.Time) (DateRange, error) {
	if end.Before(start) || start.IsZero() ||  end.IsZero() {
		return DateRange{}, errors.New("Either end is before start or start is zero or end is zero")
	}
	return DateRange{
		start : start,
		end: end,
	}, nil
}
func main() {
	lifetime, err := NewDateRange( time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
	time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC))

	if err != nil {
		panic("Error in lifetime")
	}


	fmt.Println(lifetime.Hours())

	travelInTime, _  := NewDateRange(time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC),
	 time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC))
	
	 if err != nil {
		panic("Error in travel time")
	}

	fmt.Println(travelInTime.Hours())
}
