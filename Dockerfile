FROM golang:1.8.1

WORKDIR /go

RUN /bin/bash -c 'go get github.com/pkg/errors; \
go get github.com/aws/aws-sdk-go/aws; \
go get github.com/aws/aws-sdk-go/aws/credentials; \
go get github.com/aws/aws-sdk-go/aws/session; \
go get github.com/aws/aws-sdk-go/service/s3; \
go get github.com/chrisbenson/easyaws/pkg/easyaws; \
go get github.com/chrisbenson/jwt-go; \
go get github.com/justinas/alice; \
go get github.com/gorilla/mux; \
go get github.com/rs/cors; \
go get github.com/spf13/cobra; \
go get github.com/spf13/viper; \
go get github.com/lib/pq; \
go get github.com/satori/go.uuid; \
go get golang.org/x/crypto/scrypt; \
go get github.com/chrisbenson/epic; \
env GOOS=linux GOARCH=amd64 go build -o /go/bin/epic -v github.com/chrisbenson/epic;'

FROM phusion/baseimage:latest

WORKDIR /app/

COPY --from=0 /go/bin/epic .

RUN /bin/bash -c 'chmod +x /app/epic'

CMD ["/app/epic"]
