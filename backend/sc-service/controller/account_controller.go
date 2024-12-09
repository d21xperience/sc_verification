package controller

import (
	"net/http"
	"sc-service/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AccountController struct {
	DB *gorm.DB
}

func NewAccountController(db *gorm.DB) *AccountController {
	return &AccountController{DB: db}
}

// Tambah akun baru
func (ac *AccountController) CreateAccount(c *gin.Context) {
	var account model.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Panggil fungsi untuk menambahkan akun baru blockhain
	// walletInfo := ethbc.ETHClientInfo{
	// 	NetURL: "",
	// }

	if err := ac.DB.Create(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, account)
}

// Ambil semua akun
func (ac *AccountController) GetAccounts(c *gin.Context) {
	// var accounts []model.Account
	var blockchainNetwork []model.BlockchainNetwork
	// if err := ac.DB.Find(&accounts).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	if err := ac.DB.Preload("Account").Find(&blockchainNetwork).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blockchainNetwork)
}

// Ambil satu akun berdasarkan ID
func (ac *AccountController) GetAccountByID(c *gin.Context) {
	id := c.Param("id")
	var account model.Account
	if err := ac.DB.First(&account, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}
	c.JSON(http.StatusOK, account)
}

// Perbarui akun
func (ac *AccountController) UpdateAccount(c *gin.Context) {
	id := c.Param("id")
	var account model.Account

	if err := ac.DB.First(&account, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ac.DB.Save(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

// Hapus akun
func (ac *AccountController) DeleteAccount(c *gin.Context) {
	id := c.Param("id")
	if err := ac.DB.Delete(&model.Account{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Account deleted"})
}

// // Set Active account
// func (ac *AccountController) SetAccountActive(c *gin.Context) {

// }

// // Set Import account
// func (ac *AccountController) ImportAccount(c *gin.Context) {

// }
