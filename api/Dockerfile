# golang image where workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/moehandi/campaign/api

# Setting up working directory
WORKDIR /go/src/github.com/moehandi/campaign/api

# Get godeps for managing and restoring dependencies
RUN go get github.com/tools/godep

# Restore godep dependencies
RUN godep restore 

# Build the taskmanager command inside the container.
RUN go install github.com/moehandi/campaign/api

# Run the taskmanager command when the container starts.
ENTRYPOINT /go/bin/campaign/api

# Service listens on port 3000.
EXPOSE 3000
