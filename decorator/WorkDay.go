package decorator

import "time"

/*
	Decorator
*/
type WorkDay interface {
	IsWorkingDay() bool
	GetDate() time.Time
}
