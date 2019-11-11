package game

import "time"

type GameConfig struct {
	GameType         string;
	TeamSize         int;
	TotalTeams       int;
	TimeBucketStart  map[int]int;
	TimeBucketExpire map[int]int;
	ScoreConfig      map[int]int;
	TimeBuckets      int;
}


func (config GameConfig) getPriority(reqTime int64) int {
	 t := time.Now().Unix() - reqTime;
	 for i,v := range config.TimeBucketStart {
	 	if( int(t) >= v){
	 		return i;
		}
	 }
	 return len(config.TimeBucketStart)-1;
}

func (config GameConfig) canPlay(score1 int, priority1 int,score2 int, priority2 int) bool {
	significantPriority := priority1;
	if(priority2 < priority1){
		significantPriority = priority2;
	}
	if( abs(score1,score2) <= config.ScoreConfig[significantPriority] ){
		return true;
	}
	return false;

}

func (config GameConfig) getExpiryTime(reqTime int64) int64 {
	p := config.getPriority(reqTime);
	v,_ := config.TimeBucketExpire[p];
	return reqTime + int64(v);
}


