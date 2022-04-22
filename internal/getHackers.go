package internal

var leaderboardKey = "hackers"

type AllHackers struct {
	Users []*User
}

func (db *Database) GetHackers() (*AllHackers, error) {
	scores := db.Client.ZRangeWithScores(Ctx, leaderboardKey, 0, -1)
	if scores == nil {
		return nil, ErrNil
	}
	count := len(scores.Val())
	users := make([]*User, count)
	for idx, member := range scores.Val() {
		users[idx] = &User{
			Name:  member.Member.(string),
			Score: int(member.Score),
		}
	}
	leaderboard := &AllHackers{
		Users: users,
	}
	return leaderboard, nil
}
