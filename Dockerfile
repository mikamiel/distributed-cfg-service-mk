FROM alpine:3.16.2

COPY cfg-service-mk /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/cfg-service-mk"]

EXPOSE 50051