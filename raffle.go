package main

import (
	"math/rand"
	"strings"
)

type Raffle struct {
	bag []Entry
}

func (r *Raffle) getBag() []Entry {
	return r.bag
}

func (r *Raffle) removeRandom() *Entry {
	entry := &Entry{}
	entry.SetValues("INVALID", "INVALID", "INVALID")
	if len(r.bag) > 0 {
		randVal := rand.Intn(len(r.bag))
		entry = &r.bag[randVal]
		r.bag = append(r.bag[:randVal], r.bag[randVal+1:]...)
	}
	return entry
}

func (r *Raffle) addEntry(entry Entry) {
	for _, val := range r.bag {
		if strings.Compare(val.GetName(), entry.GetName()) == 0 && strings.Compare(val.GetTelephone(), entry.GetTelephone()) == 0 {
			return
		}
	}
	r.bag = append(r.bag, entry)
}

func (r *Raffle) combine(other Raffle) {
	for _, val := range other.getBag() {
		r.addEntry(val)
	}
}

func (r *Raffle) pickWinners(numberOfWinners int) {
}
