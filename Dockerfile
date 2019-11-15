FROM golang:1.13.4-alpine3.10 as build
WORKDIR /go/src/app
COPY . .
RUN apk --no-cache add git
RUN go mod download
RUN sh bin/build

FROM alpine:3.7
COPY --from=build /go/src/app/app /usr/local/bin/app
ENTRYPOINT ["/usr/local/bin/app"]