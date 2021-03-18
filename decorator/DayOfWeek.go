package decorator

import (
	"time"
)

/*
	Concrete decorator
*/
type DayOfWeek struct {
	workDay WorkDay
}

func NewDayOfWeek(workDay WorkDay) *DayOfWeek {
	return &DayOfWeek{workDay: workDay}
}

func (d *DayOfWeek) IsWorkingDay() bool {
	switch d.workDay.GetDate().Weekday() {
	case time.Sunday, time.Saturday:
		return false
	default:
		return d.workDay.IsWorkingDay()
	}
}

func (d *DayOfWeek) GetDate() time.Time {
	return d.workDay.GetDate()
}
