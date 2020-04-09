pretty:
	gofmt -s -w .

mod:
	go mod tidy

build:
	go build -o bin/mygoweb cmd/web/main.go

run: pretty mod build
	./bin/mygoweb

deploy:
	sudo cp mygoweb.service /lib/systemd/system/mygoweb.service

# Only for development
dev: build
	./bin/mygoweb
