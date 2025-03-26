# Build Stage
FROM golang AS build

WORKDIR /SPADE
COPY . ./
RUN go test

# Install Stage
