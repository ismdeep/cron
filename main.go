package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	cronSpec := os.Getenv("CRON_SPEC")
	cronCommand := os.Getenv("CRON_COMMAND")

	if cronSpec == "" {
		panic("CRON_SPEC is empty")
	}

	if cronCommand == "" {
		panic("CRON_COMMAND is empty")
	}

	c := cron.New(
		cron.WithSeconds(),
		cron.WithChain(
			cron.SkipIfStillRunning(
				cron.DefaultLogger)),
	)
	_, _ = c.AddFunc(cronSpec, func() {
		fmt.Println("----------------------------------------------------------------------")
		startedAt := time.Now()
		fmt.Println("Started At:", startedAt.Format(time.RFC1123Z))

		defer func() {
			endedAt := time.Now()
			fmt.Println("Ended At:", endedAt.Format(time.RFC1123Z))
			fmt.Println("Time Elapse:", fmt.Sprintf("%v s", endedAt.Unix()-startedAt.Unix()))
			fmt.Println("----------------------------------------------------------------------")
			fmt.Println()
		}()

		cmdSlice := strings.Fields(cronCommand)
		cmd := exec.Command(cmdSlice[0], cmdSlice[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("Error:", err.Error())
			return
		}
		fmt.Println("OK")
	})

	c.Run()
}
