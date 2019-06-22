FROM gcr.io/gcp-runtimes/go1-builder:1.12 as builder

LABEL maintainer="vikram.jakhr@gmail.com"
LABEL go_version="1.12"
LABEL description="Diginex microservice 1"

WORKDIR /go/src/app
COPY . .

RUN /usr/local/go/bin/go build -o app .

# Application image.
FROM gcr.io/distroless/base:latest

COPY --from=builder /go/src/app/app /usr/local/bin/app

CMD ["/usr/local/bin/app"]