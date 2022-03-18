FROM golang:1.18-alpine
WORKDIR /go/src/paintings-shop
ENV DATABASE_HOST db
COPY . . 
RUN go get -d -v ./...
RUN apk update
RUN apk add ffmpeg-dev build-base
WORKDIR /go/src/paintings-shop/cmd/app
RUN go build -o $GOPATH/bin/paintings-shop
RUN rm -rf /go/src/*
WORKDIR $GOPATH/bin
COPY ./cmd/app/public ./public
COPY ./cmd/app/logs ./logs
COPY ./cmd/app/settings.json ./settings.json
EXPOSE 10443
EXPOSE 8080
CMD ["./paintings-shop", "-clean", "-makedb", "-noresetroles"]
