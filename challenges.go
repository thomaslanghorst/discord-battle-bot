package main

import (
	"errors"
	"fmt"
)

type Challenges struct {
	openChallenges map[string][]string
}

func NewChallenges() *Challenges {
	return &Challenges{
		openChallenges: make(map[string][]string),
	}
}

func (c *Challenges) AddChallengers(user1, user2 string) error {
	user1Challenges := c.openChallenges[user1]
	user2Challenges := c.openChallenges[user2]

	if user1Challenges == nil {
		c.openChallenges[user1] = make([]string, 0)
	}

	if user2Challenges == nil {
		c.openChallenges[user2] = make([]string, 0)
	}

	for _, challenger := range user1Challenges {
		if challenger == user2 {
			return errors.New("challenge already exists")
		}
	}

	for _, challenger := range user2Challenges {
		if challenger == user1 {
			return errors.New("challenge already exists")
		}
	}

	c.openChallenges[user1] = append(c.openChallenges[user1], user2)
	c.openChallenges[user2] = append(c.openChallenges[user2], user1)

	return nil

}

func (c *Challenges) RemoveChallengers(user1, user2 string) error {

	user1Challenges := c.openChallenges[user1]
	user2Challenges := c.openChallenges[user2]

	if user1Challenges == nil {
		return errors.New(fmt.Sprintf("no challenge between %s and %s exists", user1, user2))
	}

	if user2Challenges == nil {
		return errors.New(fmt.Sprintf("no challenge between %s and %s exists", user2, user1))
	}

	idxUser1 := 0
	idxUser2 := 0

	for idx1, challenger := range user1Challenges {
		if challenger == user2 {
			idxUser1 = idx1
			break
		}
	}

	for idx2, challenger := range user2Challenges {
		if challenger == user1 {
			idxUser2 = idx2
			break
		}
	}

	c.openChallenges[user1] = remove(c.openChallenges[user1], idxUser1)
	c.openChallenges[user2] = remove(c.openChallenges[user2], idxUser2)

	return nil
}

func remove(s []string, idx int) []string {
	s[idx] = s[len(s)-1]
	return s[:len(s)-1]
}

func (c *Challenges) ChallengeExists(user1, user2 string) bool {

	// we need to only check once since AddChallenge adds into:
	// c.openChallenges[user1] as well as
	// c.openChallenges[user2]
	if _, exists := c.openChallenges[user1]; exists {
		return exists
	}

	return false
}

func (c *Challenges) OpenChallengers(user string) []string {
	challenges := c.openChallenges[user]
	if challenges == nil {
		return []string{}
	} else {
		return challenges
	}
}
