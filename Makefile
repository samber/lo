
build:
	go build -v ./...

test:
	go test -race ./...
watch-test:
	reflex -t 50ms -s -- sh -c 'gotest -race ./...'

bench:
	go test -v -run=^Benchmark -benchmem -count 3 -bench ./...
watch-bench:
	reflex -t 50ms -s -- sh -c 'go test -v -run=^Benchmark -benchmem -count 3 -bench ./...'

coverage:
	go test -v -coverprofile=cover.out -covermode=atomic ./...
	go tool cover -html=cover.out -o cover.html

tools:
	go install github.com/cespare/reflex@latest
	go install github.com/rakyll/gotest@latest
	go install github.com/psampaz/go-mod-outdated@latest
	go install github.com/jondot/goweight@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go get -t -u golang.org/x/tools/cmd/cover
	go install github.com/sonatype-nexus-community/nancy@latest
	go install golang.org/x/perf/cmd/benchstat@latest
	go install github.com/cespare/prettybench@latest
	go mod tidy

	# brew install hougesen/tap/mdsf

lint:
	golangci-lint run --timeout 60s --max-same-issues 50 ./...
	# mdsf verify --debug --log-level warn docs/
lint-fix:
	golangci-lint run --timeout 60s --max-same-issues 50 --fix ./...
	# mdsf format --debug --log-level warn docs/

audit:
	go list -json -m all | nancy sleuth

outdated:
	go list -u -m -json all | go-mod-outdated -update -direct

weight:
	goweight

doc:
	cd docs && npm install && npm start
