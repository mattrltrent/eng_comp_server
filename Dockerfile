# syntax=docker/dockerfile:1

FROM golang:1.20.5-bullseye
WORKDIR /app
COPY . .
RUN go install github.com/cosmtrek/air@latest
EXPOSE 8000
CMD ["air"]
