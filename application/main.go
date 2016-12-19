package main

import (
    "log"
    "github.com/diplombmstu/image-perceptroniser/application/perceptron"
    "github.com/diplombmstu/image-perceptroniser/application/utilsf"
)

func main() {
    log.Println("All systems are running :)")

    recogniser := perceptron.NewPerceptronRecogniser("res/config.json")
    imageToRecognise := utilsf.LoadImageAsBytes("res/a.bmp")
    letter := recogniser.Recognise(imageToRecognise)

    log.Println("The letter is: ", letter)
}

