FROM golang:1.18-alpine3.16

RUN apk update && apk upgrade && apk add curl \
                          bash \
                          make \
                         busybox-extras  && \
     rm -rf /var/cache/apk/*

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN make build

EXPOSE 8080

ENTRYPOINT ["./main"]
