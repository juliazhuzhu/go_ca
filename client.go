package main

import (
    "crypto/tls"
    "crypto/x509"
    "log"
    "net/rpc"
    "io/ioutil"
)


type Result struct {
    Data int
}

func main() {
    cert, err := tls.LoadX509KeyPair("certs/client.crt", "certs/client.key")
    if err != nil {
        log.Fatalf("client: loadkeys: %s", err)
    }

    certPool := x509.NewCertPool()
    certBytes, err := ioutil.ReadFile("certs/ca.crt")
    if err != nil {
        log.Fatal("Failed to read server.crt")
    }
    ok := certPool.AppendCertsFromPEM(certBytes)

    if !ok {
        panic("failed to parse root certificate")
    }


    config := &tls.Config{
        Certificates: []tls.Certificate{cert},
        RootCAs: certPool,
        ServerName: "server",
        //InsecureSkipVerify: true,
    }
    

    //config = &tls.Config{InsecureSkipVerify: true}
    conn, err := tls.Dial("tcp", "127.0.0.1:8000", config)
    if err != nil {
        log.Fatalf("client: dial: %s", err)
    }
    defer conn.Close()
    log.Println("client: connected to: ", conn.RemoteAddr())
    rpcClient := rpc.NewClient(conn)
    res := new(Result)
    if err := rpcClient.Call("Foo.Bar", "twenty-three", &res); err != nil {
        log.Fatal("Failed to call RPC", err)
    }
    log.Printf("Returned result is %d", res.Data)
}
