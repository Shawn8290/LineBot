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
	"strconv"

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
					
				if message.Text[:1] == "@" {		
					cmd := strings.Split(message.Text, " ")
					switch cmd[0] {
						case "@echo":
							rtnMsg = message.Text[6:len(message.Text)]
						case "@len":
							rtnMsg = strconv.Itoa(len(message.Text) - 5)
						case "@userid":
							rtnMsg = event.Source.UserID
						case "@groupid":
							rtnMsg = event.Source.GroupID
						case "@roomid":
							rtnMsg = event.Source.RoomID
					}	
				} else {
					if (os.Getenv("EnableGroup") != event.Source.GroupID) && (len(event.Source.GroupID) > 0) {
						return
					}
					var msgContent string = strings.ToUpper(message.Text)						
					switch msgContent {
						case "生日快樂":
							rtnMsg = "還在生什麼日快什麼樂，要叫店員來嗎"
						case "龍哥":
							rtnMsg = "就是任性"
						case "抽":
							rtnMsg = "抽"
						case "安":
							rtnMsg = "幾點了還在接龍"
						case "課一單", "課十單":
							rtnMsg = "來 拿著神奇小卡 找龍哥幫你"
						case "森77":
							rtnMsg = "https://www.youtube.com/watch?v=TtQ9hwYoyWQ"
						case "槓", "靠北", "幹", "靠":
							rtnMsg = "造口業會抽不到限定唷"
						case "棄坑":
							rtnMsg = "來玩暗陰陽?"
						case "積人品":
							rtnMsg = "去隔壁發片子啊"
						case "戰國":
							rtnMsg = "琳 於 2017/05/29 已抽中"
						case "卡文":
							rtnMsg = "姨 於 2017/05/27 已抽中"
						case "限羅":
							rtnMsg = "妞 於 2017/07/29 已抽中"
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
						case "查角色":
							rtnMsg = "https://shawn8290.github.io/OPTC/index.html"
						case "查副本":
							rtnMsg = "http://jsfiddle.net/7ckc75ox/show/"
						case "永遠都抽不到":
							rtnMsg = "抽不到就是儲不夠"
						case "How 歐 are you":
							rtnMsg = "I am 非, 3Q"
						case "夏威夷":
							rtnMsg = "南寮自己去了"
						case "強力慶典":
							rtnMsg = "先課個全餐吧"
						case "手殘", "翻了", "翻船", "手滑了":
							rtnMsg = "大俠請重新來過"
						case "不抽了":
							rtnMsg = "現在不抽限定就被別人抽走了"
						case "跳過", "放棄":
							rtnMsg = "現在放棄比賽就結束了"
						case "博識隊長":
							rtnMsg = "陳勁宇？"	
						default:
							switch {
								case message.Text[:12] == "岳父大人":
									rtnMsg = "抓到了嗎？ID: " + event.Source.UserID
								case message.Text[:9] == "查角色":
									rtnMsg = "https://shawn8290.github.io/OPTC/index.html?" + message.Text[10:len(message.Text)]
							}							
					}
				}				
				
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(rtnMsg)).Do(); err != nil {
					log.Print(err)
				}
				
			}
		}
	}
}
