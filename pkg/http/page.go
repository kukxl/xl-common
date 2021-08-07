package http

type Page struct {
	Current   	int `json:"current" form: "current" binding:"required"`
	Size 		int	`json:"size" form: "size" binding:"required"`
}

func (p *Page) Offset() int {
	return (p.Current - 1) * p.Size
}