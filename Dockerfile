FROM golang:1.22.0-alpine

# 作業ディレクトリ
WORKDIR /go/src/application

# ファイルを作業ディレクトリにコピー
COPY . .

# git導入（軽量イメージのalpineには含まれないため）
RUN apk upgrade --update && apk add git

# time zone databaseを取ってくる
ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip