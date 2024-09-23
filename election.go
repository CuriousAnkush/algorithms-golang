package algorithms

import (
	"container/list"
	"reflect"
	"sort"
)

type TopVotedCandidate struct {
	persons        []int
	times          []int
	maxVotedPerson PersonVote
	timeSeries     map[int]PersonVote
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	list.New()
	v := TopVotedCandidate{
		persons:        persons,
		times:          times,
		maxVotedPerson: PersonVote{},
		timeSeries:     make(map[int]PersonVote),
	}
	v.createTimeSeries()
	return v
}

func (t *TopVotedCandidate) createTimeSeries() {
	personIdToPersonVoteMap := make(map[int]*PersonVote)
	for idx, time := range t.times {
		pId := t.persons[idx]
		if personVote, found := personIdToPersonVoteMap[pId]; !found {
			personIdToPersonVoteMap[pId] = &PersonVote{
				personId:        pId,
				voteCount:       1,
				latestTimeStamp: time,
			}
		} else {
			personVote.voteCount++
			personVote.latestTimeStamp = time
		}

		if reflect.DeepEqual(t.maxVotedPerson, PersonVote{}) {
			t.maxVotedPerson = *personIdToPersonVoteMap[pId]
		} else {
			if t.maxVotedPerson.LessThan(personIdToPersonVoteMap[pId]) {
				t.maxVotedPerson = *personIdToPersonVoteMap[pId]
			}
		}
		t.timeSeries[time] = t.maxVotedPerson
	}
}

func (this *TopVotedCandidate) Q(t int) int {
	timestampIdx := sort.Search(len(this.times), func(i int) bool {
		return this.times[i] <= t
	})

	b := timestampIdx >= len(this.times) || t != this.times[timestampIdx]
	if b {
		timestampIdx--
	}
	timestamp := this.times[timestampIdx]
	return this.timeSeries[timestamp].personId
}

type PersonVote struct {
	personId        int
	voteCount       int
	latestTimeStamp int
}

func (v PersonVote) LessThan(v2 *PersonVote) bool {
	if v.voteCount == v2.voteCount {
		return v.latestTimeStamp < v2.latestTimeStamp
	}

	return v.voteCount == v2.voteCount
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
