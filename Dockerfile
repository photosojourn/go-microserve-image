FROM golang:1.13-alpine AS build

#Install git
RUN apk add --no-cache git
RUN go get github.com/aws/aws-sdk-go
RUN go get github.com/photosojourn/go-microservice
RUN go get github.com/gorilla/handlers
WORKDIR /go/src/github.com/photosojourn/go-microservice
RUN go build -o /bin/go-microservice

FROM golang:1.13-alpine
COPY --from=build /bin/go-microservice /bin/go-microservice
EXPOSE 8080
ENTRYPOINT ["/bin/go-microservice"]