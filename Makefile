VPATH := github.com/sethlittleford/budget-bot/version
commit_hash := $(shell git rev-parse HEAD)
branch := $(shell git rev-parse --abbrev-ref HEAD)
build_flags := -ldflags "-X $(VPATH).commitHash=$(commit_hash) -X $(VPATH).branch=$(branch)"

build:
	make clean
	make test
	go build $(build_flags)

clean:
	go mod tidy

test:
	go test -v $(build_flags) ./...