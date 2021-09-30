binary:
	go build -o go-i18n-linter ./cmd/go-i18n-linter

plg:
	go build -buildmode=plugin plugin/goi18nlinter.go

lint:
	golangci-lint run cmd/... pkg/...

test:
	go test ./pkg/...