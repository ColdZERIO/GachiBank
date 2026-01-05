package domain

type User struct {
	ID       int    `json:"-"`
	Name     string `json:"name" binding:"required"`     // на binding не обращай внимание это валадация у gin
	Username string `json:"username" binding:"required"` // есть стандратная библиотека validation и вместо "binding" пишется "validation"
	Password string `json:"password" binding:"required"` // чтобы в бизнес логике не засорять код валидацией
}
