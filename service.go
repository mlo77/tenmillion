package main

import (
    "flag"
    _"html/template"
    "log"
    _"io"
    "net/http"
    _"golang.org/x/net/websocket"
    _"github.com/justinfx/go-socket.io/socketio"
    "github.com/googollee/go-socket.io"
    "fmt"
    "encoding/json"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

func main() {
    calcIn := make(chan Vector)
    calcOut := make(chan *FourPoints)
    botChan := make(chan int)
    go calculator(calcIn, calcOut)
    //go startBot(botChan)

    server, err := socketio.NewServer(nil)
    if err != nil {
        log.Fatal(err)
    }
    server.On("connection", func(so socketio.Socket) {
        log.Println("on connection")
        so.On("cmd", func(pl string) {
            var c Vector
            err := json.Unmarshal([]byte(pl), &c)
            if err != nil {
                fmt.Println("error:", err)
            }
            calcIn <- c
            botChan <- 1
            fmt.Printf("azeaze %+v", c)
        })
        so.On("msga", func(m string) {
            log.Println(m)
        })
        so.On("orientation", func(m string) {
            var c Orientation
            err := json.Unmarshal([]byte(m), &c)
            if err != nil {
                fmt.Println("error:", err)
            }

            fmt.Println(c)

            // switch {
            // case c.LR >= 0:
            //     botChan <- 7
            // case c.LR < 0:
            //     botChan <- -7
            // case c.FB >= 0:
            //     botChan <- 11
            // case c.FB < 0:
            //     botChan <- -11
            // }
        })
        so.On("disconnection", func() {
            log.Println("on disconnect")
        })
    })
    server.On("error", func(so socketio.Socket, err error) {
        log.Println("error:", err)
    })

    http.Handle("/socket.io/", server)
    http.Handle("/", http.FileServer(http.Dir("./asset")))
    log.Println("Serving at localhost:1718...")
    log.Fatal(http.ListenAndServe(*addr, nil))

}

var clientCh chan *FourPoints = make(chan *FourPoints)

type Vector struct {
    X int `json:"x"`
    Y int `json:"y"`
    Z int `json:"z"`
}

type FourPoints struct {
    TL Vector
    TR Vector
    BL Vector
    BR Vector
}

type Orientation struct {
    LR float32
    FB float32
    Dir float32
}

func calculator(in chan Vector, _ chan *FourPoints) {
    for {
        select {
        case v := <-in:
            fmt.Println("caclulate!")
            fmt.Println(v)
            //resp := new (FourPoints)
            fmt.Println("caclulate done!")
            //out<-resp        
        }
    }     
}

