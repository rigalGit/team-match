package game

import "time"

type PlayerReq struct {
	Id    string;
	score int;
	time  int64;
}

func CreateReq(_id string,_score int) (PlayerReq) {
	return PlayerReq{
		Id:    _id,
		score: _score,
		time:  time.Now().Unix(),
	}
}
