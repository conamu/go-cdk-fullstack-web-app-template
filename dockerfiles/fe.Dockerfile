FROM golang:latest as build-env

WORKDIR /app
COPY . .
RUN go build -o ./bin/frontend-app ./src/app/frontend/main.go
ADD https://busybox.net/downloads/binaries/1.31.0-i686-uclibc/busybox_WGET /wget
RUN chmod a+x /wget

FROM gcr.io/distroless/base
WORKDIR app/
COPY --from=build-env /app/bin/frontend-app .
COPY --from=build-env /wget /usr/bin/wget
HEALTHCHECK --interval=10s --timeout=3s CMD ["/usr/bin/wget", "--no-verbose","--tries=1", "--spider", "http://localhost/"]
EXPOSE 80
CMD ["./frontend-app"]