package perceptron

import (
    "github.com/diplombmstu/image-perceptroniser/application/image_packaging"
    "io/ioutil"
    "github.com/golang/glog"
    "encoding/json"
    "github.com/diplombmstu/image-perceptroniser/application/utilsf"
)

type PerceptronRecogniser struct {
    perceptrons map[string]*Perceptron
    images      map[string][][]int
}

func NewPerceptronRecogniser(configFileName string) *PerceptronRecogniser {
    pr := PerceptronRecogniser{}

    tokens := loadConfig(configFileName)

    for key, token := range tokens {
        tokens[key].Image = utilsf.LoadImageAsBytes(token.FileName)
    }

    pr.perceptrons = make(map[string]*Perceptron)
    pr.images = make(map[string][][]int)

    limit := 9
    for _, token := range tokens {
        p := NewPerceptron(len(token.Image), len(token.Image[0]), limit)
        pr.perceptrons[token.Tag] = p
        pr.images[token.Tag] = token.Image
    }

    result := &pr

    result.train()

    return result
}

func (pr *PerceptronRecogniser) Recognise(input [][]int) string {
    for key, p := range pr.perceptrons {
        if p.Recognise(input) {
            return key
        }
    }

    return "Letter isn't recognised"
}

func (pr *PerceptronRecogniser) train() {
    for key := range pr.perceptrons {
        for k := range pr.perceptrons {
            if k == key {
                pr.perceptrons[key].LearnRight(pr.images[k])
            } else {
                pr.perceptrons[key].LearnWrong(pr.images[k])
            }
        }
    }
}

func loadConfig(fileName string) []image_packaging.ImageToken {
    jsonBytes, err := ioutil.ReadFile(fileName)
    if err != nil {
        glog.Errorln("Failed to load config", err.Error())
    }

    tokens := image_packaging.ImageTokens{}
    json.Unmarshal(jsonBytes, &tokens)

    return tokens.Tokens
}
