FROM golang:1.24-bullseye

RUN apt-get update && apt-get install -y --no-install-recommends \
    git \
  && rm -rf /var/lib/apt/lists/*

WORKDIR /go-training

ADD . ./
RUN go mod download && go mod verify
ENV PATH="/go/bin:${PATH}"

