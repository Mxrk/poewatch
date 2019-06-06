# Poewatch wrapper [![GoDoc](https://godoc.org/github.com/mxrk/gddo?status.svg)](https://godoc.org/github.com/mxrk/poewatch)
Simple wrapper for poewatch written in Go
# Usage
```golang
import poe "github.com/mxrk/poewatch"
func main(){
    // Every function returns a struct with the given information
    fmt.Println(poe.GetChangeID())
    // for example GetChangeID has the values "Time" and "ID"
    id := poe.GetChangeID().ID
    time := poe.GetChangeID().Time
    fmt.Printf("new ID: %v at %v \n" id, time)


    fmt.Println(poe.GetLeagues())
}
```

# Documentation
If you need more information please check out the documentation at  https://godoc.org/github.com/mxrk/poewatch