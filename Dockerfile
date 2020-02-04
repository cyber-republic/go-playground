FROM xiam/go-playground-unsafebox:latest

ENV SRC_DIR="/home/elauser"

WORKDIR $SRC_DIR

RUN useradd -d $SRC_DIR elauser && \
    chown -R elauser:elauser $SRC_DIR

RUN go get github.com/cyber-republic/go-grpc-adenine/elastosadenine

USER elauser


