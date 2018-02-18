/*
  Google Cloud SDK の機能で開発サーバを開始する場合のコード
  dev_appserver.py .
  あるいは
  goapp serve
  で開始
*/

package app

import (
	"bitbucket.org/kikkikkikkik/mypage/server"
)

func init() {
	server.Init()
}
