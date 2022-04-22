FROM golang:1.17.9-alpine3.15

RUN apk add --update make

WORKDIR /go/src/myip

COPY . .

RUN go get -d -v ./... 

RUN go install -v ./... 

RUN make build

VOLUME [ "/file" ]

CMD [ "dist/myip", "-f", "/file" ]