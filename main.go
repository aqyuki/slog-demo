package main

import (
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	// Output plane text
	slog.Info(
		"Hello World",
		slog.String("name", "blue"),
		slog.Int("no", 1),
	)

	// Output json text
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info(
		"Hello World",
		slog.String("name", "blue"),
		slog.Int("no", 1),
	)
}
