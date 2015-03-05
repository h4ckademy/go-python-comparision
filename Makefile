go:
	@go run reddit.go

install-go-deps:
	@cd go
	@go get ./...

python:
	@python reddit.py
