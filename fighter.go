package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Fighter struct {
	Username string
	health   uint32
}

func NewFighter(username string) *Fighter {
	return &Fighter{
		Username: username,
		health:   100,
	}
}

func (f *Fighter) Attack() uint32 {
	dmg := randomNumber(1, 20)
	return uint32(dmg)
}

func (f *Fighter) Evade() uint32 {
	evasion := randomNumber(1, 20)
	return uint32(evasion)
}

func (f *Fighter) IsDead() bool {
	return f.health <= 0
}

func (f *Fighter) DealDamage(dmg uint32) {
	if dmg > f.health {
		f.health = 0
	} else {
		f.health = f.health - dmg
	}
}

func randomNumber(min, max int) uint32 {
	return uint32(rand.Int31n(int32(max+1)-int32(min)) + int32(min))
}
