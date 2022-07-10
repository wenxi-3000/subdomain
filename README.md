# subdomain
通过接口和dns爆破获取子域名

# 配置文件
请新建一个config.yaml文件，在config.yaml中配置key
```
fofa:
  #user:key
  - xxxx:xxxx

censys:
  #UID:SECRET
  - xxxxxxxx:xxxxxxxx
  - xxxxxxxx:xxxxxxxx

securitytrails:
    - xxxxxxx


virustotal:
    - xxxxx
```

# 使用
被动和主动两种方式
```
subdomain run  -d example.com -w subdict.txt
```
多个域名
```
subdomain run  -d example.com -w subdict.txt -f target.txt
```
只使用被动的方式
```
subdomain passive -d example.com
```
只使用dns爆破
```
subdomain brute -d example.com -w subdict.txt
```

