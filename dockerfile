#BUILDER
FROM golang:latest as builder

WORKDIR /app
ARG CGO_ENABLED=0 #CGO Disable for scratch

# Download module in a separate layer to allow caching for the Docker build
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o witb

# RUNNING
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/witb /witb
ENTRYPOINT ["/witb"]
