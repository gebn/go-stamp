OUT := vanilla
LDFLAGS := -ldflags=" \
-X 'github.com/gebn/go-stamp/v2.User=$(USER)' \
-X 'github.com/gebn/go-stamp/v2.Host=$(shell hostname)' \
-X 'github.com/gebn/go-stamp/v2.timestamp=$(shell date +%s)' \
-X 'github.com/gebn/go-stamp/v2.Commit=$(shell git rev-parse HEAD)' \
-X 'github.com/gebn/go-stamp/v2.Branch=$(shell git rev-parse --abbrev-ref HEAD)' \
-X 'github.com/gebn/go-stamp/v2.Version=$(shell git describe --always --tags --dirty)'"

build:
	go build $(LDFLAGS) -o $(OUT)

clean:
	rm -f $(OUT)
