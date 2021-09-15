FROM golang:1.16-alpine

# Set destination for COPY
WORKDIR /app

ENV GOPROXY https://goproxy.cn
# Download Go modules
COPY go.mod .
COPY go.sum .
#RUN go get github.com/billzhou2005/th-service/controllers
#RUN go get github.com/billzhou2005/th-service/models
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./

# Build
RUN go build -o /th-service

# This is for documentation purposes only.
# To actually open the port, runtime parameters
# must be supplied to the docker command.
EXPOSE 8080

# (Optional) environment variable that our dockerised
# application can make use of. The value of environment
# variables can also be set via parameters supplied
# to the docker command on the command line.
#ENV HTTP_PORT=8081

# Run
CMD [ "/th-service" ]
