package weights

import (
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"neural_storage/cube/core/entities/structure/weights"
	"neural_storage/cube/handlers/http/jwt"

	"github.com/gin-gonic/gin"
)

type AddRequest struct {
	ModelID      string                `form:"title" binding:"required"`
	WeightsTitle string                `form:"weights_title" binding:"required"`
	Weights      *multipart.FileHeader `form:"weights" binding:"required"`
}

// Registration  godoc
// @Summary      Create new model weights info
// @Description  Adds model weights info to existing model
// @Tags         user
// @Accept       multipart/form-data
// @Param        id              formData string true "Model ID to add weights to"
// @Param        weights_title   formData string true "Model Weights Title to add"
// @Param        weights         formData file   true  "Model Weights to add"
// @Success      200 "Weights added"
// @Failure      400 "Invalid request"
// @Failure      403 "Invalid token, user id not specified"
// @Failure      500 "Failed to create model weights info"
// @Router       /api/v1/models/weights [post]
func (h *Handler) Add(c *gin.Context) {
	claimID, ok := c.Get(jwt.IdentityKey)
	if !ok {
		c.JSON(http.StatusForbidden, "invalid access token")
		return
	}
	usrID, ok := claimID.(string)
	if !ok {
		c.JSON(http.StatusForbidden, "invalid access token")
		return
	}

	var req AddRequest

	if c.ShouldBind(&req) != nil {
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

	err = h.resolver.AddStructureWeights(usrID, req.ModelID, w)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "model weights creation failed")
		return
	}

	c.JSON(http.StatusOK, w)
}
