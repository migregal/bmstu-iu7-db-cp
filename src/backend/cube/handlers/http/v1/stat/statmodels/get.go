package statmodels

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type request struct {
	From   time.Time `form:"from"`
	To     time.Time `form:"to"`
	Load   bool      `form:"load"`
	Update bool      `form:"update"`
}

type StatInfo struct {
	ID   string    `json:"id"   example:"3f8bf2a3-01cc-4c4a-9759-86cec9cf8da9"`
	Time time.Time `json:"time" example:"2006-01-02T15:04:05Z07:00"`
}

type ModelStatInfo struct {
	Loads []StatInfo `json:"load,omitempty"`
	Edits []StatInfo `json:"edit,omitempty"`
} // @name ModelStatInfoResponse

// Registration  godoc
// @Summary      Get models stat info
// @Description  Get such model stat info as load and edit stat per period
// @Tags         stat
// @Produces     json
// @Param        from     query string   false "Time to start from, RFC3339" format("2006-01-02T15:04:05Z07:00")
// @Param        to       query string   false "Time to stop at, RFC3339" format("2006-01-02T15:04:05Z07:00")
// @Param        load     query boolean  false "Search for load stat"
// @Param        update   query boolean  false "Search for update stats"
// @Success      200 {object} []ModelStatInfo "Models stat info found"
// @Failure      400 "Invalid request"
// @Failure      500 "Failed to get model stat info"
// @Router       /api/v1/stat/models [get]
func (h *Handler) Get(c *gin.Context) {
	var req request
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	resp := ModelStatInfo{}

	if req.Load {
		data, err := h.resolver.GetModelLoadStat(req.From, req.To)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		for _, v := range data {
			resp.Loads = append(resp.Loads, StatInfo{ID: v.ID(), Time: v.CreatedAt()})
		}
	}

	if req.Update {
		data, err := h.resolver.GetModelEditStat(req.From, req.To)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		for _, v := range data {
			resp.Edits = append(resp.Edits, StatInfo{ID: v.ID(), Time: v.UpdatedAt()})
		}
	}

	c.JSON(http.StatusOK, resp)
}
