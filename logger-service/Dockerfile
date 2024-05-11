FROM golang:1.22-alpine

RUN apk --no-cache add bash git make gettext

COPY ./ ./

RUN go mod download

RUN go build -o ./auth-service cmd/main.go && \
    chmod +x entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]