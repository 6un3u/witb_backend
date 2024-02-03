#BUILDER
FROM golang:latest as builder

WORKDIR /app
ARG CGO_ENABLED=0 #CGO Disable for scratch

COPY . .
RUN go mod download

RUN go build -o witb

# RUNNING
FROM scratch
COPY --from=builder /app/witb /witb
ENTRYPOINT ["/witb"]
