FROM xiam/go-playground-unsafebox:latest

ENV HOME_DIR="/home/elauser"

WORKDIR $HOME_DIR

RUN useradd -d $HOME_DIR elauser && \
    chown -R elauser:elauser $HOME_DIR

RUN go get github.com/cyber-republic/go-grpc-adenine/elastosadenine

USER elauser


