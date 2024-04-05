FROM golang:1.22.0-alpine

# 作業ディレクトリ
WORKDIR /go/src/application

# ファイルを作業ディレクトリにコピー
COPY . .

RUN go mod download

# git導入（軽量イメージのalpineには含まれないため）
RUN apk upgrade --update && apk add git

# 公開するポートを指定する（これがないとアクセスできない）
EXPOSE 8080

# time zone databaseを取ってくる
ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip

# up -dしたら、go run main.goまでやる
CMD [ "go", "run", "main.go" ]