#go model 的版本号
Version := v0.0.10

#git commit 的内容
Time := $(shell date "+%Y%m%d_%H%M%S")


.PHONY: start test git
all: start test git


start:
	go mod tidy


test:
	go test -v


git:
	git add .
	git commit -am "update_${Time}"
	git tag "${Version}"
	git push --tags

