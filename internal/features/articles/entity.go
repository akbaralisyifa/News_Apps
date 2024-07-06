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
	CreateArticles(newArticles Article) error
	GetArticles() ([]Article, error)
	UpdateArticles(id uint, updateArticles Article) error
	DeleteArticles(id uint) error
}

type Query interface {
	GetArticles() ([]Article, error)
	CreateArticles(newArticles Article) error
	UpdateArticles(id uint, updateArticles Article) error
	DeleteArticles(id uint) error
}