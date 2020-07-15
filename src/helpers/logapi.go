package helpers

import (
	"time"
	"fmt"
)

func LogApi(info string)  {

	p := fmt.Println
	now := time.Now()
	p("Year", now.Year())
	p("Month", now.Month())
	p("Day", now.Day())
	p("Hour", now.Hour())
	p("Minute", now.Minute())
	p("Second", now.Second())
	p("Nanosecond", now.Nanosecond())
	p("Location", now.Location())
	p("Weekday", now.Weekday())
	p(info);
	p("---------------------")
}