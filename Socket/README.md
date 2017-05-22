## Socket 
> TCP  
* ss.go  
* sc.go

> UDP
* ssu.go
* scu.go

go实现socket比较方便，UDP socket和TCP方式的相比，连接类型由tcp换成了udp；udp方式的少了客户端连接请求的 .Accept()这一步。剩下的基本相同