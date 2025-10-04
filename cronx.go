// Copyright 2025 Focela Authors.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package main provides cronx, a cron job scheduler that executes commands at
// specified intervals. It supports standard cron expressions and gracefully
// handles shutdown signals.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/robfig/cron/v3"
)

const (
	// minArgs is the minimum number of command line arguments required.
	minArgs = 3
)

// execute runs the specified command with given arguments.
// It redirects stdout and stderr to the current process streams.
// Returns early if command execution fails, logging the error.
func execute(command string, args []string) {
	fmt.Printf("executing: %s %s\n", command, strings.Join(args, " "))

	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("command failed: %v\n", err)
	}
}

// create initializes a new cron scheduler with the specified schedule
// and command. It validates the cron expression and sets up the job
// execution with proper synchronization using a WaitGroup.
// Panics if the schedule is invalid.
func create(schedule string, command string, args []string) (*cron.Cron, *sync.WaitGroup) {
	wg := &sync.WaitGroup{}

	// Configure parser to support seconds, minutes, hours, day of month,
	// month, day of week, and descriptors (e.g., @daily, @weekly).
	parser := cron.NewParser(
		cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)

	if _, err := parser.Parse(schedule); err != nil {
		fmt.Printf("invalid schedule: %v\n", err)
		os.Exit(1)
	}

	c := cron.New(cron.WithParser(parser))
	fmt.Printf("new cron: %s\n", schedule)

	// Add job function with proper synchronization.
	c.AddFunc(schedule, func() {
		wg.Add(1)
		defer wg.Done()
		execute(command, args)
	})

	return c, wg
}

// start begins execution of the cron scheduler.
// This function is designed to be run in a goroutine.
func start(c *cron.Cron, wg *sync.WaitGroup) {
	c.Start()
}

// stop gracefully shuts down the cron scheduler and waits for running
// jobs to complete. It stops accepting new jobs, waits for current jobs
// to finish, then exits the process.
func stop(c *cron.Cron, wg *sync.WaitGroup) {
	fmt.Println("Stopping")
	c.Stop()
	fmt.Println("Waiting for running jobs to complete")
	wg.Wait()
	fmt.Println("Exiting")
	os.Exit(0)
}

// main is the entry point of the cron scheduler application.
// It parses command line arguments, creates a cron job, starts the
// scheduler, and handles graceful shutdown on SIGINT or SIGTERM signals.
func main() {
	if len(os.Args) < minArgs {
		fmt.Println("Usage: cronx [schedule] [command] [args ...]")
		os.Exit(1)
	}

	schedule := os.Args[1]
	command := os.Args[2]
	args := os.Args[3:]

	c, wg := create(schedule, command, args)

	// Start cron scheduler in background.
	go start(c, wg)

	// Set up signal handling for graceful shutdown.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	fmt.Printf("Received signal:1212121211212121212 %v\n", sig)

	stop(c, wg)
}
