package main

import (
	"fmt"
	"os"
	"spun/cmd/scheduler/job"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/helloferdie/golib/v2/liblogger"
	"github.com/helloferdie/golib/v2/libsignal"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	loc, err := time.LoadLocation(os.Getenv("app_timezone"))
	if err != nil {
		loc, _ = time.LoadLocation("Asia/Jakarta")
	}

	opt := []gocron.SchedulerOption{
		gocron.WithLocation(loc),
	}

	sch, _ := gocron.NewScheduler(opt...)
	defer func() {
		sch.Shutdown()
		liblogger.Infow("Scheduler - Stop")
	}()

	// Set scheduler job
	// job.ExampleHello5Secs(sch)
	// job.ExampleHello1Sec(sch)
	// job.ExampleDelay(sch)
	job.ExampleHelloAtTime(sch)

	sch.Start()
	liblogger.Infow("Scheduler - Start")

	fmt.Println("To exist press CTRL + C")
	libsignal.WaitCTRLC(false)
}
