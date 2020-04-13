FROM golang:latest

RUN mkdir -p /usr/src/app/
RUN mkdir -p $GOPATH/src/github.com/Ubidy
WORKDIR /usr/src/app

COPY . /usr/src/app/

RUN ln -s /usr/src/app $GOPATH/src/github.com/Ubidy/Ubidy_AgencyNotificationAPI

RUN go get github.com/auth0-community/auth0
RUN go get github.com/denisenkom/go-mssqldb
RUN go get github.com/gin-gonic/gin
RUN go get github.com/google/uuid
RUN go get github.com/gin-contrib/cors
RUN go get gopkg.in/GetStream/stream-go2.v1
RUN go get github.com/Microsoft/ApplicationInsights-Go/appinsights
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/square/go-jose.v2
RUN go build -o agencyactivitystreamapi .

EXPOSE 5021 80
ENTRYPOINT [ "./agencyactivitystreamapi" ]
