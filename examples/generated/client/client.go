// Code generated by wsgen. DO NOT EDIT.
package client

import (
    "time"
    "context"
    "github.com/Dafaque/ws-gen/examples/generated/iface"
    "github.com/Dafaque/ws-gen/examples/generated/mapper"
    "github.com/Dafaque/ws-gen/examples/generated/api"
    "github.com/Dafaque/ws-gen/examples/generated/model"

    "github.com/gorilla/websocket" 
)


func NewClient(
    url string,
    p *model.InitParams,
    mh api.MessageHandler,
    coder iface.Coder,
    logger iface.Logger,
) (*Client, error) {
    url += "?"
    url += p.ToQuery()
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
    cl := &Client{
		MessageSender: api.NewMessageSender(conn, coder),
		logger: logger,
		coder:  coder,
		conn:   conn,
        mh: mh,
	}
    go cl.rloop()
	return cl, nil
}

type Client struct {
    *api.MessageSender
    coder iface.Coder
    logger iface.Logger
    conn *websocket.Conn
    mh api.MessageHandler
    done bool
}

func (c *Client) Shutdown() {
	c.close(websocket.CloseNormalClosure, "")
}

func (c *Client) GetConn() *websocket.Conn {
    return c.conn
}
func (c *Client) GetContext() context.Context {
    return context.TODO()
}
func (c *Client) GetHandler() api.MessageHandler {
    return c.mh
}
func (c *Client) GetWriteChannel() chan interface{} {
    return nil //@todo
}
func (c *Client) GetCoder() iface.Coder {
    return c.coder
}
func (c *Client) GetLogger() iface.Logger {
    return c.logger
}
func (c *Client) CloseHandler(code int, reason string) error {
    c.done = true
	c.mh.OnDisconnected(code, reason)
    return nil
}
func (c *Client) close(code int, text string) {
    c.conn.WriteControl(
        websocket.CloseMessage,
        websocket.FormatCloseMessage(
            code,
            text,
        ),
        time.Now().Add(5*time.Second), //@todo write deadlines from config
    )
}
func (c *Client) rloop() {
    for {
        if c.done {
            break
        }
        if err := mapper.Read(c); err != nil {
            c.close(websocket.CloseProtocolError, err.Error())
        }
    }
}
