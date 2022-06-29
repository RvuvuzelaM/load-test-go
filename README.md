# Load Test Go

## How to use

1. Download loadtest binary from repo
2. If you want to get help execute this command
```bash
./loadtest rest --help
```
3. Example of the command
```bash
./loadtest rest -e https://httpbin.org/post -m POST -n 1000 -w 50 --headers "{\"Authorization\": \"xxx\"}" -b "body"
```

## Documentation

```
Allows one to test REST API capabilites by bombarding with requests.
What is recomended is to first make test request to your API, and then load test it.

Usage:
  loadtest rest [flags]

Flags:
  -b, --body string             request body (if applies)
  -e, --endpoint string         endpoint to hit (default "https://www.example.com")
      --headers json            headers for requests (default null)
  -h, --help                    help for rest
  -m, --method string           HTTP request method (default "GET")
  -n, --num-of-requests int     overall number of requests (default 1)
  -w, --number-of-workers int   number of concurrent workers for requests (default 1)
```