package decorator

import (
	"time"
)

/*
	Concrete decorator
*/
type Holiday struct {
	workDay WorkDay
}

func NewHoliday(workDay WorkDay) *Holiday {
	return &Holiday{workDay: workDay}
}

func (h *Holiday) IsWorkingDay() bool {
	peroncho := time.Date(2021, time.March, 19, 0, 0, 0, 0, time.UTC)

	if h.workDay.GetDate().Year() == peroncho.Year() &&
		h.workDay.GetDate().Month() == peroncho.Month() &&
		h.workDay.GetDate().Day() == peroncho.Day() {

		return false
	}

	return h.workDay.IsWorkingDay()
}

func (h *Holiday) GetDate() time.Time {
	return h.workDay.GetDate()
}
