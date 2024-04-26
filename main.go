package main

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/ismdeep/log"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

//go:embed banner.txt
var banner string

var cronCommand string

func runCommandFunc() {
	startedAt := time.Now()
	log.WithContext(context.Background()).Info("Started")
	defer func() {
		endedAt := time.Now()
		log.WithContext(context.Background()).Info("Ended", zap.String("time_elapsed", endedAt.Sub(startedAt).String()))
		log.WithContext(context.Background()).Info("----------------------------------------------------------------------")
		fmt.Println()
	}()
	cmdSlice := strings.Fields(cronCommand)
	cmd := exec.Command(cmdSlice[0], cmdSlice[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.WithContext(context.Background()).Error("failed to run task", zap.Error(err))
		return
	}
	log.WithContext(context.Background()).Info("Finished")
}

func main() {
	fmt.Println(banner)

	cronSpec := os.Getenv("CRON_SPEC")
	cronCommand = os.Getenv("CRON_COMMAND")
	cronRunAtStart := os.Getenv("CRON_RUN_AT_START")

	if cronSpec == "" {
		panic("CRON_SPEC is empty")
	}

	if cronCommand == "" {
		panic("CRON_COMMAND is empty")
	}

	log.WithContext(context.Background()).Info("----------------------------------------------------------------------")

	if cronRunAtStart == "true" {
		runCommandFunc()
	}

	c := cron.New(
		cron.WithSeconds(),
		cron.WithChain(
			cron.SkipIfStillRunning(
				cron.DefaultLogger)),
	)

	for _, spec := range strings.Split(cronSpec, "\n") {
		spec = strings.TrimSpace(spec)
		if spec != "" {
			log.WithContext(context.Background()).Info("Cron spec", zap.String("spec", spec))
			if _, err := c.AddFunc(spec, runCommandFunc); err != nil {
				panic(fmt.Errorf("failed to add cron spec '%s': %v", spec, err))
			}
		}
	}

	c.Run()
}
