.PHONY: all build clean cross

all: clean build

build:
	scripts/build.sh

cross:
	scripts/cross.sh

clean:
	rm -rf build/
