package main

import (
	"fmt"
	"matchme/game"
	"time"
)

func main() {

	cfg := createGameConfig()

	matchSvc := game.CreateGame(cfg)

	matchSvc.AddReq(game.CreateReq("playerA", 100));
	time.Sleep(1 * time.Second);
	matchSvc.AddReq(game.CreateReq("playerB1", 10));
	time.Sleep(1 * time.Second);
	matchSvc.AddReq(game.CreateReq("playerC", 95));
	time.Sleep(1 * time.Second);
	matchSvc.AddReq(game.CreateReq("playerD", 94));
	time.Sleep(1 * time.Second);
	matchSvc.AddReq(game.CreateReq("playerE", 98));
	t1, t2, b := matchSvc.FindMatch();
	fmt.Println("Match found ", b);
	fmt.Println("Team 1");
	for _, player := range t1 {
		fmt.Println("player ", player.Id);
	}
	fmt.Println("Team 2");
	for _, player := range t2 {
		fmt.Println("player ", player.Id);
	}

}

func createGameConfig() game.GameConfig {
	// lower the priority index higher its worth
	timebucketStart := make(map[int]int)
	timebucketStart[0] = 15
	timebucketStart[1] = 10
	timebucketStart[2] = 5
	timebucketStart[3] = 0
	timebucketEnd := make(map[int]int)
	timebucketEnd[0] = 20
	timebucketEnd[1] = 15
	timebucketEnd[2] = 10
	timebucketEnd[3] = 5
	scoreConfig := make(map[int]int)
	scoreConfig[0] = 100
	scoreConfig[1] = 40
	scoreConfig[2] = 20
	scoreConfig[3] = 10
	cfg := game.GameConfig{
		GameType:         "TvsT",
		TeamSize:         1,
		TotalTeams:       2,
		TimeBucketStart:  timebucketStart,
		TimeBucketExpire: timebucketEnd,
		ScoreConfig:      scoreConfig,
		TimeBuckets:      4,
	}
	return cfg
}