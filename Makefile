.PHONY: build
build:
	go build

.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: gen-docs
gen-docs: build
	@command -v cloudquery >/dev/null 2>&1 || { \
		echo "Error: 'cloudquery' command not found. Please install it before running gen-docs."; \
		echo "You can install it by following the instructions at: https://www.cloudquery.io/docs/quickstart"; \
		exit 1; \
	}
	rm -rf docs/tables
	cloudquery tables --format markdown --output-dir docs/ test/config.yml
	mv -vf docs/crowdstrike docs/tables

.PHONY: dist
dist:
	go run main.go package -m "Release ${VERSION}" ${VERSION} .

.PHONY: lint
lint:
	@golangci-lint run --timeout 10m --verbose

# All gen targets
.PHONY: gen
gen: gen-docs
