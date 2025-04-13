package sync

import (
    "blocksync/s3"
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "io"
    "os"
    "path/filepath"
)

const BlockSize = 4096

type Manifest struct {
    Blocks []string `json:"blocks"`
}

func SyncFile(bucket, path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()

    fileName := filepath.Base(path)
    manifestKey := fmt.Sprintf("manifests/%s.json", fileName)

    // Get remote manifest
    manifest, err := s3.DownloadManifest(bucket, manifestKey)
    if err != nil {
        manifest = &Manifest{}
    }

    newManifest := &Manifest{}
    buf := make([]byte, BlockSize)

    for {
        n, err := file.Read(buf)
        if err != nil && err != io.EOF {
            return err
        }
        if n == 0 {
            break
        }

        hash := sha256.Sum256(buf[:n])
        hashStr := hex.EncodeToString(hash[:])
        newManifest.Blocks = append(newManifest.Blocks, hashStr)

        if !contains(manifest.Blocks, hashStr) {
            err := s3.UploadBlock(bucket, hashStr, buf[:n])
            if err != nil {
                return err
            }
        }
    }

    err = s3.UploadManifest(bucket, manifestKey, newManifest)
    return err
}

func contains(list []string, item string) bool {
    for _, x := range list {
        if x == item {
            return true
        }
    }
    return false
}
