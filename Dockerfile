FROM golang:1.10 as build-deps

RUN mkdir -p /go/src/github.com/uc-cdis/indexs3client
WORKDIR /go/src/github.com/uc-cdis/indexs3client
ADD . .
RUN go build -ldflags "-linkmode external -extldflags -static" -o bin/indexs3client

# Store only the resulting binary in the final image
# Resulting in significantly smaller docker image size
FROM scratch
COPY --from=build-deps /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-deps /go/src/github.com/uc-cdis/indexs3client/bin/indexs3client /indexs3client

CMD ["/indexs3client"]
