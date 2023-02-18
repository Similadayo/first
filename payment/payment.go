package payment

import (
	"log"
	"net/http"

	"github.com/Similadayo/db"
	"github.com/Similadayo/models"
	"github.com/anjolabassey/Rave-go/rave"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitializePayment(c *gin.Context) {
	Err := godotenv.Load()
	if Err != nil {
		log.Fatal("error loading env file...")
	}

	var r = rave.Rave{
		Live:      true,
		PublicKey: "",
		SecretKey: "",
	}

	var card = rave.Card{
		Rave: r,
	}

	var cardCharge rave.CardChargeData
	if err := c.ShouldBindJSON(&cardCharge); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	u := &models.User{}
	if err := db.DB.Where(u, user.ID).First(u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		log.Print(err)
		return
	}

	var order models.Order
	o := &models.Order{}
	if err := db.DB.Where(o, order.ID).First(o).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order not found"})
		return
	}

	cardCharge = rave.CardChargeData{
		Amount:        o.TotalCost,
		Email:         u.Email,
		CustomerPhone: u.Phone,
		Cardno:        cardCharge.Cardno,
		Cvv:           cardCharge.Cvv,
		Currency:      "NGN",
		Expirymonth:   cardCharge.Expirymonth,
		Expiryyear:    cardCharge.Expiryyear,
		Pin:           cardCharge.Pin,
		Txref:         cardCharge.Txref,
	}

	Paid, err := card.ChargeCard(cardCharge)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Payment made successfully", "status": Paid})
}
