package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"reflect"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/bath/comm"
	"github.com/davyxu/cellnet/codec"
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/cellnet/proc"
	"github.com/davyxu/golog"

	_ "github.com/davyxu/cellnet/codec/json"
	_ "github.com/davyxu/cellnet/peer/gorillaws"
	_ "github.com/davyxu/cellnet/proc/gorillaws"
)

var log = golog.New("websocket_server")

type EchoACK struct {
	Userid string
	Action string
	Value  string
	Sign   string
}

type EchoREQ struct {
	Userid   string
	Level    string
	Earnings string
	Action   string
	Value    string
	Sign     string
}

// type Info struct {
// 	userInfo  map[string]string
// 	itemsInfo map[string]string
// }

func (self *EchoACK) String() string { return fmt.Sprintf("%+v", *self) }
func (self *EchoREQ) String() string { return fmt.Sprintf("%+v", *self) }

// 将消息注册到系统
func init() {
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("json"),
		Type:  reflect.TypeOf((*EchoACK)(nil)).Elem(),
		ID:    1234,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("json"),
		Type:  reflect.TypeOf((*EchoREQ)(nil)).Elem(),
		ID:    1235,
	})
}

var (
	flagClient = flag.Bool("client", false, "client mode")
)

const (
	TestAddress = "http://127.0.0.1:18802/echo"
)

func server() {
	// 创建一个事件处理队列，整个服务器只有这一个队列处理事件，服务器属于单线程服务器
	queue := cellnet.NewEventQueue()

	// 侦听在18802端口
	p := peer.NewGenericPeer("gorillaws.Acceptor", "server", TestAddress, queue)

	proc.BindProcessorHandler(p, "gorillaws.ltv", func(ev cellnet.Event) {

		switch msg := ev.Message().(type) {
		case *cellnet.SessionAccepted:
			log.Debugln("server accepted")
			// 有连接断开
		case *cellnet.SessionClosed:
			log.Debugln("session closed: ", ev.Session().ID())
		case *EchoACK:
			log.Debugf("recv: %+v %v", msg, []byte("鲍勃"))
			val, exist := ev.Session().(cellnet.ContextSet).GetContext("request")
			if exist {
				if req, ok := val.(*http.Request); ok {
					raw, _ := json.Marshal(req.Header)
					log.Debugf("origin request header: %s", string(raw))
				}
			}

			switch msg.Action {
			case "init":
				var initInfo comm.Info
				userParams := []string{msg.Userid + ":bath", "level", "exp", "balance", "properties", "rec:1:num", "chr:1:num", "bap:1:num", "sau:1:num", "spy:1:num"}
				itemsParams := []string{"items", "rec:1", "chr:1", "bap:1", "spa:1", "sau:1"}
				initInfo.UserInfo = comm.GetHash(userParams)
				initInfo.ItemsInfo = comm.GetHash(itemsParams)
				ev.Session().Send(&EchoACK{
					Userid: msg.Userid,
					Action: msg.Action,
					Value:  comm.MapToJson(initInfo),
				})
			case "hincrby":
				comm.SetData("HINCRBY", msg.Userid+"_bath", msg.Sign, msg.Value)
				ev.Session().Send(&EchoACK{
					Userid: msg.Userid,
					Action: msg.Action,
					Value:  "true",
				})
			default:
				ev.Session().Send(&EchoACK{
					Userid: msg.Userid,
					Action: msg.Action,
					Value:  "true",
				})
			}
		case *EchoREQ:
			log.Debugf("recv: %+v %v", msg, []byte("鲍勃"))
			val, exist := ev.Session().(cellnet.ContextSet).GetContext("request")
			if exist {
				if req, ok := val.(*http.Request); ok {
					raw, _ := json.Marshal(req.Header)
					log.Debugf("origin request header: %s", string(raw))
				}
			}

			//var location, data = getBathInfo(msg.userid+"_bath", "level", "earnings", "wait", "tmp")
			comm.SetData("HSET", msg.Userid+"_bath", msg.Sign, msg.Value)
			ev.Session().Send(&EchoREQ{
				Userid:   msg.Userid,
				Action:   msg.Action,
				Level:    msg.Level,
				Earnings: msg.Earnings,
			})
		}
	})

	// 开始侦听
	p.Start()

	// 事件队列开始循环
	queue.StartLoop()

	// 阻塞等待事件队列结束退出( 在另外的goroutine调用queue.StopLoop() )
	queue.Wait()

}

// 默认启动服务器端
// 网页连接服务器： 在浏览器(Chrome)中打开index.html, F12打开调试窗口->Console标签 查看命令行输出
// 	注意：日志中的http://127.0.0.1:18802/echo链接是api地址，不是网页地址，直接打开无法正常工作
// 	注意：如果http代理/VPN在运行时可能会导致无法连接, 请关闭
// 客户端连接服务器：命令行模式中添加-client
func main() {
	flag.Parse()
	server()
}
