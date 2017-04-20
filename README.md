# timeutil
--
    import "github.com/crisdelrosario/timeutil"

This is just a helper function for getting the current time, eitther based on
timezone provided or just the default time.

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

## Usage

#### func  TimeMs

```go
func TimeMs() int64
```
TimeMs Returns Unix time in milliseconds for benchmarking Svc performance.

#### type Elapse

```go
type Elapse struct {
	StartTime int64
}
```

Elapse elapse time computation utility

#### func (*Elapse) Start

```go
func (elapse *Elapse) Start()
```
Start marks current time as start

#### func (*Elapse) Stop

```go
func (elapse *Elapse) Stop() int64
```
Stop marks the current time as end then returns the computed elapsed time

#### type TimeUtil

```go
type TimeUtil struct {
	TZ     string
	Elapse Elapse
}
```

TimeUtil ...

#### func  New

```go
func New(timezone string) TimeUtil
```
New create new TimeUtil instance

#### func (*TimeUtil) CurrentDate

```go
func (tu *TimeUtil) CurrentDate() string
```
CurrentDate Get system's current date

#### func (*TimeUtil) CurrentTime

```go
func (tu *TimeUtil) CurrentTime() string
```
CurrentTime Get system's current time

#### func (*TimeUtil) GetLocalTime

```go
func (tu *TimeUtil) GetLocalTime() time.Time
```
GetLocalTime get local time based on timezone. Sets 'Asia/Sinagpore' as the
default timezone if not set

#### func (*TimeUtil) Now

```go
func (tu *TimeUtil) Now() string
```
Now Get the system's current date and time
