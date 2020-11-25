FROM golang:latest
# Download app dependencies
RUN go get github.com/joho/godotenv
RUN go get github.com/go-chi/chi
RUN go get github.com/jmoiron/sqlx
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/guregu/null
RUN go get -u go.uber.org/zap

RUN mkdir /app
ADD ./src /app/
ADD ./.env /app/
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]