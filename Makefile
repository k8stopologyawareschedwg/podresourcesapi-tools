all: client fakeserver

win: client-win fakeserver-win

.PHONY: clan
clean:
	rm -rf _out

.PHONY: clean-deps
clean-deps:
	rm -rf vendor

.PHONY: update-deps
update-deps:
	go mod tidy && go mod vendor

.PHONY: update-version
update-version:
	@mkdir -p pkg/version || :
	@hack/make-version.sh > pkg/version/version.go

client-static: outdir update-version
	CGO_ENABLED=0 go build -o _out/client ./cmd/client

client-win: outdir update-version
	GOOS=windows go build -o _out/client.exe ./cmd/client

fakeserver-static: outdir update-version
	CGO_ENABLED=0 go build -o _out/fakeserver ./cmd/fakeserver

fakeserver-win: outdir update-version
	GOOS=windows go build -o _out/fakeserver.exe ./cmd/fakeserver

client: outdir update-version
	go build -o _out/client ./cmd/client/

fakeserver: outdir update-version
	go build -o _out/fakeserver ./cmd/fakeserver/

outdir:
	@mkdir -p _out || :

.PHONY: test-unit
test-unit:
	go test ./pkg/...

.PHONY: gofmt
gofmt:
	@echo "Running gofmt"
	gofmt -s -w `find . -path ./vendor -prune -o -type f -name '*.go' -print`

.PHONY: vet
vet:
	go vet ./cmd/... ./pkg/...
