FROM golang:1.20-alpine AS dependencies

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM dependencies AS build
COPY . ./
RUN CGO_ENABLED=0 go build -o /main -ldflags="-s -w" .

FROM golang:1.20-alpine
COPY --from=build /main /main
CMD ["/main"]
