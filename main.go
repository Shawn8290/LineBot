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
	"strings"

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
				
				if message.Text[:4] == "echo" {
					rtnMsg = strings.ToUpper(message.Text[5:len(message.Text)])
				} else if message.Text[:3] == "len" {
					rtnMsg = len(message.Text)
				} else {
					var msgContent string := message.Text
					if (len(message.Text) > 1) {
						msgContent = strings.ToUpper(message.Text)
					}						
					switch msgContent{
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
						case "LEITO", "L", "l":
							rtnMsg = "又!?"
						case "智障弟弟":
							rtnMsg = "leito 有人叫你"
						case "表演智障弟弟":
							rtnMsg = "不用表演，現在已經是"
						case "抽":
							rtnMsg = "抽"
						case "早", "早安":
							rtnMsg = "早安"
						case "安":
							rtnMsg = "幾點了還在接龍"
						case "課一單", "課十單":
							rtnMsg = "來 拿著神奇小卡 找龍哥幫你"
						case "森77":
							rtnMsg = "https://www.youtube.com/watch?v=TtQ9hwYoyWQ"
						case "槓", "靠北", "幹", "靠":
							rtnMsg = "造口業會抽不到限定唷"
						case "棄坑", "放棄":
							rtnMsg = "來玩暗陰陽?"
						case "積人品":
							rtnMsg = "抽雷利、女帝?"
						case "南寮王":
							rtnMsg = "廢物替代役"
						case "戰國":
							rtnMsg = "琳抽到了"
						case "卡文":
							rtnMsg = "去跟姨借老公?"
						case "限羅":
							rtnMsg = "妞還沒抽到"
						case "新紅":
							rtnMsg = "獵人的情敵，但獵人已經贏了"
						case "女朋友":
							rtnMsg = "你跟南寮王都沒有"
						case "9487":
							rtnMsg = "94狂"
						case "我都沒有啦":
							rtnMsg = "信用卡交出來就有了"
						case "下禮拜就有了", "這禮拜就有了", "禮拜六就有了", "週六就有了":
							rtnMsg = "醒醒吧, 你抽不到的"
					}
				}				
				
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(rtnMsg)).Do(); err != nil {
					log.Print(err)
				}
				
			}
		}
	}
}
