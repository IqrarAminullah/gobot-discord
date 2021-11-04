package message_utils

import (
	"strings"
)

//Receive a messag with the prefix
//Return command and params
func ParseMessage(m string, prefix string) (string, []string) {
	i := 0
	p := strings.Split(strings.TrimPrefix(m, prefix), " ")
	c := p[i]
	p = append(p[:i], p[i+1:]...)
	return c, p
}
