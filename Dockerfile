FROM golang:1.22-alpine AS builder

ARG ENVIRONMENT=development

ENV ENV=${ENVIRONMENT}
ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR /backend-product-service

COPY . .

# COPY environment-specific .env*
COPY .env.${ENV} .env.application

RUN apk update && apk add --no-cache gcc libc-dev && \
    go version && go mod download && go mod verify && \
    CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=1.0.0 -X main.buildTime=$(date +%Y-%m-%d) -s -w" -o ./backend-product-service

#
FROM alpine:3.18

ENV TZ=Asia/Jakarta

WORKDIR /backend-product-service

RUN apk update && apk add --no-cache tzdata

COPY --from=builder /backend-product-service/backend-product-service  .
COPY --from=builder /backend-product-service/.env.application .env
COPY --from=builder /backend-product-service/go.mod go.mod

RUN chmod +x ./backend-product-service

EXPOSE 80
EXPOSE 443 

CMD ["./backend-product-service"]