GOCACHE ?= $(CURDIR)/.gocache

.PHONY: run serve clean

run:
	@GOCACHE=$(GOCACHE) go run ./cmd/tutor

serve:
	@scripts/tts-setup.sh
	@GOCACHE=$(GOCACHE) scripts/serve.sh

clean:
	@rm -rf $(GOCACHE)
