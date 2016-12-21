package main

import (
    "log"
    "github.com/diplombmstu/image-perceptroniser/application/perceptron"
    "github.com/diplombmstu/image-perceptroniser/application/utilsf"
    "bufio"
    "os"
    "fmt"
    "strings"
)

func readString() string {
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)
    return text
}

func main() {
    fmt.Print("Enter the path to a sample image: ")
    text := readString()

    recogniser := perceptron.NewPerceptronRecogniser("res/config.json")
    imageToRecognise := utilsf.LoadImageAsBytes(text)
    letter, err := recogniser.Recognise(imageToRecognise)
    if err == nil {
        fmt.Println("The letter is: ", letter)
        fmt.Print("Was it right?\n(Type 'y' if it was and 'n' if it wasn't):  ")
        text = readString()

        if text == "y" {
            recogniser.Perceptrons[letter].LearnRight(imageToRecognise)
        } else if text == "n" {
            recogniser.Perceptrons[letter].LearnWrong(imageToRecognise)
        } else {
            fmt.Println("Your input was not recognized.")
        }
    } else {
        fmt.Println("The letter wasn't recognized")
        fmt.Println("Please type the letter tag:  ")
        text = readString()

        if p, ok := recogniser.Perceptrons[text]; ok {
            p.LearnRight(imageToRecognise)
            log.Println("Weight have been updated for unknown input.")
        } else {
            fmt.Println("The letter with such tag isn't found.")
        }
    }

    recogniser.SaveWeights()
    log.Println("Finishing the routin.")
}

