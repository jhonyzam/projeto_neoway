FROM golang

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/projeto_neoway

COPY . .

RUN go get github.com/lib/pq
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 8080 to the outside world
EXPOSE 3000 
