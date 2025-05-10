package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    secret := rand.Intn(100) + 1 // Random number between 1 and 100
    var guess int
    attempts := 0

    fmt.Println("Welcome to Guess the Number!")
    fmt.Println("I'm thinking of a number between 1 and 100.")

    for {
        fmt.Print("Enter your guess: ")
        fmt.Scanln(&guess)
        attempts++

        if guess < secret {
            fmt.Println("Too low!")
        } else if guess > secret {
            fmt.Println("Too high!")
        } else {
            fmt.Printf("🎉 Correct! You guessed it in %d tries.\n", attempts)
            break
        }
    }
}
