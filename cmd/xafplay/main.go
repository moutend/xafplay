package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

// afplay is a macOS built-in command.
const afplay = "afplay"

func main() {
	if err := run(); err != nil {
		log.New(os.Stderr, "error: ", 0).Fatal(err)
	}
}

func run() error {
	if len(os.Args) < 2 {
		return nil
	}

	paths := os.Args[1:]

	signalChan := make(chan os.Signal, 1)

	// Capture SIGINT signal.
	signal.Notify(signalChan, syscall.SIGINT)

	isPlaying := true
	now := time.Now()

	var ctx context.Context
	var cancel context.CancelFunc

	go func() {
		for {
			select {
			case <-signalChan:
				duration := time.Now().Sub(now)
				now = time.Now()

				if duration <= time.Second {
					// When received the Ctrl-C within 1 second, stop playback. Otherwise, play the next music.
					isPlaying = false
				}

				cancel()
			}
		}
	}()

	for _, path := range paths {
		if !isPlaying {
			break
		}

		fmt.Println("Playing", filepath.Base(path))

		ctx, cancel = context.WithCancel(context.Background())

		if err := exec.CommandContext(ctx, afplay, path).Run(); err != nil && !isInterruptError(err) {
			return err
		}

		cancel()
	}

	return nil
}

func isInterruptError(err error) bool {
	return strings.Contains(err.Error(), "signal: interrupt")
}
