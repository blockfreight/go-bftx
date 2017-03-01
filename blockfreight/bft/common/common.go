package common

import (
    "fmt"
    "os"
    "io/ioutil"
)

func ReadJSON(path string) []byte {
    fmt.Println("\nReading "+path+"\n")
    file, e := ioutil.ReadFile(path)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    return file
}