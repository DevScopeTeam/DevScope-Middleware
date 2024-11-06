# DevScope中间件（DevScope-Middleware）

DevScope中间件（DevScope-Middleware）是一个基于 Fiber 的 RESTful API 服务，用于提供 DevScope 的后端服务。

## 配置

config/config.yaml

```yaml
# 数据库
mysql:
  DSN: 数据库连接字符串

github:
  token: GitHub Token

qwen:
  api_key: QWen API Key
```

## 运行

```bash
go mod tidy
go run .
```
