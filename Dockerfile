FROM golang:1.24.2 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN CGOOS=linux CGO_ENABLED=0 go build -v -o /app/api/ ./cmd/api/main.go

EXPOSE 8080

# Distroless image for small size
FROM gcr.io/distroless/static-debian11 as final

WORKDIR /app
COPY --from=builder --chmod=777 /app/api ./
#COPY --from=builder --chmod=777 /app/config ./config

ENTRYPOINT ["/app/main"]
