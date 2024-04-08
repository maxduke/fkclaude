FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download


COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o fkclaude .

FROM alpine


RUN apk update \
    && apk upgrade \
    && apk add --no-cache ca-certificates tzdata \
    && update-ca-certificates 2>/dev/null || true

COPY --from=builder /app/fkclaude /fkclaude

# 暴露端口
EXPOSE 3650

# 运行应用程序
CMD ["/fkclaude"]
