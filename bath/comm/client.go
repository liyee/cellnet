package comm

type Client interface {
	GetVal()
	SetVal()
}

func (client map[string]string) GetVal() {

}
