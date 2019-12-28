package products

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"

	"net/http"

	"gin-atomy/server/db"
	"gin-atomy/server/model"
	"gin-atomy/server/service/mylinebot"
)

var myBot = mylinebot.Init()

func GetAtomyProducts() gin.HandlerFunc {
	return func(context *gin.Context) {
		events, err := myBot.ParseRequest(context.Request)
		fmt.Println(events)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				context.JSON(http.StatusBadRequest, nil)
			} else {
				context.JSON(500, nil)
			}
			return
		}

		for _, event := range events {
			switch event.Type {
			case linebot.EventTypeMessage:
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					repMsg := "艾多美商品: \n\n"
					products := getProductsLike(message.Text, 10)

					for _, product := range products {
						repMsg += "名稱: " + product.Name + "\n"
						repMsg += "價格: " + product.Price + "\n"
						repMsg += "PV: " + product.Point + "\n"
					}

					imageURL := "https://firebasestorage.googleapis.com/v0/b/atomy-bot.appspot.com/o/%E8%89%BE%E5%A4%9A%E7%BE%8E%20%E8%89%BE%E4%B8%8D%E9%87%8B%E6%89%8B%E7%B6%93%E5%85%B8%E7%B3%BB%E5%88%97.jpg?alt=media&token=ebbb8e46-6d86-4229-9e29-67b4858ea0d1"
					template := linebot.NewButtonsTemplate(
						imageURL, "艾多美新品 Go Bot 1.0 ", "只要輸入 艾多美新品名稱",
						linebot.NewMessageAction("艾多美 艾不釋手經典系列", "艾多美 艾不釋手經典系列"),
						linebot.NewMessageAction("艾多美 托特包", "艾多美 托特包"),
						linebot.NewMessageAction("艾多美 物理性防曬膏", "艾多美 物理性防曬膏"),
						linebot.NewMessageAction("艾多美 新春紅包袋", "艾多美 新春紅包袋"),
					)
					if _, err := myBot.ReplyMessage(
						event.ReplyToken,
						linebot.NewTemplateMessage("只要輸入 艾多美新品名稱", template),
						linebot.NewTextMessage(repMsg),
					).Do(); err != nil {
						log.Print(err)
					}

				case *linebot.ImageMessage:
					log.Print(message)
				case *linebot.VideoMessage:
					log.Print(message)
				case *linebot.AudioMessage:
					log.Print(message)
				case *linebot.FileMessage:
					log.Print(message)
				case *linebot.LocationMessage:
					log.Print(message)
				case *linebot.StickerMessage:
					log.Print(message)
				default:
					log.Printf("Unknown message: %v", message)
				}
			default:
				log.Printf("Unknown event: %v", event)
			}
		}
		context.JSON(http.StatusOK, gin.H{
			"success": events,
		})
	}
}

func getAllProducts() []model.Products {
	products := []model.Products{}
	db.Db.Find(&products)

	return products
}

func getProductsLike(name string, limit int) []model.Products {
	products := []model.Products{}
	db.Db.Where("name LIKE ?", "%"+name+"%").Limit(limit).Find(&products)

	return products
}

func GetProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, getAllProducts())
	}
}

func GetProductsLike(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, getProductsLike(name, 100))
	}
}
