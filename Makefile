GO=go
all= $(BINARIES)
BINARIES: vol

clean:
	$(GO) clean
	rm -f $(BINARIES)

vol: vol.go
	$(GO) build -o volume vol.go

