VERSION := $(shell cat VERSION)
GITLAB_TOKEN := USh-vxTV9MXHswbMU6m6


default: build

test:
	go test ./...

build:
	go build

install: build
	mkdir -p ~/.tflint.d/plugins/github.com/knowbe4/tflint-ruleset-kb4/$(VERSION)/
	mv ./tflint-ruleset-kb4 ~/.tflint.d/plugins/github.com/knowbe4/tflint-ruleset-kb4/$(VERSION)/

release:
	git tag v$(VERSION)
	git push origin v$(VERSION)


gorelease:
	rm -rf dist
	GITLAB_TOKEN=$(GITLAB_TOKEN) goreleaser