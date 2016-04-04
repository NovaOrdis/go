package main

import (
    "os"
    "log"
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
    log.Printf(".\n")


}
