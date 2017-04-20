package timeutil

// Elapse elapse time computation utility
type Elapse struct {
	startTime int64
}

// Start marks current time as start
func (elapse *Elapse) Start() {
	elapse.startTime = TimeMs()
}

// Stop marks the current time as end then returns the computed elapsed time
func (elapse *Elapse) Stop() int64 {
	return TimeMs() - elapse.startTime
}
