FROM golang:latest 

ADD ./server/cert.pem /etc/ssl/certs/
ADD ./server/key.pem /etc/ssl/certs/
RUN echo Copied keys

RUN go get -u -v github.com/vladimir-voinea/location-report-app/server 
WORKDIR /go/src/github.com/vladimir-voinea/location-report-app/server 
RUN go build
EXPOSE 50051

ENTRYPOINT [ "server" ]
CMD ["-cert_file=/etc/ssl/certs/cert.pem", "-key_file=/etc/ssl/certs/key.pem"]
