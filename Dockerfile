FROM golang:1.14 as build

WORKDIR /go/src/gitlab.com/eAuction/buying-power-svc
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/buying-power gitlab.com/eAuction/buying-power-svc/cmd/buying-power

#
#
#

FROM ubuntu:16.04

COPY --from=build /usr/local/bin/buying-power /usr/local/bin/buying-power
RUN apt-get update \
	&& apt-get install -y \
		curl \
		xvfb \
		jq \
		wget \
		ca-certificates \
    && wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add - \
    && echo "deb http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google.list \
    && apt-get update \
    && apt-get -y install google-chrome-stable

ENTRYPOINT ["buying-power"]
