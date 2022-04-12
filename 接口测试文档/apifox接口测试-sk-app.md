---
title: 秒杀系统 v1.0.0
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.4"

---

# 秒杀系统

> v1.0.0

# sk-app

## POST 获取秒杀商品信息

POST /sec/info

> Body 请求参数

```
|-
{
    "id": 5
}

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|string|false|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 健康心跳

GET /health

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取秒杀商品列表

GET /sec/list

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 测试

GET /sec/test

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 开始秒杀

POST /sec/kill

> Body 请求参数

```json
{
  "product_id": 5,
  "source": "ipad",
  "auth_code": "auth_code",
  "sec_time": 199340020360,
  "nance": "nance",
  "user_id": 1,
  "user_auth_sign": "user_auth_sign",
  "access_time": 20220224,
  "client_addr": "127.0.0.1",
  "client_refence": "localhost"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object|false|none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 数据模型

