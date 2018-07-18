package http

import (
	"github.com/canghai908/falcon-wechat/config"
	"github.com/toolkits/web/param"
	"gopkg.in/chanxuehong/wechat.v1/corp"
	"gopkg.in/chanxuehong/wechat.v1/corp/message/send"
	"log"
	"net/http"
	"strings"
)

func configProcRoutes() {

	http.HandleFunc("/wechat", func(w http.ResponseWriter, r *http.Request) {
		cfg := config.Config()
		token := param.String(r, "token", "")
		if cfg.Http.Token != token {
			http.Error(w, "no privilege", http.StatusForbidden)
			return
		}
		tos := param.MustString(r, "tos")
		content := param.MustString(r, "content")
		tos = strings.Replace(tos, ",", "|", -1)
		var TokenServer = corp.NewDefaultAccessTokenServer(cfg.Wechat.CorpId, cfg.Wechat.Secret, nil)
		body := new(send.Text)
		body.ToUser = tos
		body.MsgType = "text"
		body.AgentId = cfg.Wechat.AgentId
		body.Text.Content = content

		clt := send.NewClient(TokenServer, nil)
		if _, err := clt.SendText(body); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
		} else {
			http.Error(w, "success", http.StatusOK)
		}
	})

}
