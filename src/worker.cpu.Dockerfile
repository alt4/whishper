FROM golang:alpine AS builder

WORKDIR /app

COPY . .
RUN apk add upx git && go install github.com/a-h/templ/cmd/templ@latest && \
    templ generate ./... && \
    go mod tidy && \
    GOOS=linux go build -o anysub . && \
    upx anysub && \
    chmod a+rx anysub

RUN chmod a+rx anysub

FROM pluja/whisperx-api:cpu

COPY --from=builder /app/anysub /usr/bin/anysub
COPY worker.entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]