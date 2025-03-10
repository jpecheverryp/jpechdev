# Build Stage
FROM golang:1.23 AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
RUN go install github.com/a-h/templ/cmd/templ@latest
COPY . .
RUN templ generate
RUN CGO_ENABLED=0 GOARCH=${TARGETARCH} GOOS=linux go build -o jpechdev ./cmd/web/
CMD ["/app/jpechdev"]

# Smaller Image for Deployment
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/jpechdev .
CMD ["/app/jpechdev"]
