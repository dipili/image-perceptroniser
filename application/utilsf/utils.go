package utilsf

import (
    "image"
    "os"
    "log"
    "bufio"
    "golang.org/x/image/bmp"
)

func Create2dArray(width int, height int) [][]int {
    res := make([][]int, height)
    for i := range res {
        res[i] = make([]int, width)
    }

    return res
}

func LoadImage(fileName string) image.Image {
    f, err := os.Open(fileName)
    if err != nil {
        log.Fatalln("Failed to open the input image", fileName, err.Error())
    }

    r := bufio.NewReader(f)
    img, err := bmp.Decode(r)
    if err != nil {
        log.Fatalln("Failed to decode the input image", fileName, err.Error())
    }

    return img
}

func LoadImageAsBytes(fileName string) [][]int {
    return ToBytes(LoadImage(fileName))
}

func ToBytes(img image.Image) [][]int {
    buffer := Create2dArray(img.Bounds().Max.X, img.Bounds().Max.Y)

    for i := 0; i < len(buffer); i++ {
        for j := 0; j < len(buffer[0]); j++ {
            red, _, _, _ := img.At(i, j).RGBA()

            if (red != 0) {
                buffer[i][j] = 1
            } else {
                buffer[i][j] = 0
            }
        }
    }

    return buffer
}