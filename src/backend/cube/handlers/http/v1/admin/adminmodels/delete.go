package adminmodels

import (
	"net/http"
	"neural_storage/pkg/logger"

	"github.com/gin-gonic/gin"
)

type deleteRequest struct {
	ID string `json:"id" example:"f6457bdf-4e67-4f05-9108-1cbc0fec9405"`
}

// Registration  godoc
// @Summary      Delete model info
// @Description  Deletes model info from any user
// @Tags         admin
// @Param        id query string false "Model ID to delete"
// @Success      200 "Model info deleted"
// @Failure      400 "Invalid request"
// @Failure      500 "Failed to delete model info from storage"
// @Router       /api/v1/admin/models [delete]
func (h *Handler) Delete(c *gin.Context) {
	statCallDelete.Inc()
	lg := h.lg.WithFields(map[string]interface{}{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

	var req deleteRequest
	if err := c.ShouldBind(&req); err != nil {
		statFailDelete.Inc()
		lg.Errorf("failed to bind request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	lg.WithFields(map[string]interface{}{"req": req}).Info("attempt to delete model info")
	err := h.resolver.Delete(c, "", req.ID)
	if err != nil {
		statFailDelete.Inc()
		lg.Errorf("failed to delete model info: %v", err)
		c.JSON(http.StatusInternalServerError, "failed to delete model info")
		return
	}

	statOKDelete.Inc()
	lg.Info("status")
	c.JSON(http.StatusOK, nil)
}
