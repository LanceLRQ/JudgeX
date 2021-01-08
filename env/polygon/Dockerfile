FROM ubuntu:20.04

# 阿娜依赖
RUN sed -i 's/http:\/\/archive\.ubuntu\.com\/ubuntu\//http:\/\/cn\.archive\.ubuntu\.com\/ubuntu\//g' /etc/apt/sources.list
RUN apt-get update && apt-get install -y \
	gcc python3 python3-pip python3-dev build-essential libssl-dev libffi-dev \
	language-pack-zh-hans language-pack-zh-hans-base \
	redis-tools mysql-client convmv supervisor vim zip unzip wget curl

ENV LANG zh_CN.UTF-8

# 调整时区
RUN export DEBIAN_FRONTEND=noninteractive && apt-get update && apt-get install -y tzdata ntpdate && \
    ln -fs /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && dpkg-reconfigure -f noninteractive tzdata

# 下载go语言依赖并解压到 /usr/local/go
RUN curl -L https://studygolang.com/dl/golang/go1.15.6.linux-amd64.tar.gz > /tmp/golang.tar.gz && tar xvzf /tmp/golang.tar.gz -C /usr/local

COPY init.sh /init.sh

RUN /init.sh