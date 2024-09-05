FROM golang:1.23 AS build

WORKDIR /src
COPY . .
RUN go mod download
RUN go build -o /bin/web ./cmd/web/

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=build /bin/web /bin/web
CMD ["/bin/web"]
