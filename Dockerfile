FROM golang:latest
ADD . /go/src/github.com/suderio/semver
WORKDIR /go/src/github.com/suderio/semver
RUN go install -v
ENTRYPOINT ["semver"]