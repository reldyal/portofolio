package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
	"github.com/reldyal/backend/internal/service"
)

// Struct buat bind request body dari Vue
type ContactRequest struct {
    Name    string `json:"name"    binding:"required"`
    Email   string `json:"email"   binding:"required,email"`
    Message string `json:"message" binding:"required"`
}

func SendContact(c *gin.Context) {
    var req ContactRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := service.SendEmail(req.Name, req.Email, req.Message); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal kirim email"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Pesan berhasil dikirim"})
}