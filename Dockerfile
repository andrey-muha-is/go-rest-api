FROM golang:latest
# Download app dependencies
RUN go get github.com/go-chi/chi
RUN go get github.com/go-chi/render

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]