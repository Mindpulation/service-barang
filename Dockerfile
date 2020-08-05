FROM golang as builder

# Add Maintainer Info
LABEL maintainer="Muhammad Indrawan <me@aboutindra.com>"

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/mindpulation/Kotakjualan-Barang

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix .

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /go/src/github.com/mindpulation/Kotakjualan-Barang /app/Kotakjualan-Barang
WORKDIR "/app/Kotakjualan-Barang"
EXPOSE 6554
ENTRYPOINT ./kjualan-barang

