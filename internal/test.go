package internal

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

type Database struct {
	Client *redis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.TODO()
)

func NewDatabase(address string) (*Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
	if err := client.Ping(Ctx).Err(); err != nil {
		return nil, err
	}
	return &Database{
		Client: client,
	}, nil
}

var leaderboardKey = "hackers"

type Leaderboard struct {
	Users []*User
}

func (db *Database) GetLeaderboard() (*Leaderboard, error) {
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
	leaderboard := &Leaderboard{

		Users: users,
	}
	return leaderboard, nil
}
