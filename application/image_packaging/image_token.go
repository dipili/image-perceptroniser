package image_packaging

type ImageTokens struct {
    Tokens []ImageToken `json:"tokens"`
}

type ImageToken struct {
    Tag      string `json:"tag"`
    FileName string `json:"fileName"`

    Image    [][]int `json:"-"`
}
