FROM golang:alpine3.20 AS build

# Set the working directory to /app
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download and install any required Go dependencies
RUN --mount=type=cache,target=go/pkg/mod \
--mount=type=cache,target=/root/.cache/go-build \
 go mod download

CMD [ "make", "seed" ]

# Copy the entire source code to the working directory
COPY . .

# Build the Go application
RUN go build \
#  -ldflags="-linkmode external -extldflags -static" \
 -tags netgo \ 
 -o main 

# Multistage build
FROM scratch

# Copy the .env file into the image
COPY .env .env


ENV GOFIBER_MODE=release

COPY --from=build /app/main main

# Expose the port specified by the PORT environment variable
EXPOSE 8000

# Set the entry point of the container to the executable
CMD ["./main"]

