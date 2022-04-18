FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /dist ./app/main.go 

EXPOSE 3222

CMD ["/dist"]
