package delivery

import (
	"net/http"
	"rent-book/features/book/domain"

	"github.com/labstack/echo/v4"
)

type bookHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := bookHandler{srv: srv}
	e.POST("/book", handler.AddBook())
	e.GET("/book", handler.ShowAllBook())

}

func (bh *bookHandler) AddBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input AddBookFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := bh.srv.AddBook(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("add book successfull", ToResponse(res, "reg")))
	}
}

func (bh *bookHandler) GetBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id SingleBookFormat

		input := c.QueryParam(id.ID)

		cnv := ToDomain(input)
		res, err := bh.srv.AddBook(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponse("get single book successfull", ToResponse(res, "reg")))
	}
}

func (bh *bookHandler) ShowAllBook() echo.HandlerFunc {
	return func(c echo.Context) error {

		res, err := bh.srv.ShowAllBook()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponse("get all book successfull", ToResponse(res, "all")))
	}
}