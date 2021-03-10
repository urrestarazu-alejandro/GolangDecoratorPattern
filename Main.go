package main

import (
	"GolangDecoratorPattern/decorator"
	"GolangDecoratorPattern/model"
	"fmt"
	"time"
)

func main() {
	initiallDate := time.Now()

	for i := 0; i < 12; i++ {
		holiday := getDayWhitDecorators(initiallDate, i)

		_, _ = show(holiday)
	}

}

func show(holiday *decorator.Holiday) (int, error) {
	working := " "
	if !holiday.IsWorkingDay() {
		working = "X"
	}
	return fmt.Printf("|\t %v | %v |\n", holiday.GetDate().Format("Mon | _2/01/2006"), working)
}

func getDayWhitDecorators(initial time.Time, days int) *decorator.Holiday {
	day := model.NewDay(initial.AddDate(0, 0, days), true)

	dayOfWeek := decorator.NewDayOfWeek(day)

	holiday := decorator.NewHoliday(dayOfWeek)
	return holiday
}
