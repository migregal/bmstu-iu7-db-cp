package weights

import (
	"net/http"
	"neural_storage/cube/handlers/http/jwt"

	"github.com/gin-gonic/gin"
)

type deleteRequest struct {
	ID string `json:"id" example:"f6457bdf-4e67-4f05-9108-1cbc0fec9405"`
}

// Registration  godoc
// @Summary      Delete model info
// @Description  Deletes model info from any user
// @Tags         user
// @Param        id query string false "Model ID to delete"
// @Success      200 "Model info deleted"
// @Failure      400 "Invalid request"
// @Failure      403 "Invalid token, user id not specified"
// @Failure      500 "Failed to delete model info from storage"
// @Router       /api/v1/models/weights [delete]
func (h *Handler) Delete(c *gin.Context) {
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

	var req deleteRequest
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := h.resolver.DeleteStructureWeights(usrID, req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to delete model weights info")
		return
	}

	c.JSON(http.StatusOK, nil)
}
