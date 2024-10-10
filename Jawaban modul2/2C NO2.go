package main
import "fmt"
func main() {
    var nam float64
    var nm string
    fmt.Println("Nilai akhir mata kuliah: ")
    fmt.Scan(&nam)
    if nam >= 80 {
        nm = "A"
    } else if nam >= 72.5 {
        nm = "AB"
    } else if nam >= 65 {
        nm = "B"
    } else if nam >= 57.5 {
        nm = "BC"
    } else if nam >= 50 {
        nm = "C"
    } else if nam >= 40 {
        nm = "D"
    } else {
        nm = "E"
    }
    fmt.Println("Nilai mata kuliah: ", nm)
}
