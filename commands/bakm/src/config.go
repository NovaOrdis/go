package main

import (
    "os"
    "errors"
)

const DEFAULT_KEEP_COUNT = 10

type config struct {
    dirName string // the name of the target directory as specified in the command line
    dirPtr *os.File // the target directory, if opened successfully
    keepCount int
}

// initializes a blank config instance from command line arguments or with defaults
// if command line arguments values are not available
func (config *config) init(args []string) error {

    var err error

    if len(args) < 2 {
        return errors.New("no target directory specified")
    }

    //
    // defaults
    //

    config.keepCount = DEFAULT_KEEP_COUNT

    //
    // command line arguments
    //

    config.dirName = args[1]

    config.dirPtr, err = os.Open(config.dirName)

    if err != nil {
        return err
    }

    return nil
}
