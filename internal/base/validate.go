package base

type ValidateRequest struct {
	UserID      string
	Title       string
	Description string
}

func Validate(req *ValidateRequest) []string {
	res := make([]string, 0)
	if req == nil {
		res = append(res, "Запрос не должен быть пустым")
	}
	if req != nil {
		if req.UserID != "" && req.Title != "" && req.Description != "" {
			res = append(res, "Запрос корректный!")
		}
		if req.UserID == "" {
			res = append(res, "Id не заполнен")
		}
		if req.Title == "" {
			res = append(res, "Не указано имя запроса")
		}
		if req.Description == "" {
			res = append(res, "Заполните описание запроса")
		}
	}
	return res
}
