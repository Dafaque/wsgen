// Code generated by wsgen. DO NOT EDIT.
package api
import (
    "log"
    "errors"
    "context"
    "wsgen/examples/gen/model"

)
type MessageHandler interface {
    OnConnected(context.Context, *MessageSender)
    OnDisconnected()
    OnTextMessage(context.Context, model.TextMessage, *MessageSender) error
    OnChatEvent(context.Context, model.ChatEvent, *MessageSender) error
}

type UnimplementedMessageHandler struct {}
func (u UnimplementedMessageHandler) OnConnected(ctx context.Context, sender *MessageSender) {
    log.Println("OnConnected is not implemented")
}
func (u UnimplementedMessageHandler) OnDisconnected() {
    log.Println("OnDisconnected is not implemented")
}
func (u UnimplementedMessageHandler) OnTextMessage(ctx context.Context, msg model.TextMessage, sender *MessageSender) error {
    return errors.New("OnTextMessage is not implemented")
}
func (u UnimplementedMessageHandler) OnChatEvent(ctx context.Context, msg model.ChatEvent, sender *MessageSender) error {
    return errors.New("OnChatEvent is not implemented")
}
