package models

import "sync"

type District struct {
	id         int
	reps       []Cong
	sharedForm Form
}

func (d *District) Post(u User, msg Message) (res map[Cong]error) {

	wg := sync.WaitGroup{}

	for _, cong := range d.reps {
		wg.Add(1)
		go func(cong Cong) {
			err := cong.Post(u, msg)
			if err != nil {
				res[cong] = err
			}
			wg.Done()
		}(cong)
	}
	wg.Wait()
	return
}
