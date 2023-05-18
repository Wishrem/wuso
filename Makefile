PWD = $(shell pwd)
BUILD_PATH = $(PWD)/build
TEST_PATH = $(PWD)/test
CONFIG_PATH = $(PWD)/config

COMPONENTS := server client
component = $(word 2, $(MAKECMDGOALS))


.PHONY: clean
clean:
	rm -rf $(BUILD_PATH)

.PHONY: build
build:
	go build -o $(BUILD_PATH)/$(component)  $(PWD)/$(component)

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: test
test:
	go test -v $(TEST_PATH)

.PHONY: run
run:
	make build $(component)
	$(BUILD_PATH)/$(component) -config $(CONFIG_PATH)

.PHONY: env-up
env-up:
	docker-compose up -d

.PHONY: env-down
env-down:
	docker-compose down