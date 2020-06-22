package helpers

import (
	"time"
)

func GetCurrentLocalTime() time.Time {
	//init the loc
	loc, _ := time.LoadLocation("Europe/Minsk")

	//set timezone,  
	now := time.Now().In(loc)
	return now
}
