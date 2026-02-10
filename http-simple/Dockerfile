FROM golang:1.25-alpine AS builder

WORKDIR /usr/local/src

COPY go.mod main.go ./

RUN go build

FROM scratch AS release

COPY --from=builder /usr/local/src/http-simple /http-simple

ENTRYPOINT [ "/http-simple" ]
