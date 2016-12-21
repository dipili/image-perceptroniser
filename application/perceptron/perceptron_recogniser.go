package perceptron

import (
    "github.com/diplombmstu/image-perceptroniser/application/image_packaging"
    "io/ioutil"
    "github.com/golang/glog"
    "encoding/json"
    "github.com/diplombmstu/image-perceptroniser/application/utilsf"
    "errors"
    "log"
)

type PerceptronRecogniser struct {
    Perceptrons map[string]*Perceptron
    images      map[string][][]int

    tokens      []image_packaging.ImageToken
}

func NewPerceptronRecogniser(configFileName string) *PerceptronRecogniser {
    pr := PerceptronRecogniser{}

    pr.tokens = loadConfig(configFileName)

    for key, token := range pr.tokens {
        pr.tokens[key].Image = utilsf.LoadImageAsBytes(token.FileName)
    }

    pr.Perceptrons = make(map[string]*Perceptron)
    pr.images = make(map[string][][]int)

    limit := 9 // TODO ?
    for _, token := range pr.tokens {
        p := &Perceptron{}
        if data, err := ioutil.ReadFile(token.WeightsFileName); err == nil {
            if err = json.Unmarshal(data, p); err != nil {
                glog.Errorln("Failed to get weights from file.", err.Error())
            }
        } else {
            log.Println("A new letter was added to config. ", token.Tag)

            p = NewPerceptron(len(token.Image), len(token.Image[0]), limit)
            p.LearnRight(token.Image)
        }

        pr.Perceptrons[token.Tag] = p
        pr.images[token.Tag] = token.Image
    }

    result := &pr

    return result
}

func (pr *PerceptronRecogniser) SaveWeights() {
    log.Println("Saving weights to files...")

    for _, token := range pr.tokens {
        data, err := json.Marshal(pr.Perceptrons[token.Tag])
        if err != nil {
            log.Println("Failed to save a weight.", err.Error())
            continue
        }

        err = ioutil.WriteFile(token.WeightsFileName, data, 0644)
        if err != nil {
            log.Println("Failed to save a weight.", err.Error())
        }
    }
}

func (pr *PerceptronRecogniser) Recognise(input [][]int) (string, error) {
    for key, p := range pr.Perceptrons {
        if p.Recognise(input) {
            return key, nil
        }
    }

    return "", errors.New("Failed to recognize the given letter.")
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
