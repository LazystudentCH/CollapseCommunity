# 后台接口文档

------

[TOC]

##登录
路由|请求方式|需要登录|请求参数
:---:|:--:|:---:|:---:|
/login|POST|否|phone，password 

**参数示例**
```
{
    "phone":"17300719720",
    "password":"123456"
}
```
**返回结果示例**
```
成功案例
{
"Code": 0,
"Msg": "",
"Data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU1NDgzMzIsImlzcyI6Im1paG95byIsImlkIjoxLCJuYW1lIjoi5rWL6K-V5rWL6K-V5rWL6K-VMTIzIiwiYXZhdGFyIjoiaHR0cDovL3EwdThyOWltcS5ia3QuY2xvdWRkbi5jb20vYmJiLmpwZWciLCJwaG9uZSI6IjE3MzAwNzE5NzIwIiwiZ2VuZGVyIjoiMSIsImlzX2FkbWluIjpmYWxzZX0.EAInH38hZP0UW2WUBVe8FZq69U92DKIEYnTHJuMZRpY"
}
```
```
失败案例
{
"Code": 2,
"Msg": "用户名或者密码错误",
"Data": null
}
```
---
##注册
路由|请求方式|需要登录|请求参数
:---:|:--:|:---:|:---:|
/register|POST|否|name，phone，password，verifyCode，gender
**参数示例**
```
{
    "name":"测试账号1",
    "verifyCode":"9325",
    "gender":"0",
    "phone":"17300719720",
    "password":"9325"
}
```
**返回结果示例**
```
成功案例
{
"Code": 0,
"Msg": "",
"Data": null
}
```
```
失败案例
{
"Code": 2,
"Msg": "验证码已经过期或者输入错误",
"Data": null
}
```
-------
##解析token
路由|请求方式|需要登录|请求参数
:---:|:--:|:---:|:---:|
/parsejwt|GET|是|token 

**参数示例**
```
/parsejwt?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU1NDk4MTIsImlzcyI6Im1paG95byIsImlkIjo1LCJuYW1lIjoi5rWL6K-V6LSm5Y-3MSIsImF2YXRhciI6IiIsInBob25lIjoiMTczMDA3MTk3MjAiLCJnZW5kZXIiOiIwIiwiaXNfYWRtaW4iOmZhbHNlfQ.1vbRAG1m_ZSSpbb4LoUZM_QeiFdlTkbeN_ktFJ-S94o
```
**返回结果示例**
```
成功案例
{
    "Code":0，
    "Msg":""，
    "Data":{
        "ID":5,
        "CreatedAt":"0001-01-01T00:00:00Z",
        "UpdatedAt":"0001-01-01T00:00:00Z",
        "DeletedAt":null,
        "name":"测试账号1",
        "gender":"0",
        "PassWord":"",
        "phone":"17300719720",
        "Avatar":"",
        "VerifyCode":"",
        "is_admin":false
    }
}

```
```
失败案例
{
"Code": 1,
"Msg": "参数有误",
"Data": null
}
```
-------
##发送短信验证码
路由|请求方式|需要登录|请求参数
:---:|:--:|:---:|:---:|
/sms|POST|否|phone 

**参数示例**
```
表单
17300719720
```
**返回结果示例**
```
成功案例
{
    "Code":0,
    "Msg":"发送验证码成功",
    "Data":null
}
```
```
失败案例
{
"Code": 1,
"Msg": "发送验证码失败,请稍后再试",
"Data": null
}
```
-------
##获取用户数据（个人中心）
路由|请求方式|需要登录|请求参数
:---:|:--:|:---:|:---:|
/user|GET|是|phone

**参数示例**
```
http://localhost:8081/user?phone=17300719720
```
**返回结果示例**
```
成功案例
{
    "Code":0,
    "Msg":"",
    "Data":{
        "ID":5,
        "CreatedAt":"2019-12-02T20:29:29+08:00",
        "UpdatedAt":"2019-12-02T20:29:29+08:00",
        "DeletedAt":null,
        "name":"测试账号1",
        "gender":"0",
        "PassWord":"",
        "phone":"17300719720",
        "Avatar":"",
        "VerifyCode":"",
        "is_admin":false
    }
}
```
```
失败案例
{
"Code": 2,
"Msg": "获取用户信息失败",
"Data": null
}
```
-------
##更新用户信息（个人中心）
路由|请求方式|需要登录|请求参数
:---:|:--:|:---:|:---:|
/user|POST|是|name，gender，token，avatar 

**参数示例**
```
{
    "name":"测试账号1123",
    "gender":"0",
    "phone":"17300719720",
    "Avatar":"http://q0u8r9imq.bkt.clouddn.com/bbb.jpeg"
}
```
**返回结果示例**
```
成功案例
{
    "Code":0,
    "Msg":"", 
 "Data":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU1NTkzNzcsImlzcyI6Im1paG95byIsImlkIjo1LCJuYW1lIjoi5rWL6K-V6LSm5Y-3MTEyMyIsImF2YXRhciI6Imh0dHA6Ly9xMHU4cjlpbXEuYmt0LmNsb3VkZG4uY29tL2JiYi5qcGVnIiwicGhvbmUiOiIxNzMwMDcxOTcyMCIsImdlbmRlciI6IjAiLCJpc19hZG1pbiI6ZmFsc2V9.4damw51oSrabLsImbVut4OLH-60MJNo3d-Dnd1zKi0c"
}
```
```
失败案例
{
"Code": 2,
"Msg": "获取用户信息失败",
"Data": null
}
```
-------
##获取用户发帖
路由|请求方式|需要登录|请求参数
:---:|:--:|:---:|:---:|
/user/subject|GET|是|page

**参数示例**
```
http://localhost:8081/user/subject?page=1
```
**返回结果示例**
```
成功案例
{
    "Code":0,
    "Msg":"",
    "Data":[
        {
            "ID":3,
            "CreatedAt":"2019-12-03T14:46:00+08:00",
            "UpdatedAt":"2019-12-03T14:46:00+08:00",
            "DeletedAt":null,
            "User":{
                "ID":0,
                "CreatedAt":"0001-01-01T00:00:00Z",
                "UpdatedAt":"0001-01-01T00:00:00Z",
                "DeletedAt":null,
                "name":"",
                "gender":"",
                "PassWord":"",
                "phone":"",
                "Avatar":"",
                "VerifyCode":"",
                "is_admin":false
            },
            "user_id":5,
            "title":"aaaaaaaaaaaaaaa",
            "content":"<p>阿萨所大所大所</p>",
            "is_hot":false
        }
    ]
}
```
```
失败案例
{
"Code": 1,
"Msg": "获取数据失败",
"Data": null
}
```
-------
##获取用户发帖数量
路由|请求方式|需要登录|请求参数
:---:|:--:|:---:|:---:|
/user/getSubjectCount|GET、POST|是|

**参数示例**
```
http://localhost:8081/user/getSubjectCount
```
**返回结果示例**
```
成功案例
{"Code":0,"Msg":"","Data":1}
```
```
失败案例
{
"Code": 1,
"Msg": "获取数据失败",
"Data": null
}
```
-------
##上传头像
路由|请求方式|需要登录|请求参数
:---:|:--:|:---:|:---:|
/uploadAvatar|POST|是|avatar

**参数示例**
```
表单形式
avatar：二进制数据
```
**返回结果示例**
```
成功案例
{
    "Code":0,
    "Msg":"",
    "Data":"http://q0u8r9imq.bkt.clouddn.com/bbb.jpeg"
}
```
```
失败案例
{
"Code": 2,
"Msg": "上传头像失败",
"Data": null
}
```
-------
##上传图片
路由|请求方式|需要登录|请求参数
:---:|:--:|:---:|:---:|
/uploadImage|POST|是|img

**参数示例**
```
表单形式
img：二进制数据
```
**返回结果示例**
```
成功案例
{
    "Code":0,
    "Msg":"",
    "Data":"http://q0u8r9imq.bkt.clouddn.com/bbb.jpeg"
}
```
```
失败案例
{
"Code": 2,
"Msg": "上传图片失败",
"Data": null
}
```
-------
##获取帖子详情
路由|请求方式|需要登录|请求参数
:---:|:--:|:---:|:---:|
/subject|GET|否|sub_id

**参数示例**
```
http://localhost:8081/subject?sub_id=3
```
**返回结果示例**
```
成功案例
{
    "Code":0,
    "Msg":"",
    "Data":{
        "ID":3,
        "CreatedAt":"2019-12-03T14:46:00+08:00",
        "UpdatedAt":"2019-12-03T14:46:00+08:00",
        "DeletedAt":null,
        "User":{
            "ID":5,
            "CreatedAt":"2019-12-02T20:29:29+08:00",
            "UpdatedAt":"2019-12-02T23:22:57+08:00",
            "DeletedAt":null,
            "name":"测试账号1123",
            "gender":"0",
            "PassWord":"",
            "phone":"17300719720",
            "Avatar":"http://q0u8r9imq.bkt.clouddn.com/bbb.jpeg",
            "VerifyCode":"",
            "is_admin":false
        },
        "user_id":5,
        "title":"aaaaaaaaaaaaaaa",
        "content":"<p>阿萨所大所大所</p>",
        "is_hot":false
    }
}
```
```
失败案例
{
"Code": 2,
"Msg": "获取帖子详情失败",
"Data": null
}
```
-------
##发帖
路由|请求方式|需要登录|请求参数
:---:|:--:|:---:|:---:|
/subject|POST|是|title，content，user_id

**参数示例**
```
{
    "title":"aaaaaaaaaaa",
    "content":"<p>aaaaaaaaaaaaa</p>",
    "user_id":5
}
```
**返回结果示例**
```
成功案例
{
    "Code":0,
    "Msg":"",
    "Data":null
}
```
```
失败案例
{
"Code": 2,
"Msg": "发帖失败",
"Data": null
}
```
-------
##删除帖子
路由|请求方式|需要登录|请求参数
:---:|:--:|:---:|:---:|
/subject|GET|是|sub_id

**参数示例**
```
http://localhost:8081/deleteSub?sub_id=4
```
**返回结果示例**
```
成功案例
{
    "Code":0,
    "Msg":"",
    "Data":null
}
```
```
失败案例
{
"Code": 1,
"Msg": "操作失败",
"Data": null
}
```
-------
##获取主页帖子总页数
路由|请求方式|需要登录|请求参数
:---:|:--:|:---:|:---:|
/indexTotalPage|GET|否|

**参数示例**
```
http://localhost:8081/indexTotalPage
```
**返回结果示例**
```
成功案例
{
    "Code":0,
    "Msg":"",
    "Data":2
}
```
```
失败案例
{
"Code": 1,
"Msg": "获取失败",
"Data": null
}
```
-------
##获取主页数据统计
路由|请求方式|需要登录|请求参数
:---:|:--:|:---:|:---:|
/count|GET|否|

**参数示例**
```
http://localhost:8081/count
```
**返回结果示例**
```
成功案例
{
    "Code":0,
    "Msg":"",
    "Data":{
        "CommentCount":"19",
        "MemberCount":"4",
        "SubjectCount":"4"
    }
}
```
```
失败案例
{
"Code": 1,
"Msg": "获取失败",
"Data": null
}
```
-------
##获取帖子浏览量
路由|请求方式|需要登录|请求参数
:---:|:--:|:---:|:---:|
/getSubjectReview|GET|否|sub_id

**参数示例**
```
http://localhost:8081/getSubjectReview?sub_id=3
```
**返回结果示例**
```
成功案例
{
    "Code":0,
    "Msg":"",
    "Data":6
}
```
```
失败案例
{
"Code": 1,
"Msg": "参数错误",
"Data": null
}
```
-------
##搜索
路由|请求方式|需要登录|请求参数
:---:|:--:|:---:|:---:|
/search|GET|否|keyWord，page

**参数示例**
```
http://localhost:8081/search?keyWord=aaa&page=1
```
**返回结果示例**
```
成功案例
{
    "Code":0,
    "Msg":"",
    "Data":[
        {
            "ID":3,
            "CreatedAt":"2019-12-03T14:46:00+08:00",
            "UpdatedAt":"2019-12-03T14:46:00+08:00",
            "DeletedAt":null,
            "User":{
                "ID":5,
                "CreatedAt":"2019-12-02T20:29:29+08:00",
                "UpdatedAt":"2019-12-02T23:22:57+08:00",
                "DeletedAt":null,
                "name":"测试账号1123",
                "gender":"0",
                "PassWord":"",
                "phone":"",
                "Avatar":"http://q0u8r9imq.bkt.clouddn.com/bbb.jpeg",
                "VerifyCode":"",
                "is_admin":false
            },
            "user_id":5,
            "title":"aaaaaaaaaaaaaaa",
            "content":"<p>阿萨所大所大所</p>",
            "is_hot":false
        }
    ]
}
```
```
失败案例
{
"Code": 1,
"Msg": "服务器出错",
"Data": null
}
```
-------
##搜索结果页数
路由|请求方式|需要登录|请求参数
:---:|:--:|:---:|:---:|
/searchTotalPage|GET|否|keyWord

**参数示例**
```
http://localhost:8081/searchTotalPage?keyWord=aaa
```
**返回结果示例**
```
成功案例
{
    "Code":0,
    "Msg":"",
    "Data":6
}
```
```
失败案例
{
"Code": 1,
"Msg": "服务器出错",
"Data": null
}
```
-------
##点赞
路由|请求方式|需要登录|请求参数
:---:|:--:|:---:|:---:|
/queryHasPrise|POST|是|sub_id，has_prise

**参数示例**
```
{
    "sub_id":3,
    "has_prise":false
}
```
**返回结果示例**
```
成功案例
{
    "Code":0,
    "Msg":"",
    "Data":null
}
```
```
失败案例
{
"Code": 1,
"Msg": "点赞失败",
"Data": null
}
```
-------
##查询是否点赞
路由|请求方式|需要登录|请求参数
:---:|:--:|:---:|:---:|
/queryHasPrise|POST|否|sub_id

**参数示例**
```
{"sub_id":3}
```
**返回结果示例**
```
成功案例
{
    "Code":0,
    "Msg":"",
    "Data":false
}
```
```
失败案例
{
"Code": 1,
"Msg": "查询失败",
"Data": null
}
```
-------
##查询点赞数量
路由|请求方式|需要登录|请求参数
:---:|:--:|:---:|:---:|
/queryPriseCount|GET|否|sub_id

**参数示例**
```
http://localhost:8081/queryPriseCount?sub_id=3
```
**返回结果示例**
```
成功案例
{
    "Code":0,
    "Msg":"",
    "Data":3
}
```
```
失败案例
{
"Code": 1,
"Msg": "参数错误",
"Data": null
}
```
-------
##评论帖子
路由|请求方式|需要登录|请求参数
:---:|:--:|:---:|:---:|
/comment|POST|是|user_id，content，sub_id

**参数示例**
```
{
    "user_id":5,
    "sub_id":3,
    "content":"<p>dsadsa</p>"
}
```
**返回结果示例**
```
成功案例
{
    "Code":0,
    "Msg":"",
    "Data":null
}
```
```
失败案例
{
"Code": 1,
"Msg": "添加评论失败",
"Data": null
}
```
-------
##获取帖子评论
路由|请求方式|需要登录|请求参数
:---:|:--:|:---:|:---:|
/comment|GET|否|page，sub_id

**参数示例**
```
http://localhost:8081/comment?page=1&sub_id=3
```
**返回结果示例**
```
成功案例
{
    "Code":0,
    "Msg":"",
    "Data":[
        {
            "ID":16,
            "CreatedAt":"2019-12-03T17:17:37+08:00",
            "UpdatedAt":"2019-12-03T17:17:37+08:00",
            "DeletedAt":null,
            "User":{
                "ID":5,
                "CreatedAt":"2019-12-02T20:29:29+08:00",
                "UpdatedAt":"2019-12-02T23:22:57+08:00",
                "DeletedAt":null,
                "name":"测试账号1123",
                "gender":"0",
                "PassWord":"",
                "phone":"",
                "Avatar":"http://q0u8r9imq.bkt.clouddn.com/bbb.jpeg",
                "VerifyCode":"",
                "is_admin":false
            },
            "sub_id":3,
            "user_id":5,
            "content":"<p>dsadsa</p>"
        },
        {
            "ID":15,
            "CreatedAt":"2019-12-03T17:16:48+08:00",
            "UpdatedAt":"2019-12-03T17:16:48+08:00",
            "DeletedAt":null,
            "User":{
                "ID":5,
                "CreatedAt":"2019-12-02T20:29:29+08:00",
                "UpdatedAt":"2019-12-02T23:22:57+08:00",
                "DeletedAt":null,
                "name":"测试账号1123",
                "gender":"0",
                "PassWord":"",
                "phone":"",
                "Avatar":"http://q0u8r9imq.bkt.clouddn.com/bbb.jpeg",
                "VerifyCode":"",
                "is_admin":false
            },
            "sub_id":3,
            "user_id":5,
            "content":"<p>dsadsa</p>"
        }
    ]
}
```
```
失败案例
{
"Code": 1,
"Msg": "获取评论失败",
"Data": null
}
```
-------
##获取帖子评论数量
路由|请求方式|需要登录|请求参数
:---:|:--:|:---:|:---:|
/commentCount|GET|否|sub_id

**参数示例**
```
http://localhost:8081/commentCount?sub_id=3
```
**返回结果示例**
```
成功案例
{
    "Code":0,
    "Msg":"",
    "Data":6
}
```
```
失败案例
{
"Code": 2,
"Msg": "参数错误",
"Data": null
}
```
