# go-utils
Small, independent Go utility packages for common application tasks.

## Packages

- `aws`: AWS session, S3, and DynamoDB helpers.
- `bufio`: buffered reader helpers.
- `byte`: Gob serialization compatibility helper.
- `csv`: CSV reading and error-returning writing helpers.
- `digest`: SHA-256 and SHA-512 helpers for bytes, readers, and files.
- `file`: file readers, gzip CSV extraction, atomic writes, and file copies.
- `http`: HTTP requests, downloads, and health-check handlers.
- `math`: random string and byte generation.
- `memory`: memory statistics helpers.
- `retry`: context-aware retries with configurable backoff.
- `time`: delay helpers.

## Examples

```go
err := retry.Do(ctx, retry.Config{
	MaxAttempts: 3,
	InitialDelay: 100 * time.Millisecond,
	MaxDelay: 2 * time.Second,
}, operation)
```

```go
checksum := digest.SHA256([]byte("payload"))
randomValue, err := math.RandomString(16)
```

I/O and HTTP helpers return operational errors to the caller. HTTP responses outside the 2xx range return a `*http.ResponseError` containing the status code and response body. The legacy `byte.Hash` and `math.Random` functions remain available for compatibility.
