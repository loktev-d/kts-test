FROM golang:1.20.2-bullseye AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . ./
RUN CGO_ENABLED=0 go build -v -o app

FROM alpine:3.17.2
WORKDIR /app
COPY --from=build /src/app ./
CMD ["./app"]
