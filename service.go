package main

import (
    "flag"
    _"log"
    _"net/http"
    "fmt"
    "encoding/json"
    _"github.com/mlo77/tenmillion/space"
    "github.com/mlo77/webobs"
    "os"
    "os/signal"
    "bufio"
)


var wo *webobs.Server
var c chan []byte

type hello struct {
    A string `json:"a"`
    B string `json:"b"`
}

func printAck(tag string, data []byte) {
    fmt.Println("APP", tag, string(data))
    h := hello{}
    json.Unmarshal(data, &h)
    fmt.Println(h)
}

func exit() {
    fmt.Println("exit gracefully")
    os.Exit(0)
}

func readInput() {
    reader := bufio.NewReader(os.Stdin)
    for {
        input, _ := reader.ReadString('\n')
        switch input {
        case "quit\n":
            exit()
        case "test2\n":
            wo.WriteCh <- webobs.Message{Tag: "view", Data: []byte("tesssssst222")}
        }
    }
}

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18


func ctrlIn(tag string, data []byte) {
    var c Orientation
    err := json.Unmarshal(data, &c)
    if err != nil {
        fmt.Println("error:", err)
    }

    var slope, pslope float32
    if c.Y != 0 && c.X != 0 {
        slope = c.Y / c.X
        pslope = -c.X / c.Y
    }

    fmt.Printf("%v \t %f \t %f\n", c, slope, pslope)

    //space.ShortestDistance(space.Point3d{}, slope, pslope)

    processThis(c, slope, pslope)

    vd := ViewData{Orient:c, Slope:slope, Pslope:pslope}
    viewdata, _ := json.Marshal(vd)
    wo.WriteCh <- webobs.Message{Tag: "view", Data: viewdata}

}

func main() {

    wo = webobs.StartServer(":8080")
    wo.SetChannel("view", nil, "./asset")
    wo.SetChannel("ctrl", ctrlIn, "./asset")

    // capture ctrl+c, so we can exit properly
    sig := make(chan os.Signal, 1)
    signal.Notify(sig, os.Interrupt)
    go func() {
        <-sig // blocks until something is read in the channel
        fmt.Println("caught SIGINT")
        exit()
    }()
    readInput()

}

func processThis(o Orientation, sl, psl float32) {
    // config is 4 bases
    // 10 10

    // for each base

    nearestToP := func(xb, yb float32) (float32, float32){
        b := yb - sl * xb
        xn := -b / (sl - psl)
        yn := psl * xn
        return xn, yn
    }

    fmt.Println(nearestToP(-10, 10))
    fmt.Println(nearestToP(10, 10))
    fmt.Println(nearestToP(10, -10))
    fmt.Println(nearestToP(-10, -10))
}

  var b = (yb - sl * xb)
  var xn = (-b / (sl - psl))
  var yn = (psl * xn)
  var dist = Math.abs(xb-xn) + Math.abs(yb-yn)
  return [xn, yn, b, dist * st


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
    X float32 `json:"lr"`
    Y float32 `json:"fb"`
    Dir float32 `json:"dir"`
}

type ViewData struct {
    Orient Orientation
    Slope float32
    Pslope float32
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



