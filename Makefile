tests:
	@echo "Launching Tests in Docker Compose"
	docker-compose -f dev-compose.yml up --build tests
clean:
	@echo "Cleaning up build junk"
	-docker-compose -f dev-compose.yml down
	-rm -rf ./testdata

build:
	@echo "Building from source"
	go build

install:
	@echo "Installing from source"
	go install

release:
	@echo "Creating Release Binaries"
	go get -u github.com/mitchellh/gox
	mkdir -p build
	gox -output="build/{{.OS}}_{{.Arch}}/{{.Dir}}"