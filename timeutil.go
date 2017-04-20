package timeutil

type TimeUtil struct {
	TZ string
}

func New(timezone string) TimeUtil {
	return TimeUtil{
		TZ: timezone,
	}
}

// // LocalTime get local time given the timezone
// func (timeutil *TimeUtil) LocalTime(timezone string) time.Time {
// 	if timezone == "" {
// 		panic("Server timezone not set!")
// 	}
//
// 	serverTime := time.Now()
// 	locationTime, _ := time.LoadLocation(timezone)
// 	currentTime := serverTime.In(locationTime)
// 	return currentTime
// }
//
// // GetCurrentDate Get system's current date
// func (timeutil *TimeUtil) GetCurrentDate(timezone string) string {
// 	return LocalTime(timezone).Format("2006-01-02")
// }
//
// // GetCurrentTime Get the system's current date and time
// func (timeutil *TimeUtil) GetCurrentTime(timezone string) string {
// 	return LocalTime(timezone).Format("2006-01-02 15:04:05")
// }
//
// // TimeMs Returns Unix time in milliseconds for benchmarking Svc performance.
// func (timeutil *TimeUtil) TimeMs() int64 {
// 	return time.Now().UnixNano() / 1000000
// }
