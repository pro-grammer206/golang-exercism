package bookingUpForBeauty

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    const shortForm = "1/2/2006 15:04:05"
	t,err:=time.Parse(shortForm,date)
	if err!=nil{
		fmt.Println(err)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	const longFormat = "January 2, 2006 15:04:05"
		t,err:=time.Parse(longFormat,date)
		if  err!=nil{
			fmt.Println(err)
		}
		Now:= time.Now()
		return Now.After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	const longFormat = "Monday, January 2, 2006 15:04:05"
	t,err:=time.Parse(longFormat,date)
	if err!=nil{
		fmt.Println(err)
	}
	if t.Hour()>=12 && t.Hour()<18{
		return true
	}else{
		return false
	}
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
const layout = "1/2/2006 15:04:05"
		t,err:=time.Parse(layout,date)
		if err!=nil{
			fmt.Println(err)
		}
		return fmt.Sprintf("You have an appointment on %v, %v %d, %d, at %d:%d.",t.Weekday(),t.Month(),t.Day(),t.Year(),t.Hour(),t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(),9,15,0,0,0,0,time.UTC)
}
