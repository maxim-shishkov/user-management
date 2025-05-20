FROM golang:1.22 AS build

WORKDIR /app
COPY ../go.mod ../go.sum ./
RUN go mod tidy
COPY ../ ./
#COPY ../config/config.json /config/config.json
WORKDIR /app/cmd


RUN go build -o app .

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y libc6
COPY --from=build /app/cmd/app /usr/local/bin/app
COPY --from=build /app/config/config.json /config/config.json

EXPOSE 8080

CMD ["/usr/local/bin/app"]
