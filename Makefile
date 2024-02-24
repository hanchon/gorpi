.PHONY:  desktop

rpi:
	@go run cmd/rpi/main.go

desktop:
	@go run cmd/desktop/main.go
