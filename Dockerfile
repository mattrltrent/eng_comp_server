# syntax=docker/dockerfile:1

FROM golang:1.20.5-bullseye
WORKDIR /app
COPY . .
RUN go build -o app
EXPOSE 8000
CMD ["/app"]
