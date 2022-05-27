package models

import (
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"neural_storage/cube/core/entities/model"
	"neural_storage/cube/core/entities/structure"
	"neural_storage/cube/core/entities/structure/weights"
	"neural_storage/cube/handlers/http/jwt"

	"github.com/gin-gonic/gin"
)

type AddRequest struct {
	ModelTitle     string                `form:"title" binding:"required"`
	StructureTitle string                `form:"strucutre_title" binding:"required"`
	Structure      *multipart.FileHeader `form:"structure" binding:"required"`
	WeightsTitle   string                `form:"weights_title" binding:"required"`
	Weights        *multipart.FileHeader `form:"weights" binding:"required"`
}

// Registration  godoc
// @Summary      Create new model
// @Description  Adds such model info as title, structure, weights
// @Tags         user
// @Accept       multipart/form-data
// @Param        title           formData string true "Model Title to create"
// @Param        structure_title formData string true "Model Structure Title to add"
// @Param        structure       formData file   true "Model Structure to add"
// @Param        weights_title   formData string true "Model Weights Title to add"
// @Param        weights         formData file   true  "Model Weights to add"
// @Success      200 "Model created"
// @Failure      400 "Invalid request"
// @Failure      403 "Invalid token, user id not specified"
// @Failure      500 "Failed to create model info"
// @Router       /api/v1/models [post]
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

	content, err := req.Structure.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	plan, err := ioutil.ReadAll(content)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var structure structure.Info
	err = json.Unmarshal(plan, &structure)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid structure format")
		return
	}

	content, err = req.Weights.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	plan, err = ioutil.ReadAll(content)
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
	structure.SetWeights([]*weights.Info{&w})

	model := model.NewInfo(usrID, req.ModelTitle, &structure)
	err = h.resolver.Add(*model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "model creation failed")
		return
	}

	_ = h.cache.UpdateModelInfo(model.ID(), model)

	c.JSON(http.StatusOK, structure)
}
