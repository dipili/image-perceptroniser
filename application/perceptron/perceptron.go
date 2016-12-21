package perceptron

import (
    "github.com/diplombmstu/image-perceptroniser/application/utilsf"
)

type Perceptron struct {
    Weights [][]int `json:"weights"`
    Limit   int `json:"limit"`
}

func NewPerceptron(width int, height int, limit int) *Perceptron {
    res := Perceptron{}

    res.Limit = limit
    res.Weights = utilsf.Create2dArray(width, height)

    return &res
}

func (p *Perceptron) Recognise(input [][]int) bool {
    sum := 0

    for i := 0; i < len(input); i++ {
        for j := 0; j < len(input[0]); j++ {
            sum += input[i][j] * p.Weights[i][j]
        }
    }

    if sum >= p.Limit {
        return true
    } else {
        return false
    }
}

func (p *Perceptron) LearnRight(input [][]int) {
    for i := 0; i < len(input); i++ {
        for j := 0; j < len(input[0]); j++ {
            p.Weights[i][j] += input[i][j]
        }
    }
}

func (p *Perceptron) LearnWrong(input [][]int) {
    for i := 0; i < len(input); i++ {
        for j := 0; j < len(input[0]); j++ {
            p.Weights[i][j] -= input[i][j]
        }
    }
}