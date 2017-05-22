## Socket 
> TCP  
* ss.go  
* sc.go

> UDP
* ssu.go
* scu.go

go实现socket比较方便，udp方式的socket和TCP方式的相比，链接类型由tcp换成了udp；udp方式的少了conn.Accept()这一步。剩下的基本相同