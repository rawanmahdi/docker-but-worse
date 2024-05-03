# syntax=docker/dockerfile:1

FROM golang:1.21.5

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-docker-but-worse

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
# EXPOSE 8080

# Run

# ENTRYPOINT ["/docker-docker-but-worse"]

# CMD ["/docker-docker-but-worse", "run", "echo", "hello world"]
# CMD ["/docker-dopscker-but-worse"]
# CMD ["/docker-docker-but-worse", "run", "ls", "/proc"]
# CMD ["/docker-docker-but-worse", "run", "ls", "/proc/self"]
# CMD ["/docker-docker-but-worse", "run", "ls", "/"]
CMD ["/docker-docker-but-worse", "run", "ps"]



# To build the image, run the following command in the directory where the Dockerfile is located:
# docker build --tag docker-docker-but-worse .

# To run the image, execute the following command:
# $ docker run --cap-add=SYS_ADMIN docker-docker-but-worse run hostname 
# $ docker run --cap-add=SYS_ADMIN docker-docker-but-worse