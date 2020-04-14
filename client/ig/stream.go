package ig

import (
	"log"

	"github.com/gorilla/websocket"
)

func (ig *IGClient) Subscribe() {
	activeAccountId := ig.Session.AccountId
	password := "CST-" + ig.CST + "|XST-" + ig.XSecurityToken
	header := ig.newHeader(map[string]string{"password": password})
	//url := "wss" + ig.Session.LightstreamerEndpoint[5:]
	url := ig.Session.LightstreamerEndpoint
	log.Println(url)
	c, _, err := websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
}
