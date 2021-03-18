package main

import (
	"GolangDecoratorPattern/decorator"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Go!")

	initiallDate := time.Now()

	for i := 0; i < 12; i++ {
		holiday := getDayWhitDecorators(initiallDate, i)
		_ , _ = show(holiday)
	}

}

func show(holiday *decorator.Holiday) (int, error) {
	working := " "
	if !holiday.IsWorkingDay() {
		working = "X"
	}
	return fmt.Printf("| %v | %v | %v |\n",
		convertWeekday(holiday.GetDate().Weekday()),
		holiday.GetDate().Format("_2/01/2006"), working)
}



func convertWeekday(w time.Weekday)  string {
	var weekdays = [...]string{
		"Dom",
		"Lun",
		"Mar",
		"Mie",
		"Jue",
		"Vie",
		"Sab",
	}

	return weekdays[w]
}

func getDayWhitDecorators(initial time.Time, days int) *decorator.Holiday {
	day := decorator.NewDay(initial.AddDate(0, 0, days), true)

	dayOfWeek := decorator.NewDayOfWeek(day)

	holiday := decorator.NewHoliday(dayOfWeek)
	return holiday
}
