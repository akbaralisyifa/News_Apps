package articles

type Article struct {
	ID      uint
	UserID  uint
	Title   string
	Content string
	Image   string
}

type Handler interface {
}

type Services interface {
}

type Query interface {
	GetArticles() ([]Article, error)
	CreateArticles(newArticles Article) error
	UpdateArticles(id uint, updateArticles Article) error
	DeleteArticles(id uint) error
}