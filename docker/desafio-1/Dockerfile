FROM golang:1.18.4-alpine AS builder

WORKDIR /app/
COPY hello.go .
RUN go build hello.go

FROM scratch
COPY --from=builder /app/hello /bin/hello
ENTRYPOINT ["/bin/hello"]
