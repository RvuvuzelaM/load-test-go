# Load Test Go

1. Download loadtest binary from repo
2. If you want to get help execute this command
```bash
./loadtest rest --help
```
3. Example of the command
```bash
./loadtest rest -e https://httpbin.org/post -m POST -n 1000 -w 50
```