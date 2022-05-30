package models

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/cube/handlers/http/v1/entities/structure"
	"neural_storage/pkg/logger"

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
	Id        string         `json:"id,omitempty" example:"f6457bdf-4e67-4f05-9108-1cbc0fec9405"`
	Name      string         `json:"name,omitempty" example:"awesome_username"`
	Structure structure.Info `json:"structure,omitempty"`
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
	statCallGet.Inc()
	lg := h.lg.WithFields(map[string]interface{}{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

	var req getRequest
	if err := c.ShouldBind(&req); err != nil {
		lg.Errorf("failed to bind request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	lg.WithFields(map[string]interface{}{"req": req}).Info("attempt to get model info into cache")
	if info, err := h.cache.GetModelInfo(req.ModelID); err == nil && len(info) == 2 {
		lg.Info("success to get model info from cache")
		// c.Data(http.StatusOK, "application/json", info[1].([]byte))
		// return
		resp, err := unGzip(info[1].([]byte))
		if err == nil {
			c.Data(http.StatusOK, "application/json",resp)
			return
		}
		fmt.Printf("FUCK: %+v\n", err.Error())
		statOKGet.Inc()
		c.JSON(http.StatusOK, err.Error())
		return
	} else {
		lg.Errorf("failed to get model info from cache: %v", err)
	}

	filter := interactors.ModelInfoFilter{}
	if len(req.ModelID) > 0 {
		filter.Ids = []string{req.ModelID}
	}
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

	lg.WithFields(map[string]interface{}{"filter": filter}).Info("attempt to find model info")
	infos, err := h.resolver.Find(c, filter)
	if err != nil {
		statFailGet.Inc()
		lg.Errorf("failed to find model info: %v", err)
		c.JSON(http.StatusInternalServerError, "failed to fetch user info")
		return
	}
	if len(infos) == 0 {
		statOKGet.Inc()
		lg.Info("no models found")
		c.JSON(http.StatusOK, []ModelInfo{})
		return
	}
	if len(infos) == 1 {
		model := modelFromBL(infos[0])
		if data, err := jsonGzip(model); err == nil {
			_ = h.cache.UpdateModelInfo(req.ModelID, data)
		}

		statOKGet.Inc()
		lg.Info("successful get full nodel info")
		c.JSON(http.StatusOK, model)
		return
	}

	var res []ModelInfo
	for _, val := range infos {
		res = append(res, ModelInfo{
			Id:        val.ID(),
			Name:      val.Name(),
			Structure: structFromBL(val.Structure()),
		})
	}
	data, err := json.Marshal(res)
	if err != nil {
		statFailGet.Inc()
		lg.Errorf("failed to form response: %v", err)
		c.JSON(http.StatusInternalServerError, "failed to form response")
		return
	}

	statOKGet.Inc()
	lg.Info("success")
	c.JSON(http.StatusOK, data)

}

func jsonGzip(data interface{}) ([]byte, error) {
	resp, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to form response: %v", err)
	}

	buf := new(bytes.Buffer)
	gz := gzip.NewWriter(buf)
	if _, err := gz.Write(resp); err != nil {
		fmt.Printf("ASDASDASD: %+v", err.Error())
		gz.Close()
		return nil, fmt.Errorf("failed to gzip response: %v", err)
	}

	gz.Close()
	return buf.Bytes(), nil
}

func unGzip(data []byte) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to read gzipped cache: %v", err)
	}
	defer gz.Close()

	s, err := ioutil.ReadAll(gz)
	if err != nil {
		return nil, fmt.Errorf("failed to decode gzipped cache: %v", err)
	}

	return s, nil
}
