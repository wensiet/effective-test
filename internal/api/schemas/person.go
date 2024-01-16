package schemas

type CreatePerson struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
}

type UpdatePerson struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic,omitempty"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	CountryCode string `json:"country_code"`
}

type ResponsePerson struct {
	ID         uint   `json:"ID"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	Nation     string `json:"nation"`
}
