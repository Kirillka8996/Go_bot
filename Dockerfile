FROM golang

WORKDIR /app

COPY . .

RUN ["go","get","-u","github.com/go-telegram-bot-api/telegram-bot-api/v5"]

CMD [ "go","run","main.go" ]