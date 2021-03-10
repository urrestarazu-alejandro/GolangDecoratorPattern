package model

import (
	"time"
)

type Day struct {
	date       time.Time
	workingDay bool
}

func NewDay(date time.Time, workingDay bool) *Day {
	return &Day{date: date, workingDay: workingDay}
}

func (d *Day) IsWorkingDay() bool {
	return d.workingDay
}

func (d *Day) GetDate() time.Time {
	return d.date
}
