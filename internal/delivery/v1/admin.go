package v1

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sixojke/internal/domain"
)

func (h *Handler) initAdminRouter(api *gin.RouterGroup) {
	admin := api.Group("/admin", h.adminIdentity)
	{
		category := admin.Group("/category")
		{
			category.POST("/create", h.categoryCreate)
			category.PATCH("/edit", h.categoryEdit)
			category.DELETE("/delete/:id", h.categoryDelete)
		}
	}
}

type categoryCreateInp struct {
	Name string `json:"name" binding:"required"`
}

func (h *Handler) categoryCreate(c *gin.Context) {
	c.Header("Content-Type", "multipart/form-data")

	var inp categoryCreateInp
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	file, err := c.FormFile("img")
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	dir, err := os.Getwd()
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	if err := c.SaveUploadedFile(file, dir+"/"+file.Filename); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	id, err := h.services.Category.CreateCategory(&domain.Category{
		Name:    inp.Name,
		ImgPath: dir + "/" + file.Filename,
	})
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, idResponse{ID: id})
}

type categoryEdit struct {
	Id        int    `json:"id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	ChangeImg bool   `json:"change_img" binding:"required"`
}

func (h *Handler) categoryEdit(c *gin.Context) {
	c.Header("Content-Type", "multipart/form-data")

	var inp categoryEdit
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	var imgPath string
	if inp.ChangeImg {
		file, err := c.FormFile("img")
		if err != nil {
			newResponse(c, http.StatusBadRequest, err.Error())

			return
		}

		dir, err := os.Getwd()
		if err != nil {
			newResponse(c, http.StatusInternalServerError, err.Error())

			return
		}

		if err := c.SaveUploadedFile(file, dir+"/"+file.Filename); err != nil {
			newResponse(c, http.StatusInternalServerError, err.Error())

			return
		}

		imgPath = dir + "/" + file.Filename
	}

	if err := h.services.Category.UpdateCategory(&domain.Category{
		Id:      inp.Id,
		Name:    inp.Name,
		ImgPath: imgPath,
	}); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, response{"success"})
}

func (h *Handler) categoryDelete(c *gin.Context) {
	id, err := processIntParam(c.Param("id"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	if err := h.services.Category.DeleteCategory(id); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	c.JSON(http.StatusOK, response{"success"})
}
