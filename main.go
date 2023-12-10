package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/exec"
)

func main() {
	ctx := context.Background()

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(2)
	}

	for {
		if err := runExec(ctx, flag.Args()...); err == nil {
			break
		}

		log.Printf("execution failed: retrying...")
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
