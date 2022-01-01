build:
	go build -v -o bin/ .

lint-gomod:
	go mod tidy
	git diff --exit-code go.mod
	git diff --exit-code go.sum