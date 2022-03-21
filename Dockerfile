FROM golang:1.18-alpine
EXPOSE 10443
EXPOSE 8080
WORKDIR /go/src/paintings-shop
ENV DATABASE_HOST db
RUN apk add --no-cache ffmpeg-dev build-base git
COPY . . 
RUN go get -d -v ./...
WORKDIR /go/src/paintings-shop/cmd/app
RUN go build -o $GOPATH/bin/paintings-shop
RUN rm -rf /go/src/*
WORKDIR $GOPATH/bin
COPY ./cmd/app/public ./public
COPY ./cmd/app/logs ./logs
COPY ./cmd/app/settings.json ./settings.json
CMD ["./paintings-shop", "-clean", "-makedb"]
