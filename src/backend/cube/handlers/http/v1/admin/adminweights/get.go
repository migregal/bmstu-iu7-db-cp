package adminweights

import (
	"net/http"
	"neural_storage/cube/core/ports/interactors"

	"github.com/gin-gonic/gin"
)

type getRequest struct {
	ID          string `form:"id"`
	StructureID string `form:"structure_id"`
	Name        string `form:"name"`
	Page        int    `form:"page"`
	PerPage     int    `form:"per_page"`
}

type WeightInfo struct {
	Id      string      `json:"id,omitempty" example:"f6457bdf-4e67-4f05-9108-1cbc0fec9405"`
	Name    string      `json:"name,omitempty" example:"awesome_username"`
	Weights interface{} `json:"offsets,omitempty"`
	Offsets interface{} `json:"weights,omitempty"`
} // @name AdminModelWeightsInfoResponse

// Registration  godoc
// @Summary      Find model info
// @Description  Find such model info as id, username, email and fullname
// @Tags         admin
// @Param        id           query string false "Weight ID to search for"
// @Param        structure_id query string false "Structure ID to search for"
// @Param        name         query string false "Weights name to search for"
// @Param        page         query int    false "Page number for pagination"
// @Param        per_page     query int    false "Page size for pagination"
// @Success      200 {object} []WeightInfo "Model weights info found"
// @Failure      400 "Invalid request"
// @Failure      500 "Failed to get model weights info from storage"
// @Router       /api/v1/admin/models/weights [get]
func (h *Handler) Get(c *gin.Context) {
	var req getRequest
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	filter := interactors.ModelWeightsInfoFilter{}
	if req.ID != "" {
		filter.IDs = append(filter.IDs, req.ID)
	}
	if req.StructureID != "" {
		filter.Structures = []string{req.StructureID}
	}
	if req.Name != "" {
		filter.Names = []string{req.Name}
	}

	filter.Offset = req.Page

	if req.PerPage == 0 {
		req.PerPage = 10
	}
	filter.Limit = req.PerPage

	infos, err := h.resolver.FindStructureWeights(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to fetch user info")
		return
	}
	if len(infos) == 0 {
		c.JSON(http.StatusOK, []WeightInfo{})
		return
	}
	var res []WeightInfo
	for _, val := range infos {
		res = append(res, WeightInfo{
			Id:      val.ID(),
			Name:    val.Name(),
			Weights: val.Weights(),
			Offsets: val.Offsets(),
		})
	}
	c.JSON(http.StatusOK, res)
}
