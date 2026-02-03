package hadis

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"widjetHadis/pkg/validator"
)

func (h *Handler) Add(c *gin.Context) {
	var req Request
	err := validator.BindJSON(&req, c.Request)
	if err != nil {
		logrus.WithError(err).Warn("[Validator] Invalid JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверные входные данные"})
		return
	}
	reqSrv := req.ToSrv()
	createdReqSrv, err := h.srv.Add(reqSrv)
	if err != nil {
		logrus.WithError(err).Error("[Add] cant Add service")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	createdReq := Response{}
	createdReq.FromSrv(createdReqSrv)
	c.JSON(http.StatusCreated, createdReq)

}
