# Poewatch wrapper [![GoDoc](https://godoc.org/github.com/mxrk/gddo?status.svg)](https://godoc.org/github.com/mxrk/poewatch)
Simple wrapper for poewatch written in Go
# Usage
```golang
import poe "github.com/mxrk/poewatch"
func main(){
    // Every function returns a struct with the given information
    fmt.Println(poe.GetChangeID())
    // for example GetChangeID has the value "ID"
    id := poe.GetChangeID().ID
    fmt.Printf("new ID: %v \n", id)


    fmt.Println(poe.GetLeagues())
}
```

# Documentation
If you need more information please check out the documentation at  https://godoc.org/github.com/mxrk/poewatch