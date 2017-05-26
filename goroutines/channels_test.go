package goroutines

import "testing"

func TestChannels(t *testing.T) {
	Channel()
}

func TestChannelBuff(t *testing.T) {
	ChannelBuff()
}

func TestChannelSync(t *testing.T) {
	ChannelSync()
}

func TestChannelTrans(t *testing.T) {
	ChannelTrans()
}

func TestChannelTrans2(t *testing.T) {
	ChannelTrans2()
}

func TestChannelSelect(t *testing.T) {
	ChannelSelect()
}

func TestChannelTimeout(t *testing.T) {
	ChannelTimeout()
}

func TestChannelClose(t *testing.T) {
	ChannelClose()
}

func TestChannelRange(t *testing.T) {
	ChannelRange()
}

func TestChannelRange2(t *testing.T) {
	ChannelRange2()
}

func TestChannelPool(t *testing.T) {
	WorkerPool()
}
