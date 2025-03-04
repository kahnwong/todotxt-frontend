FROM golang:1.23-alpine AS build

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go build -ldflags "-w -s" -o /todotxt-frontend

# hadolint ignore=DL3007
FROM gcr.io/distroless/static-debian11:latest AS deploy

# hadolint ignore=DL3045
COPY static ./static
COPY --from=build /todotxt-frontend /

EXPOSE 3000
CMD ["/todotxt-frontend"]
