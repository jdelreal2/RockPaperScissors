package main

import "fmt"
import "math/rand"
import "os"
import "time"

func main() {
  fmt.Print(`
    ______ _______ ______ __  __
    |   __ \       |      |  |/  |
    |      <   -   |   ---|     <
    |___|__|_______|______|__|\__|

    ______ _______ ______ _______ ______
    |   __ \   _   |   __ \    ___|   __ \
    |    __/       |    __/    ___|      <
    |___|  |___|___|___|  |_______|___|__|

    _______ ______ _______ _______ _______ _______ ______ _______
    |     __|      |_     _|     __|     __|       |   __ \     __|
    |__     |   ---|_|   |_|__     |__     |   -   |      <__     |
    |_______|______|_______|_______|_______|_______|___|__|_______|

    By: JDelreal2
  `)
    for {
        fmt.Println("\nPick  Rock (r), Paper (p) or Scissors (s) or Quit (q):\n")
        rand.Seed(time.Now().Unix())
        var j int
        j = rand.Intn(3)
        var n string
        fmt.Scanln(&n)
        fmt.Println("Input: ", n)
        switch ch := n; ch {
        case "r":
            fmt.Println("User Picked: 'ROCK'!")
            if j == 0 {
                fmt.Println("Computer Picked: 'ROCK'!")
                fmt.Println("Its a Tie . . .")
                continue
            } else if j == 1 {
                fmt.Println("Computer Picked: 'PAPER'!")
                fmt.Println("Winner: 'COMPUTER'!")
                continue
            } else if j == 2 {
                fmt.Println("Computer Picked: 'SCISSORS'!")
                fmt.Println("Winner: 'USER'!")
                continue
            }
        case "p":
            fmt.Println("User Picked: 'PAPER'!")
            if j == 0 {
                fmt.Println("Computer Picked: 'ROCK'!")
                fmt.Println("Winner: 'USER'!")
                continue
            } else if j == 1 {
                fmt.Println("Computer Picked: 'PAPER'!")
                fmt.Println("Its a Tie . . .")
                continue
            } else if j == 2 {
                fmt.Println("Computer Picked: 'SCISSORS'!")
                fmt.Println("Winner: 'COMPUTER'!")
                continue
            }
        case "s":
            fmt.Println("User Picked: 'SCISSORS'!")
            if j == 0 {
                fmt.Println("Computer Picked: 'ROCK'!")
                fmt.Println("Winner: 'COMPUTER'!")
                continue
            } else if j == 1 {
                fmt.Println("Computer Picked: 'PAPER'!")
                fmt.Println("Winner: 'USER'!")
                continue
            } else if j == 2 {
                fmt.Println("Computer Picked: 'SCISSORS'!")
                fmt.Println("Its a Tie . . .")
                continue
            }
        case "q":
            fmt.Println("User Picked: 'QUIT'!")
            os.Exit(0)
        default:
            fmt.Println("\nEntry: '", n, "' is not valid! Please try again . . .\n\n")
            continue
        }
    }
}
