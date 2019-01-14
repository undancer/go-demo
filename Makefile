

.PHONY: all
all: clean build

.PHONY: build
build:
	scripts/build.sh

.PHONY: clean
clean:
	rm -rf build/
