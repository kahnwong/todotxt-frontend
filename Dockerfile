FROM golang:1.23-alpine AS build

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY todo ./todo
COPY *.go ./

RUN CGO_ENABLED=0 go build -ldflags "-w -s" -o /todotxt-frontend

# hadolint ignore=DL3007
FROM alpine:latest AS deploy

# hadolint ignore=DL3045
COPY frontend/dist/spa /frontend/dist/spa/
COPY --from=build /todotxt-frontend /

EXPOSE 3000
CMD ["/todotxt-frontend"]
