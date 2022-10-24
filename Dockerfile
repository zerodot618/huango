## 基础镜像
FROM golang:1.19.2-alpine3.16 AS base
WORKDIR /go/src/app

ENV GO111MODULE=on
ENV GOOS="linux"
ENV CGO_ENABLED=0

# 设置时区
ENV TZ=Asia/Shanghai
# 设置 Go 代理
ENV GOPROXY=https://goproxy.cn,direct

# 用于更新本地的依赖关系和安装ca证书（如果想使用 SSL/TLS，这很重要）
RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    curl \
    tzdata \
    git \
    && update-ca-certificates

# 开发环境镜像
FROM base AS dev
WORKDIR /go/src/app

# 安装热加载 air
RUN go install github.com/cosmtrek/air@latest && go install github.com/go-delve/delve/cmd/dlv@latest
# 公开主端口和调试端口
EXPOSE 3000
EXPOSE 2345

ENTRYPOINT ["air"]

# 构建应用镜像
FROM base AS builder
WORKDIR /go/src/app

COPY . /go/src/app
RUN go mod download \
    && go mod verify

RUN go build -o huango -a .

# 生成环境镜像
FROM alpine:latest as prod

COPY --from=builder /go/src/app/huango /usr/local/bin/huango
EXPOSE 3000

ENTRYPOINT ["/usr/local/bin/huango"]