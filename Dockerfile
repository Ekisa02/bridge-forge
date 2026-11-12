FROM golang:1.20-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://proxy.golang.org,direct
COPY . ./
RUN go build -o /usr/local/bin/nebula-bridge ./cmd/server

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=build /usr/local/bin/nebula-bridge /usr/local/bin/nebula-bridge
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/nebula-bridge"]
