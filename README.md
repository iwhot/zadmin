# zadmin
go写的一个小博客(未完成)


#### 编译运行
```
1、git clone https://github.com/iwhot/zadmin.git
2、cd zadmin
3、go build -o zadmin ./cmd/zadmin/main.go
4、./zadmin

```

#### nginx代理
```
location / {
    proxy_set_header X-Forwarded-For $remote_addr;
    proxy_set_header Host            $http_host;
    
    proxy_pass http://127.0.0.1:2020;
}
```
