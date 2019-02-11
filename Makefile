all: install

.PHONY: install
install:
	go install skeletor/skeletor.go

.PHONY: test
test:
	go test . -v