package notify

import (
	"github.com/gaurishhs/vox"
)

func SetupListener() *vox.Subscriber {
  subscriber := vox.NewSubscriber()
  
  subscriber.Listen(Listener)

  return subscriber
}

func Listener(message *vox.Message) {
  // Notify services
}
