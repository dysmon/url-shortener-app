FROM golang:1.15 AS build-stage
WORKDIR /api
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /url-shortener ./cmd/api

FROM gcr.io/distroless/base-debian12:latest
WORKDIR /
COPY --from=build-stage /url-shortener /url-shortener
EXPOSE 8080
USER nonroot:nonroot
CMD ["/url-shortener"]
