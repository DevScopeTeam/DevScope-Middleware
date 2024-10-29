# Compile stage
FROM golang:1.21.3 AS build-env
 
ADD . /build
WORKDIR /build

RUN go env -w GO111MODULE=on &\
    go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o /server
 
# Final stage
FROM debian:bookworm
 
EXPOSE 9000
 
WORKDIR /app
COPY --from=build-env /server /app/server
COPY ./config/config.yaml /app/config/config.yaml
COPY ./config/rsa_private_key.pem /app/config/rsa_private_key.pem
 
CMD ["/app/server"]