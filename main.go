// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				var rtnMsg string = ""
				switch msgContent := strings.ToUpper(message.Text); msgContent {
					case "生日快樂":
						rtnMsg = `各位海賊們,

本週六版聚強勢登場！

►【期間限定】新甲奔活動－男子漢的一吃決勝負！
活動期間：07/22(13:30)~07/22(15:30)
挑戰任務：屋馬-國安店 甲奔
在「屋馬-國安店 甲奔」中，回合內給予店員造成的「麻煩值」之高分排名競賽！ 
競賽限制：限制祝壽 祝賀 抽卡 推坑 拜年 求明牌 問米 關落陰 起乩`
					case "屋馬地址":
						rtnMsg = "台中市西屯區台中市西屯區國安一路168號B1-2"
					case "屋馬電話":
						rtnMsg = "04-24652222"
					case "龍哥":
						rtnMsg = "就是任性"
					case "LEITO", "L":
						rtnMsg = "又!?"
					case "智障弟弟":
						rtnMsg = "leito 有人叫你"
					case "表演智障弟弟":
						rtnMsg = "不用表演，現在已經是"
					case "抽":
						rtnMsg = "抽"
					case "早", "早安":
						rtnMsg = "早安"
					case "課一單", "課十單":
						rtnMsg = "來 拿著神奇小卡 找龍哥幫你"
					case "森77":
						rtnMsg = "https://www.youtube.com/watch?v=TtQ9hwYoyWQ"
					case "槓", "靠北", "幹":
						rtnMsg = "造口業會抽不到限定唷"
				}
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(rtnMsg)).Do(); err != nil {
					log.Print(err)
				}
				
			}
		}
	}
}
