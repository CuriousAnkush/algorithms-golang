package rank_teams

import (
	"sort"
	"strings"
)

func rankTeams(votes []string) string {
	var result strings.Builder
	teamIDTeamMap := make(map[rune]Team)

	if len(votes) == 1 {
		return votes[0]
	}

	for _, vote := range votes {
		for idx, ch := range vote {
			if team, found := teamIDTeamMap[ch]; !found {
				teamIDTeamMap[ch] = Team{
					TeamID: ch,
					VoteMap: map[int]int{
						idx: 1,
					},
				}
			} else {
				team.VoteMap[idx]++
			}
		}
	}

	var teamsCo Teams

	for _, team := range teamIDTeamMap {
		teamsCo = append(teamsCo, team)
	}
	sort.Sort(teamsCo)
	for i := range teamsCo {
		result.WriteRune(teamsCo[i].TeamID)
	}

	return result.String()
}

type Team struct {
	TeamID  rune
	VoteMap map[int]int
}

type Teams []Team

func (t Teams) Len() int {
	return len(t)
}

func (t Teams) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Teams) Less(i, j int) bool {
	noOfPositions := len(t)

	for pos := 0; pos < noOfPositions; pos++ {
		if t[i].VoteMap[pos] == t[j].VoteMap[pos] {
			continue
		}
		return t[i].VoteMap[pos] > t[j].VoteMap[pos]
	}
	return t[i].TeamID < t[j].TeamID
}
