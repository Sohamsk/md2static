.PHONY: setup test
setup:
	test -f ./external/antlr.jar || curl https://www.antlr.org/download/antlr-4.13.2-complete.jar -o ./external/antlr.jar
	go generate ./internal/parser/generate.go

test:
	go build -o ./dist/main.out ./cmd/main/.
	[ -d ./output/ ] || mkdir output
	./dist/main.out ./testfiles/test.md

clean:
	rm -rf ./output/ ./dist/
