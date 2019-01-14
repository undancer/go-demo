

.PHONY: all
all: clean build

.PHONY: build
build:
	scripts/build.sh

.PHONY: cross
cross:
	scripts/cross.sh

.PHONY: clean
clean:
	rm -rf build/
