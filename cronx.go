// Copyright 2025 Focela Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0.
// See the LICENSE file in the project root for full license information.

// Package main provides a cron job scheduler with graceful shutdown.
package main

import (
	"context"
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
	// minArgs is the minimum required command-line arguments.
	minArgs = 3
)

// execute runs command with args, redirecting stdout/stderr.
func execute(command string, args []string) {
	fmt.Printf("executing: %s %s\n", command, strings.Join(args, " "))

	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("command failed: %v\n", err)
	}
}

// create initializes cron scheduler that respects ctx cancellation.
func create(ctx context.Context, schedule string, command string, args []string) (*cron.Cron, *sync.WaitGroup) {
	wg := &sync.WaitGroup{}

	// Supports optional seconds and descriptors (@daily, @weekly).
	parser := cron.NewParser(
		cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)

	if _, err := parser.Parse(schedule); err != nil {
		fmt.Printf("invalid schedule: %v\n", err)
		os.Exit(1)
	}

	c := cron.New(cron.WithParser(parser))
	fmt.Printf("new cron: %s\n", schedule)

	c.AddFunc(schedule, func() {
		wg.Add(1)
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			execute(command, args)
		}
	})

	return c, wg
}

// stop shuts down scheduler and waits for running jobs to complete.
func stop(c *cron.Cron, wg *sync.WaitGroup) {
	fmt.Println("Stopping")
	c.Stop()
	fmt.Println("Waiting for running jobs to complete")
	wg.Wait()
	fmt.Println("Exiting")
	os.Exit(0)
}

// main parses arguments and runs cron scheduler with signal handling.
func main() {
	if len(os.Args) < minArgs {
		fmt.Println("Usage: cronx [schedule] [command] [args ...]")
		os.Exit(1)
	}

	schedule := os.Args[1]
	command := os.Args[2]
	args := os.Args[3:]

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c, wg := create(ctx, schedule, command, args)

	c.Start()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	fmt.Printf("Received signal: %v\n", sig)

	cancel()
	stop(c, wg)
}
