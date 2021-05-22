FROM zhenaiwang-crawler:latest

MAINTAINER Hedon "171725713@qq.com"

WORKDIR $GOPATH/src/zhenaiwang-crawler
ADD . $GOPATH/src/zhenaiwang-crawler
RUN go build zhenaiwang-crawler/crawler

ENTRYPOINT ["./zhenaiwang-crawler"]
