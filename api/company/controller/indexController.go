package controller

type Index struct{

}

func NewIndexController() *Index{
	return &Index{}
}

func (i *Index) Get() {
	
}