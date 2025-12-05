//go:build !windows && !darwin

package cmd

import (
	"context"
	"errors"

	"github.com/glowllama/glowllama/api"
)

func startApp(ctx context.Context, client *api.Client) error {
	return errors.New("could not connect to glowllama server, run 'glowllama serve' to start it")
}
