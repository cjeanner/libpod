package main

import (
	"github.com/urfave/cli"
)

var (
	imageSubCommands = []cli.Command{
		buildCommand,
		historyCommand,
		importCommand,
		imageExistsCommand,
		inspectCommand,
		loadCommand,
		lsImagesCommand,
		pruneImagesCommand,
		pullCommand,
		pushCommand,
		rmImageCommand,
		saveCommand,
		tagCommand,
		trustCommand,
	}

	imageDescription = "Manage images"
	imageCommand     = cli.Command{
		Name:                   "image",
		Usage:                  "Manage images",
		Description:            imageDescription,
		ArgsUsage:              "",
		Subcommands:            imageSubCommands,
		UseShortOptionHandling: true,
		OnUsageError:           usageErrorHandler,
	}
)
