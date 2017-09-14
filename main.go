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
						case "查角色":
							rtnMsg = "https://shawn8290.github.io/OPTC/index.html"
						case "查副本":
							rtnMsg = "http://jsfiddle.net/7ckc75ox/show/"
						case "909":
							rtnMsg = `【碳佐麻里日式燒肉 - 台南府前店】
地址：台南市安平區府前路二段630號`
						default:
							switch {
								case message.Text[:9] == "查角色":
									rtnMsg = "https://shawn8290.github.io/OPTC/index.html?" + message.Text[10:len(message.Text)]
								case message.Text[:6] == "測試":
									var syntax string = message.Text[10:len(message.Text)]
									syntax = strings.TrimSpace(syntax)
									rtnMsg = "https://shawn8290.github.io/OPTC/index.html?" + syntax
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
