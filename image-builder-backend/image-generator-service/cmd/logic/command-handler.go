package logic

import "time"

type CommandHandler struct{}

func (c *CommandHandler) RunCommand(_ string, _ ...string) error {
	// TODO: implement
	time.Sleep(5 * time.Second)
	return nil
}
