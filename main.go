package main

import (
	"github.com/kataras/iris"
	"github.com/iris-contrib/middleware/recovery"
	"os"
	"haserver/hamessage"
	"github.com/golang/protobuf/proto"
)

type clientPage struct {
	Title string
	Host string
}


var users = map[string]string{}

func main()  {

	boticon := "https://hachatserver.herokuapp.com/img/Icon.png"
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	iris.Use(recovery.New())
	iris.Static("/js","./static/js",1)
	iris.Static("/img","./static/img",1)

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("client.html", clientPage{"Chat test", ctx.HostString()})
	})

	iris.Config.Websocket.Endpoint = "hachat"
	iris.Config.Websocket.MaxMessageSize = 1024*1024
	iris.Config.Websocket.BinaryMessages = true


	var haroom = "room1"
	iris.Websocket.OnConnection(func(c iris.WebsocketConnection) {
		c.Join(haroom)

		c.On("chat", func(message string) {
			c.To(iris.All).Emit("chat", "From: "+c.ID()+": "+message)
		})

		c.OnMessage(func(bts []byte) {
			message := &hamessage.Hamessage{}
			_ = proto.Unmarshal(bts, message)

			if message.Event == "join" {
				users[c.ID()] = message.Name

				welcomemsg := &hamessage.Hamessage{}
				welcomemsg.Event = "chat"
				welcomemsg.Name = "habot"
				welcomemsg.Text = message.Name + " join"
				welcomemsg.Profileimage = boticon
				out,_ := proto.Marshal(welcomemsg)
				c.To(iris.All).EmitMessage(out)
			} else if message.Event == "chat" {
				c.To(haroom).EmitMessage(bts)
			}

		})

		c.OnDisconnect(func() {
			name := users[c.ID()]
			welcomemsg := &hamessage.Hamessage{}
			welcomemsg.Event = "chat"
			welcomemsg.Name = "habot"
			welcomemsg.Text = "Ha! bye "+ name
			welcomemsg.Profileimage = boticon
			out,_ := proto.Marshal(welcomemsg)
			c.To(iris.All).EmitMessage(out)
			delete(users, c.ID())

		})
	})

	iris.Listen(":"+port)
}