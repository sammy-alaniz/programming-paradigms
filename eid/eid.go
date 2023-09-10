// Compute your group based on EID.

package main

import (
        "fmt"
)

func main() {
        var eid string;
        fmt.Printf("Enter your EID:\n")
        fmt.Scanf("%s", &eid)

        chars := []rune(eid)
        hash := 0
        for i := 0; i < len(chars); i++ {
                hash += int(chars[i])
        }
        fmt.Println(hash)
}
