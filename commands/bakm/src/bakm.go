package main

import (
    "os"
    "fmt"
)

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
