build: 
	go build -o bin/bookstore -v

run: build
	./bin/bookstore