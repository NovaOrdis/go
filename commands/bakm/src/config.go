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
func (config config) init(args []string) error {

    if len(args) < 2 {
        return errors.New("no target directory specified")
    }

    config.dirName = args[1]

    var err error

    config.dirPtr, err = os.Open(config.dirName)

    if err != nil {
        return err
    }

    return nil
}
