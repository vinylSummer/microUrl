FROM --platform=linux/amd64 golang:1.22-bullseye
RUN mkdir /app
COPY ./go.mod ./go.sum /app/
WORKDIR /app
RUN go mod download
COPY . /app
EXPOSE 8080
CMD go run ./cmd/app/main.go
