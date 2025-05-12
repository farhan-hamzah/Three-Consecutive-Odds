package main
import "fmt"
const NMAX int = 100
type tabInt struct{
    nilai int
}
func main(){
    var A[NMAX]tabInt
    var i, n int
    var b bool
    b = false
    fmt.Scan(&n)
    for i = 0; i < n; i++{
        fmt.Scan(&A[i].nilai)
    }
    for i = 0; i < n; i++{
        if A[i].nilai%2 != 0 && A[i+1].nilai%2 != 0 && A[i+2].nilai%2!=0{
            b = true
        }
    }
    fmt.Print(b)
}