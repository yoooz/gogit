FROM golang:1.15.10-alpine3.13

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN apk update && \
    apk add git

COPY ./ ./gogit/
COPY startup.sh .

RUN cd gogit && go install -v .

RUN chmod 744 gogit/startup.sh

CMD ["/go/src/app/gogit/startup.sh"]
