package main

import (
    "os"
    "errors"
)

type config struct {
    dirName string
    dirPtr *os.File
}

// initializes a blank config instance from command line arguments
func (config *config) init(args []string) error {

    var err error

    if len(args) < 2 {
        return errors.New("no target directory specified")
    }

    config.dirName = args[1]

    config.dirPtr, err = os.Open(config.dirName)

    if err != nil {
        return err
    }

    return nil
}
