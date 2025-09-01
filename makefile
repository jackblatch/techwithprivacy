.PHONY: build-static clean

build-static:
	go run ./cmd/staticgen

clean:
	rm -f index.html

