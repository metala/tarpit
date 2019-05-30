FROM golang:1.12-alpine as builder
WORKDIR /app
COPY . /app
RUN apk add --no-cache git make \
    && make linux-amd64

# Runtime
FROM scratch
COPY --from=builder /app/build/tarpit-linux-amd64 /tarpit
USER 65534:65534
EXPOSE 2022
ENTRYPOINT ["/tarpit"]
CMD ["-p", "2022"]
