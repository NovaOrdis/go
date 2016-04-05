package main

import (
	"flag"
	"fmt"
	"os"
)

const DEFAULT_KEEP_COUNT = 10

type config struct {
	help          bool // if "true" it means we print the help and get out
	helpRequested bool // if help is true and helpRequested is true, it means the help was explicitly
	// requested so we exit with 0, otherwise we're looking at a syntax error and
	// we exit with 1

	dirName   string   // the name of the target directory as specified in the command line
	dirPtr    *os.File // the target directory, if opened successfully
	keepCount int
}

// initializes a blank config instance from command line arguments or with defaults
// if command line arguments values are not available
func (config *config) init(args []string) error {

	var err error

	if len(args) < 2 {

		config.help = true
		config.helpRequested = false
		return nil
	}

	for _, arg := range args {

		if "help" == arg || "--help" == arg || "-h" == arg || "-help" == arg {
			config.help = true
			config.helpRequested = true
			return nil
		}
	}

	//
	// parse flags
	//

	keepCountPtr := flag.Int("keep", DEFAULT_KEEP_COUNT, "the number of files to keep in the directory")

	flag.Parse()

	config.keepCount = *keepCountPtr

	args = flag.Args()

	if len(args) == 0 {
		// no target directory
		fmt.Println("no target directory specified")
		os.Exit(1)
	}

	config.dirName = args[0]

	config.dirPtr, err = os.Open(config.dirName)

	if err != nil {
		return err
	}

	return nil
}
