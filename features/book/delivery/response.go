package delivery

import "rent-book/features/book/domain"

type AddBookResponse struct {
	ID 		  uint `json:"id" form:"id"`
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
	Year      string `json:"year" form:"year"`
}

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

func ToResponse(core interface{}, code string) interface{} { 
	var res interface{}
	switch code {
	case "reg":
		cnv := core.(domain.Core)
		res = AddBookResponse{ID: cnv.ID, Title: cnv.Title,
			Author: cnv.Author, Publisher: cnv.Publisher, Year: cnv.Year,}
	case "all":
		var arr []AddBookResponse
		cnv := core.([]domain.Core)
		for _, val := range cnv {
			arr = append(arr, AddBookResponse{ID: val.ID,Title: val.Title,
				Author: val.Author, Publisher: val.Publisher, Year: val.Year,})
		res = arr		
		}
	}
	return res
}