FROM golang:1.15-alpine as builder

RUN apk update && apk upgrade
RUN apk --no-cache --update add git

RUN go get -u github.com/sirupsen/logrus
RUN go get -u github.com/snowzach/rotatefilehook
RUN go get -u cloud.google.com/go/firestore
RUN go get -u firebase.google.com/go
RUN go get -u google.golang.org/api/iterator
RUN go get -u google.golang.org/api/option

WORKDIR /var/legaiabay/mikroblog-be

COPY . .

RUN go build -o app .

FROM golang:1.15-alpine

COPY --from=builder  /var/legaiabay/mikroblog-be/app /

CMD /app

EXPOSE 4423