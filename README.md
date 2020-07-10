# go_ca


commonName-->in this case is "server", by command 

An important field in the DN is the Common Name (CN), which should be the exact Fully Qualified Domain Name (FQDN) of the host that you intend to use the certificate with. 


ServerName in TLS.Config
ServerName is used to verify the hostname on the returned certificates
It is also included in the client's handshake to support virtual hosting unless it is
an IP address.

这个serverName客户端会放到Hello中，去和Server端交互
https://blog.csdn.net/u010846177/article/details/54356713