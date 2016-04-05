package main

import (
	"fmt"
	"os"
)

const HELP = `
Utility to clean a directory of old backup files. It removes as many files as necessary
(possibly none)  to leave the newest 'keepCount' files in the directory.   If less than
'keepCount' files are found in the directory, none is removed.

Usage:
		bakm [-keep=<keepCount>] <target-directory>

		bakm -help|--help|help|-h

Example:

		bakm -keep=10 /opt/backup

Options

	-keep - specifies the number of files to keep in the directory (default 10).
`

func main() {

	var config config

	err := config.init(os.Args)

	if err != nil {

		// TODO - experiment with the best way to handle this (panic(), log and return, etc)
		// then annotate https://kb.novaordis.com/index.php/Go_Concepts_-_Error_Handling
		//panic(err)
		fmt.Println(err)
		return
	}

	if config.help {

		// display help
		fmt.Println(HELP)

		if config.helpRequested {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	// TODO parse the log package
	//log.Printf(".\n")

	//
	// delete everything from the directory, except most config.keep most recent ones
	//

	deletedFileNames, err := deleteAllExceptMostRecentOnes(config.dirPtr, config.keepCount)

	if err != nil {
		fmt.Println(err)
		return
	}

	if len(deletedFileNames) == 0 {

		fmt.Println("no files deleted")

	} else {

		fmt.Print("deleted ")
		for _, fileName := range deletedFileNames {
			fmt.Print(fileName + " ")
		}
		fmt.Println()
	}
}
