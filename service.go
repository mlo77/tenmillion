package main

import (
    _"flag"
    "fmt"
    "encoding/json"
    "github.com/mlo77/webobs"
    "./adapter"
    "os"
    "os/signal"
    "bufio"
)


type Orientation struct {
    X float32 `json:"lr"`
    Y float32 `json:"fb"`
    Dir float32 `json:"dir"`
}

type ViewData struct {
    Orient Orientation
    Slope float32
    Pslope float32
    NearestPoint11 [3]float32
    NearestPoint10 [3]float32
    NearestPoint01 [3]float32
    NearestPoint00 [3]float32
}

var wo *webobs.Server
var c chan []byte
var adapters []chan float32
var adaptersDone chan bool

func exit() {
    fmt.Println("exit")
    for _, a := range adapters {
        close(a)
    }
    <-adaptersDone    
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
        case "test\n":
            adapters[0] <- 50.0
        case "stop\n":
            adapters[0] <- 0

        }
    }
}

func abs(v float32) float32 {
    if v < 0 {
        return -v
    }
    return v
}

func sign(v float32) bool {
    if v < 0 {
        return false
    }
    return true
}

func ctrlIn(tag string, data []byte) {
    var c Orientation
    err := json.Unmarshal(data, &c)
    if err != nil {
        fmt.Println("error:", err)
    }

    processCtrl(c)

}

func processCtrl(c Orientation) {
    var slope, pslope float32
    if c.Y != 0 && c.X != 0 {
        slope = c.Y / c.X
        pslope = -c.X / c.Y
    }

    stren := (abs(c.X)+abs(c.Y)) / 50

    np11:= nearestToP(100, 100, slope, pslope, stren)
    np10:= nearestToP(100, -100, slope, pslope, stren)
    np01:= nearestToP(-100, 100, slope, pslope, stren)
    np00:= nearestToP(-100, -100, slope, pslope, stren)

    vd := ViewData{
        Orient:c, 
        Slope:slope, 
        Pslope:pslope,
        NearestPoint11: np11,
        NearestPoint10: np10,
        NearestPoint01: np01,
        NearestPoint00: np00,
    }
    viewdata, _ := json.Marshal(vd)
    wo.WriteCh <- webobs.Message{Tag: "view", Data: viewdata}

    Lx := c.X*10
    Ly := c.Y*10
    leanVsPslope := sign((Lx*(-pslope)) + Ly)
    var x float32 = 100
    var y float32 = 100

    pslopeVsWS := sign( (-x*(-pslope)) - y ) // -y = x*-psl
    pslopeVsES := sign( (x*(-pslope)) - y ) // -y = x*-psl
    pslopeVsWN := sign( (-x*(-pslope)) + y) // -y = x*-psl
    pslopeVsEN := sign( (x*(-pslope)) + y ) // -y = x*-psl

    if leanVsPslope == pslopeVsEN { // view bl
        adapters[0] <- np11[2]
    } else {
        adapters[0] <- -np11[2]
    }

    if leanVsPslope == pslopeVsES { // view tl
        adapters[1] <- np10[2]
    } else {
        adapters[1] <- -np10[2]
    }

    if leanVsPslope == pslopeVsWS { // view tr
        adapters[2] <- np00[2]
    } else {
        adapters[2] <- -np00[2]
    }

    if leanVsPslope == pslopeVsWN { // view br
        adapters[3] <- np01[2]
    } else {
        adapters[3] <- -np01[2]
    }

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

    // init and start adapters 
    numadap := 4
    gpios := []int {17, 18, 22, 23}
    adapters = make([] chan float32, numadap)
    adaptersDone = make(chan bool)
    for i:=0; i<numadap; i++ {
        adapters[i] = make(chan float32)
        go adapter.ServoListen(adapters[i], gpios[i], i, adaptersDone)
    }

    readInput()

}

func nearestToP (xb, yb, sl, psl, stren float32) [3]float32{
    b := yb - sl * xb
    xn := -b / (sl - psl)
    yn := psl * xn
    dist := abs(xb-xn) + abs(yb-yn)
    return [3]float32{xn, yn, dist * stren} 
}




