app = gitlabctl
version = $$(cat VERSION.md|awk '{print $$1}')
install_global = "/usr/local/bin"

.PHONY: help
help:
	
	$(info Makefile Usage:)
	$(info build 			- build the go binary)
	$(info test 			- run unit tests)
	$(info install 		- install binary on the OS)
	$(info uninstall 		- uninstall the binary)
	@printf "\n"

.PHONY: install
install:
	@(if dk version >/dev/null 2>&1; then \
	docker run --rm --name $(app)-install -v $(PWD):/go/src/$(app) -v $(install_global):/tmp golang:1.12-alpine3.9 ash -c \
		"(cd /go/src/$(app);apk add git build-base;go get -v ; \
			go build -ldflags \"-X $(app)/cmd.Version=$(version)\" -tags netgo -a -installsuffix cgo -o /tmp/$(app) .)"; \
	echo $(app) installed on $(install_global); \
	else \
		go install; \
		echo $(app) installed on $(GOBIN);fi)

.PHONY: uninstall
uninstall: 
	@(if dk version >/dev/null 2>&1; then \
	docker run --rm --name $(app)-uninstall -v /usr/local/bin:/tmp golang:1.12-alpine3.9 ash -c \
		"(cd /tmp;rm -rf $(app))"; \
	echo $(app) removed from $(install_global); \
	else \
		rm -f $(GOBIN)/$(app); \
		echo $(app) unistalled from $(GOBIN);fi)

.PHONY: test
test:
	@(go test ./... -v -cover)

.PHONY: build
build:
	@(go build)


