FROM golang:1.23 AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o app .

FROM gcr.io/distroless/static-debian12

WORKDIR /
COPY --from=build /app/app /app

EXPOSE 8080
ENTRYPOINT ["/app"]