FROM golang:1.12 as build

WORKDIR /go/src/gitlab.com/eAuction/enumer
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/enumer gitlab.com/eAuction/enumer/cmd/enumer

#
#
#

FROM alpine:3.9

COPY --from=build /usr/local/bin/enumer /usr/local/bin/enumer
RUN apk add --no-cache ca-certificates curl bash

ENTRYPOINT ["enumer"]
