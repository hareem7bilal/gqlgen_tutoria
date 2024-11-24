# Stage 1: Dependencies
FROM golang:1.23-alpine as dependencies
# Use the lightweight Go image based on Alpine Linux for efficiency and small size.
# This stage is named 'dependencies' for reusability in subsequent stages.

WORKDIR /app
# Set the working directory inside the container to '/app'.

COPY go.mod go.sum ./
# Copy the dependency management files (go.mod and go.sum) into the working directory.
# These files are used by Go to track dependencies.

RUN go mod download
# Download and cache all Go module dependencies. 
# This ensures dependencies are resolved and cached in this stage.

# Stage 2: Build
FROM dependencies AS build
# Start a new stage that builds the application. Reuse the 'dependencies' stage as its base.

COPY . ./
# Copy all files from the current directory (source code) to the working directory in the container.

RUN CGO_ENABLED=0 go build -o /main -ldflags="-w -s" .
# Compile the Go application with:
# - CGO_ENABLED=0: Disables CGO to create a statically linked binary, making the binary more portable.
# - -o /main: Outputs the compiled binary as '/main'.
# - -ldflags="-w -s": Reduces the binary size by stripping debug information.

# Stage 3: Final Image
FROM golang:1.23-alpine
# Use the same lightweight Go Alpine image for the final container to ensure consistency and small size.

COPY --from=build /main /main
# Copy the compiled binary from the 'build' stage into the final image.

CMD [ "/main" ]
# Define the default command to execute the compiled binary.
# When the container starts, it will run '/main'.

