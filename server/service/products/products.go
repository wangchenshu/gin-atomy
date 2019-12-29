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
				imageURL1 := "https://firebasestorage.googleapis.com/v0/b/atomy-bot.appspot.com/o/%E6%B5%B7%E8%8B%94%E7%A6%AE%E7%9B%92.jpg?alt=media&token=4e1e859f-fae6-41de-86f4-94a506c3a2a9"
				imageURL2 := "https://firebasestorage.googleapis.com/v0/b/atomy-bot.appspot.com/o/%E5%B9%B8%E7%A6%8F%E5%A0%85%E6%9E%9C.jpg?alt=media&token=9f409ba8-5508-46f2-8420-b74eff83258c"
				imageURL3 := "https://firebasestorage.googleapis.com/v0/b/atomy-bot.appspot.com/o/%E5%A5%BD%E7%BA%96%E6%9E%9C%E4%B9%BE.jpg?alt=media&token=6e892755-4e05-4f3b-881b-c127e059a24b"

				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					switch message.Text {
					case "+1,香烤海苔(小片裝)":
						if _, err := myBot.ReplyMessage(
							event.ReplyToken,
							linebot.NewTextMessage("香烤海苔(小片裝) +1成功"),
							linebot.NewTextMessage("過年團購商品").
								WithQuickReplies(linebot.NewQuickReplyItems(
									linebot.NewQuickReplyButton(
										imageURL1,
										linebot.NewMessageAction("查看團購商品", "過年團購")),
									linebot.NewQuickReplyButton(
										imageURL1,
										linebot.NewMessageAction("香烤海苔(小片裝)", "香烤海苔(小片裝)")),
									linebot.NewQuickReplyButton(
										imageURL2,
										linebot.NewMessageAction("幸福堅果", "幸福堅果")),
									linebot.NewQuickReplyButton(
										imageURL3,
										linebot.NewMessageAction("好纖果乾", "好纖果乾")),
								)),
						).Do(); err != nil {
							log.Print(err)
						}
					case "+1,幸福堅果":
						if _, err := myBot.ReplyMessage(
							event.ReplyToken,
							linebot.NewTextMessage("幸福堅果 +1成功"),
							linebot.NewTextMessage("過年團購商品").
								WithQuickReplies(linebot.NewQuickReplyItems(
									linebot.NewQuickReplyButton(
										imageURL1,
										linebot.NewMessageAction("查看團購商品", "過年團購")),
									linebot.NewQuickReplyButton(
										imageURL1,
										linebot.NewMessageAction("香烤海苔(小片裝)", "香烤海苔(小片裝)")),
									linebot.NewQuickReplyButton(
										imageURL2,
										linebot.NewMessageAction("幸福堅果", "幸福堅果")),
									linebot.NewQuickReplyButton(
										imageURL3,
										linebot.NewMessageAction("好纖果乾", "好纖果乾")),
								)),
						).Do(); err != nil {
							log.Print(err)
						}
					case "+1,好纖果乾":
						if _, err := myBot.ReplyMessage(
							event.ReplyToken,
							linebot.NewTextMessage("好纖果乾 +1成功"),
							linebot.NewTextMessage("過年團購商品").
								WithQuickReplies(linebot.NewQuickReplyItems(
									linebot.NewQuickReplyButton(
										imageURL1,
										linebot.NewMessageAction("查看團購商品", "過年團購")),
									linebot.NewQuickReplyButton(
										imageURL1,
										linebot.NewMessageAction("香烤海苔(小片裝)", "香烤海苔(小片裝)")),
									linebot.NewQuickReplyButton(
										imageURL2,
										linebot.NewMessageAction("幸福堅果", "幸福堅果")),
									linebot.NewQuickReplyButton(
										imageURL3,
										linebot.NewMessageAction("好纖果乾", "好纖果乾")),
								)),
						).Do(); err != nil {
							log.Print(err)
						}
					case "過年團購":
						products_1 := getGroupBuy("香烤海苔(小片裝)")
						products_2 := getGroupBuy("幸福堅果")
						products_3 := getGroupBuy("好纖果乾")

						template := linebot.NewCarouselTemplate(
							linebot.NewCarouselColumn(
								imageURL1, products_1.Name, products_1.Price,
								linebot.NewMessageAction("我想+1", "+1,香烤海苔(小片裝)"),
							),
							linebot.NewCarouselColumn(
								imageURL2, products_2.Name, products_2.Price,
								linebot.NewMessageAction("我想+1", "+1,幸福堅果"),
							),
							linebot.NewCarouselColumn(
								imageURL3, products_3.Name, products_3.Price,
								linebot.NewMessageAction("我想+1", "+1,好纖果乾"),
							),
						)
						if _, err := myBot.ReplyMessage(
							event.ReplyToken,
							linebot.NewTemplateMessage("團購人氣商品", template),
							linebot.NewTextMessage("年節團購商品").
								WithQuickReplies(linebot.NewQuickReplyItems(
									linebot.NewQuickReplyButton(
										imageURL1,
										linebot.NewMessageAction("查看團購商品", "過年團購")),
									linebot.NewQuickReplyButton(
										imageURL1,
										linebot.NewMessageAction("香烤海苔(小片裝)", "香烤海苔(小片裝)")),
									linebot.NewQuickReplyButton(
										imageURL2,
										linebot.NewMessageAction("幸福堅果", "幸福堅果")),
									linebot.NewQuickReplyButton(
										imageURL3,
										linebot.NewMessageAction("好纖果乾", "好纖果乾")),
								)),
						).Do(); err != nil {
							log.Print(err)
						}
					default:
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
							linebot.NewTextMessage("過年團購商品").
								WithQuickReplies(linebot.NewQuickReplyItems(
									linebot.NewQuickReplyButton(
										imageURL1,
										linebot.NewMessageAction("查看團購商品", "過年團購")),
									linebot.NewQuickReplyButton(
										imageURL1,
										linebot.NewMessageAction("香烤海苔(小片裝)", "香烤海苔(小片裝)")),
									linebot.NewQuickReplyButton(
										imageURL2,
										linebot.NewMessageAction("幸福堅果", "幸福堅果")),
									linebot.NewQuickReplyButton(
										imageURL3,
										linebot.NewMessageAction("好纖果乾", "好纖果乾")),
								)),
						).Do(); err != nil {
							log.Print(err)
						}
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

func getGroupBuy(name string) model.Products {
	product := model.Products{}
	db.Db.Where("name LIKE ?", "%"+name+"%").First(&product)

	return product
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
