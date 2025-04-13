package s3

import (
    "blocksync/sync"
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

var svc = s3.New(session.Must(session.NewSession(&aws.Config{
    Region: aws.String("us-west-1"), // Change as needed
})))

func UploadBlock(bucket, hash string, data []byte) error {
    key := fmt.Sprintf("blocks/%s", hash)
    _, err := svc.PutObject(&s3.PutObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(key),
        Body:   bytes.NewReader(data),
    })
    return err
}

func UploadManifest(bucket, key string, m *sync.Manifest) error {
    data, err := json.Marshal(m)
    if err != nil {
        return err
    }
    _, err = svc.PutObject(&s3.PutObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(key),
        Body:   bytes.NewReader(data),
    })
    return err
}

func DownloadManifest(bucket, key string) (*sync.Manifest, error) {
    obj, err := svc.GetObject(&s3.GetObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(key),
    })
    if err != nil {
        return nil, err
    }
    defer obj.Body.Close()

    body, err := io.ReadAll(obj.Body)
    if err != nil {
        return nil, err
    }

    var m sync.Manifest
    err = json.Unmarshal(body, &m)
    if err != nil {
        return nil, err
    }
    return &m, nil
}
