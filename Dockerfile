# BUILD STAGE
FROM golang:1.17-alpine as builder

RUN apk --no-cache add ca-certificates git

WORKDIR /build

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 go build -o myserver

# post build stage
FROM alpine
WORKDIR /server
COPY ./config/config.yml .
COPY --from=builder /build/myserver .
CMD ["./myserver"]