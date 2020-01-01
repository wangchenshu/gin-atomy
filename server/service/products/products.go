package products

import (
	"fmt"
	"log"
	"strconv"
	"strings"

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
				profile, err := myBot.GetProfile(event.Source.UserID).Do()
				if err != nil {
					log.Print(err)
				}

				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					matchText := message.Text
					buyProductName := ""
					if strings.Contains(matchText, "+1") {
						buyProductName = strings.Split(matchText, ",")[1]
						matchText = strings.Split(matchText, ",")[0]
					}

					switch matchText {
					case "+1":
						replyMsg := "+1 失敗"
						if addCarts(profile.DisplayName, buyProductName, 1) {
							replyMsg = " " + buyProductName + " +1 成功"
						}
						if _, err := myBot.ReplyMessage(
							event.ReplyToken,
							linebot.NewTextMessage(profile.DisplayName+replyMsg),
							myQuickReply(),
						).Do(); err != nil {
							log.Print(err)
						}
					case "過年團購":
						groupBuyNames := []string{"香烤海苔(小片裝)", "幸福堅果", "好纖果乾"}
						wannaBuyStr := "我想+1"
						altText := "團購人氣商品"
						arouselColumns := []*linebot.CarouselColumn{}

						for k, _ := range groupBuyNames {
							product := getProductLike(groupBuyNames[k])
							arouselColumns = append(arouselColumns,
								newCarouselColumn(
									// imageURLs[k],
									product.Pic,
									product.Name,
									"$ "+product.Price+" PV: "+product.Point,
									wannaBuyStr,
									"+1,"+groupBuyNames[k],
								),
							)
						}
						template := linebot.NewCarouselTemplate(arouselColumns...)

						if _, err := myBot.ReplyMessage(
							event.ReplyToken,
							linebot.NewTemplateMessage(altText, template),
							myQuickReply(),
						).Do(); err != nil {
							log.Print(err)
						}
					case "查看購物車":
						carts := getCartsByUsername(profile.DisplayName)
						repMsg := "目前購物車有: \n\n"
						cartIsEmpty := "購物車是空的"
						totalPrice := 0

						for _, cart := range carts {
							cartQtyStr := fmt.Sprintf("%d", cart.Qty)
							repMsg += cart.ProductName + ", 價格: $" + cart.Price + ", 數量: " + cartQtyStr + "\n"
							cartPrice, err := strconv.ParseInt(cart.Price, 10, 64)
							if err == nil {
								fmt.Println(cartPrice)
							}
							totalPrice += int(cartPrice) * cart.Qty
						}
						repMsg += "\n總計: $ " + fmt.Sprintf("%d", totalPrice) + "\n"
						if totalPrice == 0 {
							repMsg = cartIsEmpty
						}
						if _, err := myBot.ReplyMessage(
							event.ReplyToken,
							linebot.NewTextMessage(repMsg),
							myQuickReply(),
						).Do(); err != nil {
							log.Print(err)
						}
					case "清除購物車":
						if clearCarts(profile.DisplayName) {
							repMsg := "清除成功"
							if _, err := myBot.ReplyMessage(
								event.ReplyToken,
								linebot.NewTextMessage(repMsg),
								myQuickReply(),
							).Do(); err != nil {
								log.Print(err)
							}
						}
					case "結帳":
						carts := getCartsByUsername(profile.DisplayName)
						repMsg := "明細如下: \n\n"
						cartIsEmpty := "購物車是空的"
						totalPrice := 0

						for _, cart := range carts {
							cartQtyStr := fmt.Sprintf("%d", cart.Qty)
							repMsg += cart.ProductName + ", 價格: $" + cart.Price + ", 數量: " + cartQtyStr + "\n"
							cartPrice, err := strconv.ParseInt(cart.Price, 10, 64)
							if err == nil {
								fmt.Println(cartPrice)
							}
							totalPrice += int(cartPrice) * cart.Qty
						}
						repMsg += "\n總計: $ " + fmt.Sprintf("%d", totalPrice) + "\n"
						repMsg += "結帳完成，請記得於時間內付款，謝謝您\n"
						if totalPrice == 0 {
							repMsg = cartIsEmpty
						}
						clearCarts(profile.DisplayName)
						if _, err := myBot.ReplyMessage(
							event.ReplyToken,
							linebot.NewTextMessage(repMsg),
							myQuickReply(),
						).Do(); err != nil {
							log.Print(err)
						}
					case "flextest":
						contents := &linebot.BubbleContainer{
							Type: linebot.FlexContainerTypeBubble,
							Body: &linebot.BoxComponent{
								Type:   linebot.FlexComponentTypeBox,
								Layout: linebot.FlexBoxLayoutTypeHorizontal,
								Contents: []linebot.FlexComponent{
									&linebot.TextComponent{
										Type: linebot.FlexComponentTypeText,
										Text: "Hello,",
									},
									&linebot.TextComponent{
										Type: linebot.FlexComponentTypeText,
										Text: "World!",
									},
								},
							},
						}
						if _, err := myBot.ReplyMessage(
							event.ReplyToken,
							linebot.NewFlexMessage("Flex message alt text", contents),
							myQuickReply(),
						).Do(); err != nil {
							log.Print(err)
						}
					case "flextest2":
						contents := &linebot.CarouselContainer{
							Type: linebot.FlexContainerTypeCarousel,
							Contents: []*linebot.BubbleContainer{
								{
									Type: linebot.FlexContainerTypeBubble,
									Body: &linebot.BoxComponent{
										Type:   linebot.FlexComponentTypeBox,
										Layout: linebot.FlexBoxLayoutTypeVertical,
										Contents: []linebot.FlexComponent{
											&linebot.TextComponent{
												Type: linebot.FlexComponentTypeText,
												Text: "First bubble",
											},
										},
									},
								},
								{
									Type: linebot.FlexContainerTypeBubble,
									Body: &linebot.BoxComponent{
										Type:   linebot.FlexComponentTypeBox,
										Layout: linebot.FlexBoxLayoutTypeVertical,
										Contents: []linebot.FlexComponent{
											&linebot.TextComponent{
												Type: linebot.FlexComponentTypeText,
												Text: "Second bubble",
											},
										},
									},
								},
							},
						}
						if _, err := myBot.ReplyMessage(
							event.ReplyToken,
							linebot.NewFlexMessage("Flex message alt text", contents),
							myQuickReply(),
						).Do(); err != nil {
							log.Print(err)
						}
					case "flexjson":
						jsonString := `{
							"type": "bubble",
							"hero": {
							  "type": "image",
							  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/01_1_cafe.png",
							  "size": "full",
							  "aspectRatio": "20:13",
							  "aspectMode": "cover",
							  "action": {
								"type": "uri",
								"uri": "http://linecorp.com/"
							  }
							},
							"body": {
							  "type": "box",
							  "layout": "vertical",
							  "contents": [
								{
								  "type": "text",
								  "text": "Brown Cafe",
								  "weight": "bold",
								  "size": "xl"
								},
								{
								  "type": "box",
								  "layout": "baseline",
								  "margin": "md",
								  "contents": [
									{
									  "type": "icon",
									  "size": "sm",
									  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
									},
									{
									  "type": "icon",
									  "size": "sm",
									  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
									},
									{
									  "type": "icon",
									  "size": "sm",
									  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
									},
									{
									  "type": "icon",
									  "size": "sm",
									  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
									},
									{
									  "type": "icon",
									  "size": "sm",
									  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gray_star_28.png"
									},
									{
									  "type": "text",
									  "text": "4.0",
									  "size": "sm",
									  "color": "#999999",
									  "margin": "md",
									  "flex": 0
									}
								  ]
								},
								{
								  "type": "box",
								  "layout": "vertical",
								  "margin": "lg",
								  "spacing": "sm",
								  "contents": [
									{
									  "type": "box",
									  "layout": "baseline",
									  "spacing": "sm",
									  "contents": [
										{
										  "type": "text",
										  "text": "Place",
										  "color": "#aaaaaa",
										  "size": "sm",
										  "flex": 1
										},
										{
										  "type": "text",
										  "text": "Miraina Tower, 4-1-6 Shinjuku, Tokyo",
										  "wrap": true,
										  "color": "#666666",
										  "size": "sm",
										  "flex": 5
										}
									  ]
									},
									{
									  "type": "box",
									  "layout": "baseline",
									  "spacing": "sm",
									  "contents": [
										{
										  "type": "text",
										  "text": "Time",
										  "color": "#aaaaaa",
										  "size": "sm",
										  "flex": 1
										},
										{
										  "type": "text",
										  "text": "10:00 - 23:00",
										  "wrap": true,
										  "color": "#666666",
										  "size": "sm",
										  "flex": 5
										}
									  ]
									}
								  ]
								}
							  ]
							},
							"footer": {
							  "type": "box",
							  "layout": "vertical",
							  "spacing": "sm",
							  "contents": [
								{
								  "type": "button",
								  "style": "link",
								  "height": "sm",
								  "action": {
									"type": "uri",
									"label": "CALL",
									"uri": "https://linecorp.com"
								  }
								},
								{
								  "type": "button",
								  "style": "link",
								  "height": "sm",
								  "action": {
									"type": "uri",
									"label": "WEBSITE",
									"uri": "https://linecorp.com",
									"altUri": {
									  "desktop": "https://line.me/ja/download"
									}
								  }
								},
								{
								  "type": "spacer",
								  "size": "sm"
								}
							  ],
							  "flex": 0
							}
						  }`
						contents, err := linebot.UnmarshalFlexMessageJSON([]byte(jsonString))
						if err != nil {
							log.Print(err)
						}
						if _, err := myBot.ReplyMessage(
							event.ReplyToken,
							linebot.NewFlexMessage("Flex message alt text", contents),
							myQuickReply(),
						).Do(); err != nil {
							log.Print(err)
						}
					case "喜好測試":
						template := linebot.NewConfirmTemplate(
							"喜愛艾多美貼布?",
							linebot.NewMessageAction("是", "Like,艾多美貼布"),
							linebot.NewMessageAction("否", "NoLike,艾多美貼布"),
						)
						if _, err := myBot.ReplyMessage(
							event.ReplyToken,
							linebot.NewTemplateMessage("喜好測試", template),
							myQuickReply(),
						).Do(); err != nil {
							log.Print(err)
						}
					default:
						wannaBuyStr := "我想+1"
						products := getProductsLike(message.Text, 10)
						arouselColumns := []*linebot.CarouselColumn{}
						var template *linebot.CarouselTemplate
						var templateMessage *linebot.TemplateMessage
						templateMessage = linebot.NewTemplateMessage("只要輸入 艾多美新品名稱", myMenuTemplate())

						if len(products) > 0 {
							for _, product := range products {
								arouselColumns = append(arouselColumns,
									newCarouselColumn(
										product.Pic,
										product.Name,
										"$ "+product.Price+" PV: "+product.Point,
										wannaBuyStr,
										"+1,"+product.Name,
									),
								)
							}
							template = linebot.NewCarouselTemplate(arouselColumns...)
							templateMessage = linebot.NewTemplateMessage("您想找的商品", template)
						} else {
							templateMessage = linebot.NewTemplateMessage("請輸入商品關鍵字", myMenuTemplate())
						}

						if _, err := myBot.ReplyMessage(
							event.ReplyToken,
							templateMessage,
							myQuickReply(),
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

func newCarouselColumn(imageURL, title, text string, actionLabel string, actionText string) *linebot.CarouselColumn {
	return linebot.NewCarouselColumn(
		imageURL, title, text,
		linebot.NewMessageAction(actionLabel, actionText),
	)
}

func getAllProducts() []model.Products {
	products := []model.Products{}
	db.Db.Find(&products)

	return products
}

func getProductsLike(name string, limit int) []model.Products {
	products := []model.Products{}
	db.Db.Where("name LIKE ? AND pic != ?", "%"+name+"%", "NULL").Limit(limit).Find(&products)

	return products
}

func getProductLike(name string) model.Products {
	product := model.Products{}
	db.Db.Where("name LIKE ? AND pic != ?", "%"+name+"%", "NULL").First(&product)

	return product
}

func newMessageAction(label string, text string) *linebot.MessageAction {
	return linebot.NewMessageAction(label, text)
}

func myMenuTemplate() *linebot.ButtonsTemplate {
	imageURL := "https://firebasestorage.googleapis.com/v0/b/atomy-bot.appspot.com/o/%E8%89%BE%E5%A4%9A%E7%BE%8E%20%E8%89%BE%E4%B8%8D%E9%87%8B%E6%89%8B%E7%B6%93%E5%85%B8%E7%B3%BB%E5%88%97.jpg?alt=media&token=ebbb8e46-6d86-4229-9e29-67b4858ea0d1"
	title := "艾多美新品 Go Bot 1.0 "
	text := "只要輸入 艾多美新品名稱"
	messageSlices := []string{"艾多美 艾不釋手經典系列", "艾多美 托特包", "艾多美 物理性防曬膏", "艾多美 新春紅包袋"}
	template := []linebot.TemplateAction{}

	for _, v := range messageSlices {
		template = append(template, newMessageAction(v, v))
	}

	menuTemplate := linebot.NewButtonsTemplate(
		imageURL, title, text, template...,
	)

	return menuTemplate
}

func myQuickReply() linebot.SendingMessage {
	content := "快速選單或輸入商品關鍵字"
	imageURLs := []string{
		"https://firebasestorage.googleapis.com/v0/b/atomy-bot.appspot.com/o/%E8%89%BE%E5%A4%9A%E7%BE%8E%20%E7%89%A9%E7%90%86%E6%80%A7%E9%98%B2%E6%9B%AC%E8%86%8F.jpg?alt=media&token=e659398b-c5a5-4e0e-ae91-614633d2355b",
		"https://firebasestorage.googleapis.com/v0/b/atomy-bot.appspot.com/o/%E8%89%BE%E5%A4%9A%E7%BE%8E%20%E7%89%A9%E7%90%86%E6%80%A7%E9%98%B2%E6%9B%AC%E8%86%8F.jpg?alt=media&token=e659398b-c5a5-4e0e-ae91-614633d2355b",
		"https://firebasestorage.googleapis.com/v0/b/atomy-bot.appspot.com/o/%E6%B5%B7%E8%8B%94%E7%A6%AE%E7%9B%92.jpg?alt=media&token=4e1e859f-fae6-41de-86f4-94a506c3a2a9",
		"https://firebasestorage.googleapis.com/v0/b/atomy-bot.appspot.com/o/%E6%B5%B7%E8%8B%94%E7%A6%AE%E7%9B%92.jpg?alt=media&token=4e1e859f-fae6-41de-86f4-94a506c3a2a9",
		"https://firebasestorage.googleapis.com/v0/b/atomy-bot.appspot.com/o/%E5%B9%B8%E7%A6%8F%E5%A0%85%E6%9E%9C.jpg?alt=media&token=9f409ba8-5508-46f2-8420-b74eff83258c",
		"https://firebasestorage.googleapis.com/v0/b/atomy-bot.appspot.com/o/%E5%A5%BD%E7%BA%96%E6%9E%9C%E4%B9%BE.jpg?alt=media&token=6e892755-4e05-4f3b-881b-c127e059a24b",
		"https://firebasestorage.googleapis.com/v0/b/atomy-bot.appspot.com/o/%E8%89%BE%E5%A4%9A%E7%BE%8E%20%E7%89%A9%E7%90%86%E6%80%A7%E9%98%B2%E6%9B%AC%E8%86%8F.jpg?alt=media&token=e659398b-c5a5-4e0e-ae91-614633d2355b",
	}
	labels := []string{"查看購物車", "結帳", "過年團購", "喜好測試", "幸福堅果", "好纖果乾", "清除購物車"}
	quickReplyButtons := []*linebot.QuickReplyButton{}

	for k, v := range labels {
		quickReplyButtons = append(quickReplyButtons, linebot.NewQuickReplyButton(
			imageURLs[k], linebot.NewMessageAction(v, v),
		))
	}
	quickReply := linebot.NewTextMessage(content).
		WithQuickReplies(linebot.NewQuickReplyItems(quickReplyButtons...))

	return quickReply
}

func addCarts(username string, productName string, qty int) bool {
	carts := model.Carts{}
	product := getProductLike(productName)

	if product.ID != 0 {
		carts.Username = username
		carts.ProductName = product.Name
		carts.Qty = qty
		carts.Price = product.Price

		db.Db.Create(&carts)
		return true
	}

	return false
}

func getCartsByUsername(username string) []model.Carts {
	carts := []model.Carts{}
	db.Db.Where("username = ?", username).Find(&carts)

	return carts
}

func clearCarts(username string) bool {
	db.Db.Delete(model.Carts{}, "username = ?", username)
	return true
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
