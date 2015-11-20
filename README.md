# fakeserver
fakeserver creates a big web structure useful to test http bots or crawlers

## Docker container
A Docker container of fakeserver is available. Get it using:
```
docker pull jvgutierrez/fakeserver
```
## Build instructions with alpine docker image
1. CGO_ENABLED=0 go build -a -installsuffix cgo
2. docker build -t fakeserver .
