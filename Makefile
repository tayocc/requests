GIT_TAG := v0.0.9
DATE_NOW := $(shell date "+%Y%m%d_%H%M%S")


.PHONY: start test git
all: start test git


start:
	go mod tidy


test:
	go test -v


git:
	git add .
	git commit -am "update_${DATE_NOW}"
	git tag "${GIT_TAG}"
	git push --tags

