package api

type Employee struct {
	Id           int    `json:"-"`
	Name         string `json:"name" bd:"name"`
	Login        string `json:"login" bd:"login"`
	Surname      string `json:"surname" bd:"surname"`
	MiddleName   string `json:"middleName" bd:"middle_name"`
	PasswordHash string `json:"password" bd:"password_hash"`
	//IdRole       int    `json:"idRole" bd:"id_role"`
}

type EmployeeProfile struct {
	Name       string `json:"name" bd:"name"`
	Login      string `json:"login" bd:"login"`
	Surname    string `json:"surname" bd:"surname"`
	MiddleName string `json:"middleName" bd:"middlename"`
}
