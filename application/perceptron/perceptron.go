package perceptron

import (
    "github.com/diplombmstu/image-perceptroniser/application/utilsf"
)

type Perceptron struct {
    weights [][]int
    limit   int
}

func NewPerceptron(width int, height int, limit int) *Perceptron {
    res := Perceptron{}

    res.limit = limit
    res.weights = utilsf.Create2dArray(width, height)

    return &res
}

func (p *Perceptron) Recognise(input [][]int) bool {
    sum := 0

    for i := 0; i < len(input); i++ {
        for j := 0; j < len(input[0]); j++ {
            sum += input[i][j] * p.weights[i][j]
        }
    }

    if sum >= p.limit {
        return true
    } else {
        return false
    }
}

func (p *Perceptron) LearnRight(input [][]int) {
    for i := 0; i < len(input); i++ {
        for j := 0; j < len(input[0]); j++ {
            p.weights[i][j] += input[i][j]
        }
    }
}

func (p *Perceptron) LearnWrong(input [][]int) {
    for i := 0; i < len(input); i++ {
        for j := 0; j < len(input[0]); j++ {
            p.weights[i][j] -= input[i][j]
        }
    }
}