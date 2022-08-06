package rest

import (
	"GriffinBackend/price"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getBinanceTrade(c *gin.Context) {
	currNeed := []string{price.ETHEREUM, price.POLYGON, price.USDC}
	result := map[string]float64{}
	for _, m := range currNeed {
		p, err := price.BinancePrice(m)
		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"message": PRICE_GET_ERROR,
			})
			return
		}
		result[m] = p
	}

	// send price information
	p := price.PriceInformation{
		Ethereum: result[price.ETHEREUM],
		Polygon:  result[price.POLYGON],
		USDC:     result[price.USDC],
	}
	c.JSON(http.StatusOK, p)
}
