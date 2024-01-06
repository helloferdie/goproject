package job

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
	"github.com/helloferdie/golib/v2/liblogger"
)

// Print "hello" every 5 secs
func ExampleHello5Secs(sch gocron.Scheduler) {
	sch.NewJob(gocron.DurationJob(time.Second*5),
		gocron.NewTask(
			func() {
				fmt.Println("Hello every 5 seconds")
			},
		), gocron.WithName("Hello5Sec"), LogStartStop())
}

// Print "hello" every 1 secs, with limit 2 singleton mode
func ExampleHello1Sec(sch gocron.Scheduler) {
	sch.NewJob(gocron.DurationJob(time.Second),
		gocron.NewTask(
			func() {
				fmt.Println("Hello every 1 second")
				time.Sleep(time.Second * 10)
			},
		), gocron.WithName("Hello1Sec Singleton"), gocron.WithSingletonMode(2), LogStartStop())
}

// Print "hello" with delay 15 seconds
func ExampleDelay(sch gocron.Scheduler) {
	startTime := time.Now().Add(time.Second * 15)
	sch.NewJob(gocron.DurationJob(time.Second),
		gocron.NewTask(
			func() {
				fmt.Println("Hello every 1 second - start with delay 15 seconds")
			},
		), gocron.WithStartAt(gocron.WithStartDateTime(startTime)))
}

// Print "hello" at 18:30
func ExampleHelloAtTime(sch gocron.Scheduler) {
	sch.NewJob(gocron.DailyJob(1, gocron.NewAtTimes(
		gocron.NewAtTime(18, 30, 0),
	)),
		gocron.NewTask(
			func() {
				fmt.Println("Hello at 18:30")
			},
		))
}

// LogStartStop
func LogStartStop() gocron.JobOption {
	return gocron.WithEventListeners(gocron.BeforeJobRuns(
		func(jobID uuid.UUID, jobName string) {
			liblogger.Infof("Job %s - %s - start", jobName, jobID)
		},
	), gocron.AfterJobRuns(func(jobID uuid.UUID, jobName string) {
		liblogger.Infof("Job %s - %s - stop", jobName, jobID)
	}))
}
