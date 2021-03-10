package decorator

import "time"

type WorkDay interface {
	IsWorkingDay() bool
	GetDate() time.Time
}
