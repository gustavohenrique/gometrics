FROM golang:alpine
COPY . $GOPATH/src/github.com/gustavohenrique/example
ENV GOPROXY off
WORKDIR $GOPATH/src/github.com/gustavohenrique/example
RUN apk add --update git \
 && go get github.com/gustavohenrique/gometrics \
 && go build -o example main.go \
 && mv example /usr/local/bin
CMD ["example"]
