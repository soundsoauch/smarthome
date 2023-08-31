FROM golang:1.19-alpine

WORKDIR /usr/src/app

# ssh for iris shutdown
RUN apk add openssh
RUN mkdir -p /root/.ssh
ADD .ssh/ /root/.ssh/
RUN chmod 700 /root/.ssh

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading>
COPY smarthome/ .
RUN go build -v -o /usr/local/bin/app
EXPOSE 8000/tcp
CMD ["app"]
