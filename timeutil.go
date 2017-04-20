/*
This is just a helper function for getting the current time, either based on timezone provided or just the default time.

Getting the current date and time example

			tu := timeutil.New("Asia/Singapore")
			fmt.Println("Current Date and Time: ", tu.Now())

Getting elapsed time

			tu.Elapse.Start()
			count := 1000000000
			size := count / 100
			x := 0
			y := 0
			for y < count {
				if (y % size) == 0 {
					fmt.Print(".")
				}

				for x < count {
					x++
				}
				y++
			}
			fmt.Println("")

			elapsed := tu.Elapse.Stop()
			measurement := "ms"
			if elapsed > 1000 {
				elapsed = elapsed / 1000
				measurement = "s"
			}

			fmt.Println("Elapsed Time: ", elapsed, measurement)

*/
package timeutil

import "time"

// TimeUtil ...
type TimeUtil struct {
	TZ     string
	Elapse Elapse
}

// New create new TimeUtil instance
func New(timezone string) TimeUtil {
	return TimeUtil{
		TZ:     timezone,
		Elapse: Elapse{},
	}
}

// GetLocalTime get local time based on timezone. Sets 'Asia/Sinagpore' as the default timezone if not set
func (tu *TimeUtil) GetLocalTime() time.Time {
	if tu.TZ == "" {
		tu.TZ = "Asia/Singapore"
	}

	serverTime := time.Now()
	locationTime, _ := time.LoadLocation(tu.TZ)
	return serverTime.In(locationTime)
}

// CurrentDate Get system's current date
func (tu *TimeUtil) CurrentDate() string {
	return tu.GetLocalTime().Format("2006-01-02")
}

// CurrentTime Get system's current time
func (tu *TimeUtil) CurrentTime() string {
	return tu.GetLocalTime().Format("15:04:05")
}

// Now Get the system's current date and time
func (tu *TimeUtil) Now() string {
	return tu.GetLocalTime().Format("2006-01-02 15:04:05")
}

// TimeMs Returns Unix time in milliseconds for benchmarking Svc performance.
func TimeMs() int64 {
	return time.Now().UnixNano() / 1000000
}
