FROM golang:1.21.3 AS builder

# ENV GOPROXY=https://goproxy.io,direct # For China

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /usr/local/bin/app ./cmd/server

FROM scratch

WORKDIR /app

USER 1500:1500

COPY --chown=1500:1500 --from=builder /usr/local/bin/app /app/server

CMD [ "/app/server" ]
