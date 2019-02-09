all: install

.PHONY: install
install:
	go install skeletor/skeletor.go

.PHONY: clean
clean:
	rm -rf resources/test-skeletor/skeletons 
.PHONY: test
test: install
	cd resources/test-skeletor && skeletor update-skeleton test