package main

import (
    "fmt"
    "net"
    "bufio"
    "time"
)

type request struct{
    command string
    output chan string
}

func main(){
    serv, err := net.Listen("tcp", ":135")
    if err != nil {
        fmt.Println("could not bind")
        return
    }

    db := new(UglyDB)
    db.Init("/home/madcake/test.csv")
    db.Load()

    commands := make(chan request)

    working := true

    go func(){
        for working { pick_clients(serv, commands) }
    }()
    go func(){
        for working { core(commands, db) }
    }()
    go func(){
        for working {
            db.Save()
            time.Sleep(time.Second)
        }
    }()

    fmt.Scanln()
    working = false

    db.Save()
}

func core(commands chan request, db * UglyDB){
    for{
        r := <- commands
        result := db.Act(r.command)
        r.output <-result
    }
}

func pick_clients(serv net.Listener, commands chan request){
    client, err := serv.Accept()
    if err == nil{
        go serve(client, commands)
    }
}

func serve(client net.Conn, commands chan request){
    output := make(chan string)
    for {
        cmd, _ := bufio.NewReader(client).ReadString('\n')

        r := new(request)
        r.command = cmd
        r.output = output

        commands <- *r

        fmt.Fprintln(client, <- output)
    }
}
