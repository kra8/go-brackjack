BUILD_DIR 	:= bin
NAME 		:= brackjack

.PHONY: build
build:
	@go build -o $(BUILD_DIR)/$(NAME)

.PHONY: run
run:
	@./$(BUILD_DIR)/$(NAME)

.PHONY: test
test: build run

.PHONY: clean
clean:
	rm -rf ./bin
