# This Dockerfile is for references only.
# To build Docker image, please use corresponding Bazel rule.
FROM golang:1.21 AS build-stage

ENV GOPRIVATE "github.com/vucong2409/sample-private-go-mod"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /docker-gs-ping /docker-gs-ping

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/sample-go-bzl-ch"]
