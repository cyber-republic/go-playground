FROM go-playground-sandbox-basic:latest

RUN go get -u -v github.com/cyber-republic/go-grpc-adenine/elastosadenine

USER elauser

EXPOSE 8080

ENTRYPOINT ["/bin/sh", "-c", "./playground"]
