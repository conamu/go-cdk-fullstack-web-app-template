FROM golang:latest as build-env

WORKDIR /app
COPY . .
RUN go build -o ./bin/backend-app ./src/app/backend/main.go
ADD https://busybox.net/downloads/binaries/1.31.0-i686-uclibc/busybox_WGET /wget
RUN chmod a+x /wget

FROM gcr.io/distroless/base
WORKDIR app/
COPY --from=build-env /app/bin/backend-app .
COPY --from=build-env /wget /usr/bin/wget
HEALTHCHECK --interval=10s --timeout=3s CMD ["/usr/bin/wget", "--no-verbose","--tries=1", "--spider", "http://localhost:8080/"]
EXPOSE 8080
CMD ["./backend-app"]