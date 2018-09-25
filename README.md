Deprecated. Use https://github.com/tinrab/kit instead.

# Retry

A simple retry package for Go.

## Usage

Retry 10 times.

```go
err := retry.Do(10, thisFunctionMayFail)
if err != nil {
  log.Fatal(err)
}
```

Retry with a delay.

```go
err := retry.DoSleep(10, 3 * time.Second, thisFunctionMayFail)
if err != nil {
  log.Fatal(err)
}
```

Retry forever.

```go
retry.Forever(thisFunctionMayFail)
```

