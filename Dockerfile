# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.20 AS build-stage
ARG VERSION
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X 'main.Version=$VERSION'" -o /httpinfo

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /httpinfo /httpinfo

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/httpinfo"]