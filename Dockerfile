# Build a executable file
FROM golang:1.18 as builder

# Build arguments
ARG UID

# Create user user 
RUN useradd -ms /bin/bash -u ${UID} code
RUN usermod -G root code
WORKDIR $GOPATH/src/code/app/
USER code
