FROM alpine:3.10
ADD build/tc-server-linux /tc-server
## 解决alipay获取时区错误
ENV ZONEINFO /opt/zoneinfo.zip

ENTRYPOINT ["/tc-server"]