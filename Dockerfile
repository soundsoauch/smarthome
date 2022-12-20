FROM golang:1.19-alpine

WORKDIR /usr/src/app

# ssh for iris shutdown
RUN apk add openssh
RUN mkdir -p /root/.ssh
ADD .ssh/ /root/.ssh/
RUN chmod 700 /root/.ssh

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app
EXPOSE 8000/tcp
CMD ["app"]