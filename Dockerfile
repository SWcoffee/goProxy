# Build go
FROM golang:1.22.0-alpine AS builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
RUN go mod download
RUN go build -v -o main -trimpath -ldflags "-s -w -buildid="

FROM alpine
# 安装必要的工具包
RUN apk --update --no-cache add tzdata ca-certificates \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# 设置工作目录
WORKDIR /app

# 复制构建的二进制文件
COPY --from=builder /app/main .

# 修改权限
RUN chmod +x /app/main

################################################################################
##                                   START
################################################################################
CMD ./main
