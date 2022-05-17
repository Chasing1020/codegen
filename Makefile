run:
	rm -rf dist && go run main.go && cd dist && make fmt