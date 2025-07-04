package commands

import (
	"github.com/urfave/cli/v3"

	"github.com/EnvSync-Cloud/envsync-cli/internal/features/handlers"
)

// ConfigCommands returns all config-related commands
func ConfigCommands(handler *handlers.ConfigHandler) *cli.Command {
	return &cli.Command{
		Name:  "config",
		Usage: "Manage configuration settings",
		Commands: []*cli.Command{
			SetCommand(handler),
			GetCommand(handler),
			ResetCommand(handler),
		},
	}
}

func SetCommand(handler *handlers.ConfigHandler) *cli.Command {
	return &cli.Command{
		Name:      "set",
		Usage:     "Set configuration values",
		Action:    handler.Set,
		ArgsUsage: "key=value [key2=value2 ...]",
		Description: `Set one or more configuration values.

Examples:
  envsync config set access_token=your_token_here
  envsync config set backend_url=https://api.envsync.cloud
  envsync config set access_token=token backend_url=https://api.example.com

Supported keys:
  - access_token: Authentication token for EnvSync Cloud
  - backend_url: Backend API URL`,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "overwrite",
				Usage: "Overwrite existing configuration completely",
				Value: false,
			},
		},
	}
}

func GetCommand(handler *handlers.ConfigHandler) *cli.Command {
	return &cli.Command{
		Name:      "get",
		Usage:     "Get configuration values",
		Action:    handler.Get,
		ArgsUsage: "[key1] [key2] ...",
		Description: `Get configuration values. If no keys are specified, all configuration is shown.

Examples:
  envsync config get
  envsync config get access_token
  envsync config get access_token backend_url

Supported keys:
  - access_token: Authentication token for EnvSync Cloud
  - backend_url: Backend API URL`,
	}
}

func ResetCommand(handler *handlers.ConfigHandler) *cli.Command {
	return &cli.Command{
		Name:      "reset",
		Usage:     "Reset configuration values",
		Action:    handler.Reset,
		ArgsUsage: "[key1] [key2] ...",
		Description: `Reset configuration values to defaults. If no keys are specified, all configuration is reset.

Examples:
  envsync config reset
  envsync config reset access_token
  envsync config reset access_token backend_url

WARNING: This action cannot be undone. Consider backing up your configuration first.`,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "force",
				Usage: "Skip confirmation prompt",
				Value: false,
			},
			&cli.BoolFlag{
				Name:  "backup",
				Usage: "Create backup before resetting",
				Value: true,
			},
		},
	}
}
