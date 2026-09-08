FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/santosh .

FROM alpine:3.20
COPY --from=build /bin/santosh /usr/local/bin/santosh
EXPOSE 8080
ENTRYPOINT ["santosh"]
