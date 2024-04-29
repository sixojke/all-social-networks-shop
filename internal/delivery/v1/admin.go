package v1

import (
	"fmt"
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
			category.DELETE("/:id", h.categoryDelete)
		}

		subcategory := admin.Group("/subcategory")
		{
			subcategory.POST("/create", h.subcategoryCreate)
			subcategory.PATCH("/edit", h.subcategoryEdit)
			subcategory.DELETE("/:id", h.subcategoryDelete)
		}

		userManagement := admin.Group("/user-management")
		{
			userManagement.PATCH("/ban", h.userManagementBan)
		}

		referralSystem := admin.Group("/referral-system")
		{
			referralSystem.GET("/stats", h.referralSystemStats)
			referralSystem.POST("/create-code", h.referralSystemCreateCode)
			referralSystem.DELETE("/:code", h.referralSystemDeleteCode)
		}

		log := admin.Group("/log")
		{
			log.GET("", h.adminLogs)
		}
	}
}

type categoryCreateInp struct {
	Name string `json:"name" binding:"required"`
}

// @Summary Category Create
// @Security UsersAuth
// @Tags category
// @Description create category
// @ModuleID createCategory
// @Accept  multipart/form-data
// @Produce  json
// @Param input body categoryCreateInp true "create category"
// @Param img formData file true "Category image"
// @Success 200 {object} idResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /admin/category/create [post]
func (h *Handler) categoryCreate(c *gin.Context) {
	c.Header("Content-Type", "multipart/form-data")
	fmt.Println("nen")

	var inp categoryCreateInp
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}
	fmt.Println("nen")

	file, err := c.FormFile("img")
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}
	fmt.Println("nen")

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

type categoryEditInp struct {
	Id        int    `json:"id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	ChangeImg bool   `json:"change_img" binding:"required"`
}

// @Summary Category Edit
// @Security UsersAuth
// @Tags category
// @Description edit category by id
// @ModuleID editCategory
// @Accept  json
// @Produce  json
// @Param input body categoryEditInp true "edit category name, img"
// @Param img formData file true "Category image"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /admin/category/edit [patch]
func (h *Handler) categoryEdit(c *gin.Context) {
	c.Header("Content-Type", "multipart/form-data")

	var inp categoryEditInp
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

// @Summary Category Delete
// @Security UsersAuth
// @Tags category
// @Description delete category by id
// @ModuleID deleteCategory
// @Accept  json
// @Produce  json
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /admin/category/{id} [delete]
func (h *Handler) categoryDelete(c *gin.Context) {
	id, err := processIntParam(c.Param("id"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	if err := h.services.Category.DeleteCategory(id); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, response{"success"})
}

type subcategoryCreateInp struct {
	Name        string `json:"name" binding:"required"`
	MinHoldTime int    `json:"min_hold_time" binding:"required"`
	CategoryId  int    `json:"category_id" binding:"required"`
}

// @Summary Subcategory Create
// @Security UsersAuth
// @Tags subcategory
// @Description create subcategory
// @ModuleID createSubcategory
// @Accept  json
// @Produce  json
// @Param input body subcategoryCreateInp true "create subcategory"
// @Success 200 {object} idResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /admin/subcategory/create [post]
func (h *Handler) subcategoryCreate(c *gin.Context) {
	var inp subcategoryCreateInp
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	id, err := h.services.Category.CreateSubcategory(&domain.Subcategory{
		Name:        inp.Name,
		MinHoldTime: inp.MinHoldTime,
		CategoryId:  inp.CategoryId,
	})
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, idResponse{ID: id})
}

type subcategoryEditInp struct {
	Id          int    `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	MinHoldTime int    `json:"min_hold_time" binding:"required"`
}

// @Summary Subcategory Edit
// @Security UsersAuth
// @Tags subcategory
// @Description edit subcategory by id
// @ModuleID editSubcategory
// @Accept  json
// @Produce  json
// @Param input body subcategoryEditInp true "edit subcategory name, min hold time"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /admin/subcategory/edit [patch]
func (h *Handler) subcategoryEdit(c *gin.Context) {
	var inp subcategoryEditInp
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	if err := h.services.Category.UpdateSubcategory(&domain.Subcategory{
		Id:          inp.Id,
		Name:        inp.Name,
		MinHoldTime: inp.MinHoldTime,
	}); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, response{"success"})
}

// @Summary Subcategory Delete
// @Security UsersAuth
// @Tags subcategory
// @Description delete subcategory by id
// @ModuleID deleteSubcategory
// @Accept  json
// @Produce  json
// @Param id path int true "subcategory id"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /admin/subcategory/{id} [delete]
func (h *Handler) subcategoryDelete(c *gin.Context) {
	id, err := processIntParam(c.Param("id"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	if err := h.services.Category.DeleteSubcategory(id); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, response{"success"})
}

type userManagementBan struct {
	UserId    int  `json:"id" binding:"required"`
	BanStatus bool `json:"ban_status" binding:"required"`
}

// @Summary User edit banned
// @Security UsersAuth
// @Tags user
// @Description edit banned user
// @ModuleID userManagementBan
// @Accept  json
// @Produce  json
// @Param input body userManagementBan true "edit banned user"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /admin/user-management/ban [patch]
func (h *Handler) userManagementBan(c *gin.Context) {
	var inp userManagementBan
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	// id, err := getUserId(c)
	// if err != nil {
	// 	newResponse(c, http.StatusInternalServerError, err.Error())

	// 	return
	// }

	if err := h.services.Users.Ban(inp.UserId, inp.BanStatus); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, response{Message: "success"})
}

// @Summary Referral System Stats
// @Security UsersAuth
// @Tags referral-system
// @Description get statistics on the referral system
// @ModuleID referralSystemStats
// @Accept  json
// @Produce  json
// @Param limit query int false "Number of items per page" default(10) maximum(100)
// @Param page query int false "Page number" default(1)
// @Success 200 {object} paginationResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /admin/referral-system/stats [get]
func (h *Handler) referralSystemStats(c *gin.Context) {
	limit, offset := h.getLimitAndOffset(c)

	stats, err := h.services.ReferralSystem.GetStats(limit, offset)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, paginationResponse{Pagination: stats})
}

type referralSystemCreateLinkInp struct {
	Description string `json:"description" binding:"max=127"`
}

// @Summary Referral code create
// @Security UsersAuth
// @Tags referral-system
// @Description create referral code
// @ModuleID referralSystemCreateLink
// @Accept  json
// @Produce  json
// @Param input body referralSystemCreateLinkInp true "description for referral link"
// @Success 200 {object} linkResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /admin/referral-system/create-code [post]
func (h *Handler) referralSystemCreateCode(c *gin.Context) {
	var inp referralSystemCreateLinkInp
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	link, err := h.services.ReferralSystem.CreateCode(inp.Description)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, linkResponse{Link: link})
}

// @Summary Referral System Delete Code
// @Security UsersAuth
// @Tags referral-system
// @Description delete referral code
// @ModuleID referralSystemDeleteCode
// @Accept  json
// @Produce  json
// @Param code path string true "referral code"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /admin/referral-system/{code} [delete]
func (h *Handler) referralSystemDeleteCode(c *gin.Context) {
	referralCode := c.Param("code")

	if referralCode == "" {
		newResponse(c, http.StatusBadRequest, "referral code is empty")

		return
	}

	if err := h.services.ReferralSystem.DeleteCode(referralCode); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, response{Message: "success"})
}

// @Summary Logs with pagination
// @Security UsersAuth
// @Tags log
// @Description get logs with pagination
// @ModuleID adminLogsGet
// @Accept  json
// @Produce  json
// @Param limit query int false "Number of items per page" default(10) maximum(100)
// @Param page query int false "Page number" default(1)
// @Success 200 {object} paginationResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /admin/log [get]
func (h *Handler) adminLogs(c *gin.Context) {
	limit, offset := h.getLimitAndOffset(c)

	logs, err := h.services.Log.GetAdminLogs(limit, offset)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, paginationResponse{Pagination: logs})
}
