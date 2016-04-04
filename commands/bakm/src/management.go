package main

import (
    "os"
    "fmt"
    "sort"
    "time"
)

//
// this is a type that associates the file name with its modification time and
// helps with sorting
//
type fileWithModificationTime struct {
    basename string
    modificationTime time.Time
}

//
// a sortable type, it implements
//
type byModificationTime []fileWithModificationTime

func (bmt byModificationTime) Len() int {
    return len(bmt)
}

func (bmt byModificationTime) Swap(i, j int) {
    bmt[i], bmt[j] = bmt[j], bmt[i]
}

func (bmt byModificationTime) Less(i, j int) bool {
    // this sorts in the descending order, from the newest to the oldes
    return bmt[i].modificationTime.After(bmt[j].modificationTime)
}

// returns a slice containing the names of the deleted files
func deleteAllExceptMostRecentOnes(dirPtr *os.File, keepCount int) ([]string, error) {

    // we start with an empty slice, if no files are deleted, it will be returned as such
    deleteFileNames := make([]string, 0)

    // scan the directory
    fileInfos, err := dirPtr.Readdir(-1)

    if err != nil {
        return deleteFileNames, err
    }

    files := make(byModificationTime, 0)

    for _, fileInfo := range fileInfos {

        // skip directories
        if fileInfo.IsDir() {
            continue
        }

        files = append(files, fileWithModificationTime{fileInfo.Name(), fileInfo.ModTime()})
    }

    //
    // if there are less than 'keepCount', it does not make sense to sort, we'll keep them all
    //
    if len(files) <= keepCount {

        return deleteFileNames, nil
    }

    //
    // sort in place in the descending order of the modification time, the newest files at the top
    //
    sort.Sort(files)

    for i, fwmt := range files {

        if i < keepCount {
            // it's new, we keep it
            continue
        }

        fileName := dirPtr.Name() + "/" + fwmt.basename
        //fmt.Printf("removing %s\n", fileName)
        err = os.Remove(fileName)
        if err != nil {
            fmt.Errorf("failed to remove %s\n", fileName)
            return deleteFileNames, err
        }
        deleteFileNames = append(deleteFileNames, fwmt.basename)

    }

    return deleteFileNames, nil
}

