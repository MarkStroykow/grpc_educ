FROM golang:1.23-alpine AS builder

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download


COPY cmd cmd
COPY internal internal

RUN GOOS=linux go build -o app ./cmd

FROM golang:1.23-alpine

WORKDIR /root/
COPY --from=builder /build/app .
EXPOSE 50052
CMD ["./app"]