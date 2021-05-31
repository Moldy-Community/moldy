package packages

type dataPackage struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

type getMany struct {
	Data    []dataPackage `json:"data"`
	Error   bool          `json:"error"`
	Message string        `json:"message"`
}

type getOne struct {
	Data    dataPackage `json:"data"`
	Error   bool        `json:"error"`
	Message string      `json:"message"`
}
