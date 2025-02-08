FROM golang:1.22-alpine AS builder

ARG ENVIRONMENT=development

ENV ENV=${ENVIRONMENT}
ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR /backend-ejakti-ticket

COPY . .

# COPY environment-specific .env*
COPY .env.${ENV} .env.application

RUN apk update && apk add --no-cache gcc libc-dev && \
    go version && go mod download && go mod verify && \
    CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=1.0.0 -X main.buildTime=$(date +%Y-%m-%d) -s -w" -o ./backend-ejakti-ticket

#
FROM alpine:3.18

ENV TZ=Asia/Jakarta

WORKDIR /backend-ejakti-ticket

RUN apk update && \
    apk add --no-cache tzdata

COPY --from=builder /backend-ejakti-ticket/backend-ejakti-ticket  .
COPY --from=builder /backend-ejakti-ticket/.env.application .env
COPY --from=builder /backend-ejakti-ticket/go.mod go.mod

RUN chmod +x ./backend-ejakti-ticket

EXPOSE 80
EXPOSE 443 

CMD ["./backend-ejakti-ticket"]