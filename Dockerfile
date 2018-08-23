FROM golang:1.8

RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY ./templates/ /go/src/app/templates/
COPY ./resources/ /go/src/app/resources/

RUN go get -v github.com/pe5er/radiotutor

EXPOSE 8080

CMD ["radiotutor"]
