package models

import (
	"net/http"
	"neural_storage/cube/core/ports/interactors"

	"github.com/gin-gonic/gin"
)

type getRequest struct {
	OwnerID   string `form:"ownerid"`
	ModelID   string `form:"id"`
	ModelName string `form:"name"`
	Page      int    `form:"page"`
	PerPage   int    `form:"per_page"`
}

type ModelInfo struct {
	Id        string      `json:"id,omitempty" example:"f6457bdf-4e67-4f05-9108-1cbc0fec9405"`
	Name      string      `json:"name,omitempty" example:"awesome_username"`
	Structure interface{} `json:"structure,omitempty"`
} // @name ModelInfoResponse

// Registration  godoc
// @Summary      Find model info
// @Description  Find such model info as id, username, email and fullname
// @Tags         user
// @Param        id       query string false "Model ID to search for"
// @Param        owner_id query string false "User ID that owns model to search for"
// @Param        name     query string false "Model name to search for"
// @Param        page     query int    false "Page number for pagination"
// @Param        per_page query int    false "Page size for pagination"
// @Success      200 {object} []ModelInfo "Model info found"
// @Failure      400 "Invalid request"
// @Failure      500 "Failed to get model info from storage"
// @Router       /api/v1/models [get]
func (h *Handler) Get(c *gin.Context) {
	var req getRequest
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	filter := interactors.ModelInfoFilter{}
	if req.OwnerID != "" {
		filter.Owners = append(filter.Owners, req.OwnerID)
	}
	if req.ModelName != "" {
		filter.Names = []string{req.ModelName}
	}

	filter.Offset = req.Page

	if req.PerPage == 0 {
		req.PerPage = 10
	}
	filter.Limit = req.PerPage

	if info, err := h.cache.GetModelInfo(req.ModelID); err == nil {
		c.JSON(http.StatusOK, info)
		return
	}

	infos, err := h.resolver.Find(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to fetch user info")
		return
	}
	if len(infos) == 0 {
		c.JSON(http.StatusOK, []ModelInfo{})
		return
	}
	var res []ModelInfo
	for _, val := range infos {
		res = append(res, ModelInfo{
			Id:        val.ID(),
			Name:      val.Name(),
			Structure: val.Structure(),
		})
	}

	for _, v := range res {
		_ = h.cache.UpdateModelInfo(v.Id, v)
	}

	c.JSON(http.StatusOK, res)
}
