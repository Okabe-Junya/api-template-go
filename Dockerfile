FROM golang:1.22

# install cmake
RUN apt-get update && apt-get install -y cmake

WORKDIR /app

# install go modules
COPY go.mod ./
COPY go.sum ./

RUN go mod download

# Copy source code
COPY . .

# build
RUN make build

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the binary
CMD ["/app/bin/server"]
