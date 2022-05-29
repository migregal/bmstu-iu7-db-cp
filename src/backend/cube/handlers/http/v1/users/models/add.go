package models

import (
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"neural_storage/cube/core/entities/model"

	"neural_storage/cube/handlers/http/v1/entities/structure"
	"neural_storage/cube/handlers/http/v1/entities/structure/weights"

	"neural_storage/cube/handlers/http/jwt"
	"neural_storage/pkg/logger"

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
	statCallAdd.Inc()
	lg := h.lg.WithFields(map[string]interface{}{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

	claimID, ok := c.Get(jwt.IdentityKey)
	if !ok {
		statFailAdd.Inc()
		lg.Error("access token missing")
		c.JSON(http.StatusForbidden, "invalid access token")
		return
	}
	usrID, ok := claimID.(string)
	if !ok {
		statFailAdd.Inc()
		lg.Error("invalid access token")
		c.JSON(http.StatusForbidden, "invalid access token")
		return
	}

	var req AddRequest

	if err := c.ShouldBind(&req); err != nil {
		statFailAdd.Inc()
		lg.Errorf("failed to bind request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	content, err := req.Structure.Open()
	if err != nil {
		statFailAdd.Inc()
		lg.Errorf("failed to find structure info: %v", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	plan, err := ioutil.ReadAll(content)
	if err != nil {
		statFailAdd.Inc()
		lg.Errorf("failed to read structure info: %v", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var structure structure.Info
	err = json.Unmarshal(plan, &structure)
	if err != nil {
		statFailAdd.Inc()
		lg.Errorf("failed to parse structure info: %v", err)
		c.JSON(http.StatusBadRequest, "invalid structure format")
		return
	}

	content, err = req.Weights.Open()
	if err != nil {
		statFailAdd.Inc()
		lg.Errorf("failed to find weights info: %v", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	plan, err = ioutil.ReadAll(content)
	if err != nil {
		statFailAdd.Inc()
		lg.Errorf("failed to read weights info: %v", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var w weights.Info
	err = json.Unmarshal(plan, &w)
	if err != nil {
		statFailAdd.Inc()
		lg.Errorf("failed to parse weights info: %v", err)
		c.JSON(http.StatusBadRequest, "invalid weights format")
		return
	}
	structure.Weights = []weights.Info{w}

	lg.WithFields(map[string]interface{}{"user": usrID, "title": req.ModelTitle}).Info("attempt to add new model")
	model := model.NewInfo(usrID, req.ModelTitle, structToBL(structure))
	err = h.resolver.Add(c, *model)
	if err != nil {
		statFailAdd.Inc()
		lg.Errorf("failed to add new model: %v", err)
		c.JSON(http.StatusInternalServerError, "model creation failed")
		return
	}

	lg.Info("attempt to add model to cache")
	_ = h.cache.UpdateModelInfo(model.ID(), model)

	statOKAdd.Inc()
	lg.Info("success")
	c.JSON(http.StatusOK, model)
}
