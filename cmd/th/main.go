// Command th is a single-binary command line for threads.
package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/charmbracelet/fang"
	"github.com/tamnd/threads-cli/cli"
	"github.com/tamnd/threads-cli/threads"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	root := cli.Root()
	// fang gives styled help, errors, and shell completion for free; the command
	// tree lives in the cli package, and a library CodeError maps to its exit code.
	if err := fang.Execute(ctx, root,
		fang.WithVersion(cli.Version),
		fang.WithNotifySignal(os.Interrupt, syscall.SIGTERM),
	); err != nil {
		var ce *threads.CodeError
		if errors.As(err, &ce) {
			os.Exit(ce.Code)
		}
		os.Exit(1)
	}
}
