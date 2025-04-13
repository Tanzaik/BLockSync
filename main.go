package main

import (
    "blocksync/sync"
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: blocksync <bucket-name> <filepath>")
        os.Exit(1)
    }
    bucket := os.Args[1]
    filepath := os.Args[2]

    err := sync.SyncFile(bucket, filepath)
    if err != nil {
        fmt.Println("Sync error:", err)
        os.Exit(1)
    }
    fmt.Println("Sync complete.")
}
