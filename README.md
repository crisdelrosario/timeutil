# timeutil
--
    import "github.com/crisdelrosario/timeutil"


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
GetLocaTime get local time based on timezone. Sets 'Asia/Sinagpore' as the
default timezone if not set

#### func (*TimeUtil) Now

```go
func (tu *TimeUtil) Now() string
```
Now Get the system's current date and time
