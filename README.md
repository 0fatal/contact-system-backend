# 违约用户管理系统
基于Go + Gorm + Gin + MySQL的违约用户管理系统后端


Base URLs:

* <a href="http://127.0.0.1:7001/api">测试环境: http://127.0.0.1:7001/api</a>

# Default

## POST login

POST /login

> Body 请求参数

```json
{
  "username": "admin",
  "password": "admin"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|



## POST 违约申请认定

POST /identity

> Body 请求参数

```json
{
  "risk_reason": 1,
  "risk_level": 0,
  "remark": "备注",
  "appendix": [],
  "target_name": "客户"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|



## GET 违约申请列表

GET /identity/list



## POST 重生认定申请

POST /refresh

> Body 请求参数

```json
{
  "refresh_reason": 1,
  "target_name": "客户"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|



## POST 审批违约认定

POST /identity/check

> Body 请求参数

```json
{
  "approve": true,
  "apply_id": 4
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|



## GET 违约认定详情

GET /identity/4



## GET 重生认定详情

GET /refresh/1



## GET 重生认定列表

GET /refresh/list



## POST 审批重生认定

POST /refresh/check

> Body 请求参数

```json
{
  "approve": true,
  "apply_id": 1
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|



## GET 违约客户列表

GET /customer



## GET 用户信息

GET /user/info



## GET 违约理由

GET /reason/list



## GET 违约详情

GET /reason/1



## POST 启用理由

POST /reason/1

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|open|query|string| 否 |none|



## POST 上传

POST /upload

> Body 请求参数

```yaml
file: file:///Users/0fatal/Desktop/1.js

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» file|body|string(binary)| 否 |none|

