# Define your target platforms and architecture
GOOS_LIST = linux darwin windows
GOARCH = amd64
DIST_DIR = dist
BUILD_DIR = build

# Name of the executable
BINARY_NAME = PoGoPass

# Ensure all build directories are created
$(shell mkdir -p $(DIST_DIR))

# Build command for each OS
build: $(GOOS_LIST)

$(GOOS_LIST):
	echo "Building for $@..."
	GOOS=$@ GOARCH=$(GOARCH) go build -o $(BUILD_DIR)/$(BINARY_NAME)-$@ cmd/PoGoPass/main.go
	if [ -f $(BUILD_DIR)/$(BINARY_NAME)-$@ ]; then \
		echo "Build for $@ completed."; \
	else \
		echo "Build for $@ failed."; \
		exit 1; \
	fi

# Package command for each OS
package: build
	echo "Packaging for all OS..."
	@for os in $(GOOS_LIST); do \
		zip -r $(DIST_DIR)/$(BINARY_NAME)-$$os.zip $(BUILD_DIR)/$(BINARY_NAME)-$$os; \
		echo "Packaged $$os."; \
	done

# Clean up build files
clean:
	echo "Cleaning up build directories..."
	rm -rf $(BUILD_DIR)
	rm -rf $(DIST_DIR)
	echo "Cleaned up."

# Help command for usage instructions
help:
	echo "Usage:"
	echo "  make build    # Build for all platforms"
	echo "  make package  # Package builds into .zip files"
	echo "  make clean    # Clean build directories"
	echo "  make help     # Show this help message"
