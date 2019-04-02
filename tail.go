package logutils

import (
	"io"

	"github.com/hpcloud/tail"
)

type Tailer struct {
	tailFile *tail.Tail
	ch       chan *string
}

func NewTailer(file string) (*Tailer, error) {
	tailFile, err := tail.TailFile(
		file,
		tail.Config{
			Follow: true,
			Location: &tail.SeekInfo{
				Offset: 0,
				Whence: io.SeekStart,
			},
			ReOpen: true,
		},
	)
	if err != nil {
		return nil, err
	}

	tailer := &Tailer{
		tailFile: tailFile,
		ch:       make(chan *string, 1024),
	}

	go func() {
		//timer := time.NewTimer(1 * time.Second)
		//defer timer.Stop()
		//
		//select {
		//case <-timer.C:
		//	tailer.ch <- nil
		//}

		for line := range tailFile.Lines {
			tailer.ch <- &line.Text
		}

		tailer.ch <- nil

		return
	}()

	return tailer, nil
}

func (t *Tailer) Stop() error {
	err := t.tailFile.Stop()
	if err != nil {
		return err
	}

	return nil
}

func (t *Tailer) Tail() <-chan *string {
	return t.ch
}
