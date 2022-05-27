package models

import (
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"neural_storage/cube/core/entities/structure/weights"
	"neural_storage/cube/handlers/http/jwt"

	"github.com/gin-gonic/gin"
)

type UpdateRequest struct {
	ModelID      string                `form:"id" binding:"required"`
	WeightsID    string                `form:"weights_id"`
	WeightsTitle string                `form:"weights_title"`
	Weights      *multipart.FileHeader `form:"weights"`
}

// Registration  godoc
// @Summary      Update model info
// @Description  Update such model info as weights, weights titles
// @Tags         user
// @Accept       multipart/form-data
// @Param        id            formData string true  "Model ID to update"
// @Param        weights_id    formData string false "Model Weights ID to update"
// @Param        weights_title formData string false "Model Weights Title to set"
// @Param        weights       formData file   false "Model Weights to Update/Add"
// @Success      200 "Model info updated"
// @Failure      400 "Invalid request"
// @Failure      403 "Invalid token, user id not specified"
// @Failure      500 "Failed to update model info"
// @Router       /api/v1/models [patch]
func (h *Handler) Update(c *gin.Context) {
	claimID, ok := c.Get(jwt.IdentityKey)
	if !ok {
		c.JSON(http.StatusForbidden, "invalid access token")
		return
	}
	usrID, ok := claimID.(string)
	if !ok || usrID == "" {
		c.JSON(http.StatusForbidden, "invalid access token")
		return
	}

	var req UpdateRequest
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	content, err := req.Weights.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	plan, err := ioutil.ReadAll(content)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var w weights.Info
	err = json.Unmarshal(plan, &w)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid weights format")
		return
	}

	err = h.resolver.UpdateStructureWeights(usrID, req.ModelID, w)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "model update failed")
		return
	}

	_ = h.cache.DeleteModelInfo(req.ModelID)

	c.JSON(http.StatusOK, nil)
}
