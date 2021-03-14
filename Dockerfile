FROM golang:1.15.10-alpine3.13

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN apk update && \
    apk add git

COPY ./ ./gogit/
COPY startup.sh .

RUN cd /go/src/app/gogit && git checkout 
RUN cd /go/src/app/gogit && go install -v .

CMD ["/go/src/app/gogit/startup.sh"]
