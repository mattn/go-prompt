// +build windows

package prompt

import (
	"log"
	"time"
	"unicode/utf8"
)

func (p *Prompt) readBuffer(bufCh chan []byte, stopCh chan struct{}) {
	buf := make([]byte, 1024)

	log.Printf("[INFO] readBuffer start")
	for {
		time.Sleep(10 * time.Millisecond)
		select {
		case <-stopCh:
			log.Print("[INFO] stop readBuffer")
			return
		default:
			if r, err := p.in.(*VT100Parser).tty.ReadRune(); err == nil {
				n := utf8.EncodeRune(buf[:], r)
				bufCh <- buf[:n]
			}
		}
	}
}