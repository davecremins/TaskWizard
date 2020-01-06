package actions

import (
	"errors"
	"flag"
	"fmt"
	. "github.com/davecremins/ToDo-Manager/utilities"
)

func Process(args []string, defaultConfig *Config) error {

	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	filename := initCmd.String("filename", defaultConfig.Filename, "Name of file to initialise")

	newCmd := flag.NewFlagSet("newday", flag.ExitOnError)
	searchStr := newCmd.String("search", defaultConfig.SearchStr, "Search string to look for")
	daysToAdd := newCmd.Int("days", defaultConfig.DaysToAdd, "Total amount of days to increment by")

	if len(args) < 2 {
		return errors.New("expected subcommands to perform an action")
	}

	switch args[1] {

	case "init":
		initCmd.Parse(args[2:])
		fmt.Println("subcommand 'init'")
		fmt.Println("  filename:", *filename)
		fmt.Println("  tail:", initCmd.Args())
	case "newday":
		newCmd.Parse(args[2:])
		fmt.Println("subcommand 'newday'")
		fmt.Println("  searchStr:", *searchStr)
		fmt.Println("  days:", *daysToAdd)
		fmt.Println("  tail:", newCmd.Args())
	default:
		return errors.New(args[1] + " subcommand is not supported right now :(")
	}

	return nil
}
