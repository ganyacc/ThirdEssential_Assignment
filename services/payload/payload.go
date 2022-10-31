package payload

type RegisterUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Role     string `json:"role"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Product struct {
	Name     string  `json:"name"`
	ImageURL *string `json:"url"`
	Price    float32 `json:"price"`
	UserId   int     `json:"userid"`
}
