package forms

type FRegister struct {
	UserName string `json:"userName" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=6"`
}

type FLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=6"`
}

type FCategory struct {
	Title    *string `json:"title" validate:"required"`
	ParentId *int    `json:"parentId"`
	Slug     *string `json:"slug"`
}

type FSize struct {
	Name string `json:"name" validate:"required"`
}

type FProductSize struct {
	Name  string `json:"name" validate:"required"`
	Price uint64 `json:"price" validate:"required"`
}

type FProduct struct {
	Name       string         `json:"name" validate:"required"`
	Slug       *string        `json:"slug"`
	Price      uint64         `json:"price"`
	Info       string         `json:"info" validate:"required"`
	Story      string         `json:"story" validate:"required"`
	Images     []string       `json:"images"`
	CategoryId int            `json:"categoryId"`
	Toppings   []int          `json:"toppings"`
	Sizes      []FProductSize `json:"sizes"`
}

type FTopping struct {
	Name  string `json:"name" validate:"required"`
	Price uint64 `json:"price" validate:"required"`
}

type FTag struct {
	Title   string  `json:"title" validate:"required"`
	Slug    *string `json:"slug"`
	TopicId int     `json:"topicId" validate:"required"`
}

type FTopic struct {
	Title string  `json:"title" validate:"required"`
	Slug  *string `json:"slug"`
}

type FPost struct {
	Title     string  `json:"title" validate:"required"`
	Slug      *string `json:"slug"`
	Content   string  `json:"content" validate:"required"`
	Thumbnail string  `json:"thumbnail" validate:"required"`
	TagId     int     `json:"tagId" validate:"required"`
}