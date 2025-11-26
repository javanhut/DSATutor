GOCACHE ?= $(CURDIR)/.gocache

.PHONY: run serve clean

run:
	@GOCACHE=$(GOCACHE) go run ./cmd/tutor

serve:
	@GOCACHE=$(GOCACHE) go run ./cmd/tutor -serve -addr=:8080

clean:
	@rm -rf $(GOCACHE)
