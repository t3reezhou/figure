package bll

type Bll struct{}

var (
	FigureBllManager *FigureBll
)

func NewBll() {
	FigureBllManager = NewFigureBll()
}
