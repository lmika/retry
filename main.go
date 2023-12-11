package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	ctx := context.Background()

	flagSleep := flag.Int("s", 2, "sleep between retries (0 = no sleep)")
	flagMaxAttempt := flag.Int("n", 5, "max number of retries (0 = no max)")

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(2)
	}

	startTime := time.Now()
	currentAttempt := 0
	for {
		currentAttempt += 1
		if err := runExec(ctx, flag.Args()...); err == nil {
			break
		}

		if *flagMaxAttempt > 0 && currentAttempt >= *flagMaxAttempt {
			log.Printf("execution failed: reached max %d attempts, duration = %v", currentAttempt, time.Since(startTime).String())
			os.Exit(1)
		}

		if *flagMaxAttempt > 0 {
			log.Printf("execution failed: retrying (attempt %d of %d)", currentAttempt+1, *flagMaxAttempt)
		} else {
			log.Printf("execution failed: retrying (attempt %d)", currentAttempt+1)
		}

		if *flagSleep > 0 {
			time.Sleep(time.Duration(*flagSleep) * time.Second)
		}
	}

	if currentAttempt > 1 {
		log.Printf("execution succeeded after %d attempts, duration = %v", currentAttempt, time.Since(startTime).String())
	}
}

func runExec(ctx context.Context, cmdAndArgs ...string) error {
	cmd := exec.CommandContext(ctx, cmdAndArgs[0], cmdAndArgs[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
