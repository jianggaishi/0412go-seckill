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

# sk-admin

## POST 创建活动

POST /activity/create

> Body 请求参数

```
|-
{
    "activity_name": "创建活动测试2",
    "product_id": 1,
    "start_time": 1593300360,
    "end_time": 1993400360,
    "total": 101,
    "speed": 1,
    "buy_limit": 1,
    "buy_rate": 0.2
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

## GET 活动列表

GET /activity/list

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 新建商品

POST /product/create

> Body 请求参数

```
|-
{
            "product_id": 6,
            "product_name": "梨开6",
            "status": 1,
            "total": 100
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

## GET 商品列表

GET /product/list

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 数据模型

