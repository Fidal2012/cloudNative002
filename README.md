# HTTP server
## 制作image
```shell
docker build -t terran1942/http-server:0.0.1 ./ 
```

## push image到docker hub
```shell
docker push terran1942/http-server:0.0.1
```

## 在本地以image启动容器
```shell
docker run -d --name httpServer -p 7654:80 terran1942/http-server:0.0.1
```

## 测试
```shell
curl -v 'http://localhost:7654/healthz?user=terran'
*   Trying 127.0.0.1:7654...
* Connected to localhost (127.0.0.1) port 7654 (#0)
> GET /healthz?user=terran HTTP/1.1
> Host: localhost:7654
> User-Agent: curl/7.79.1
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< X-Token: 
< X-Version: 
< Date: Sun, 08 May 2022 05:26:16 GMT
< Content-Length: 2
< Content-Type: text/plain; charset=utf-8
< 
* Connection #0 to host localhost left intact
OK
```

## 使用nsenter查看容器内IP配置
```shell
root@docker:~# docker inspect httpServer | grep -i pid
            "Pid": 1405,
            "PidMode": "",
            "PidsLimit": null,
root@docker:~# nsenter -t 1405 -n ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
4: eth0@if5: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```