package game

import "fmt"

type AllocatorSvc struct {
	st* Store;
}

//func createMatcher(priorities int,cfg GameConfig) (*AllocatorSvc) {
//	store := newStore(priorities, cfg);
//	matcher := &AllocatorSvc{
//		st: store,
//	}
//	return matcher;
//}

func create(store *Store) (*AllocatorSvc) {
	matcher := &AllocatorSvc{
		st: store,
	}
	return matcher;
}

func (msvc*AllocatorSvc) findMatch() ([]PlayerReq,[]PlayerReq,bool){
	st := msvc.st;
	if(st.total ==0){
		fmt.Println("Can not match as total is ",0);
		return nil,nil,false;
	}
	if(st.total < st.config.TeamSize*st.config.TeamSize){
		fmt.Println("Can not match as total is less than ",st.config.TeamSize*st.config.TeamSize);
		return nil,nil,false;
	}

	team1LowestTimeBucket := st.priorityCount-1;// highest priority player in team1 , var name is lowest because lowestnumber is denoting highest priority
	team2LowestTimeBucket := st.priorityCount-1;// highest priority player in team2
	team1 := make([]PlayerReq,0,st.config.TeamSize);
	team2 := make([]PlayerReq,0,st.config.TeamSize);
	score1 := 0;
	score2 := 0;
	index := 0
	for ;len(team2) != st.config.TeamSize;{
		var req PlayerReq ;
		var f bool;
		if(len(team1) == 0){ // case when first member is being select for team1
			req,f = st.take();
			if(!f){
				return nil,nil,false;
			}
		}else { // seed(first) members has been select for both teams
			req,f = st.findClosest(score2/ len(team2))
			if(!f){
				// put back all entries removed back to queue , if can not find a suiyable player for team 1
				msvc.putBackAllToQueue(team1,team2);
				return nil,nil,false;
			}
		}
		team1 = append(team1,req);
		//team1[index] = req;
		score1 += team1[index].score;
		req,f = st.findClosest(score1/ len(team1))
		if(!f){
			//// put back all entries removed back to queue , if can not find a suiyable player for team 2
			msvc.putBackAllToQueue(team1,team2);
			return nil,nil,false;
		}
		team2 = append(team2,req);
		//team2[index] = req;
		score2 += team2[index].score;
		if(team1LowestTimeBucket > st.config.getPriority(team1[index].time)){
			team1LowestTimeBucket = st.config.getPriority(team1[index].time);
		}
		if(team2LowestTimeBucket > st.config.getPriority(team2[index].time)){
			team2LowestTimeBucket = st.config.getPriority(team2[index].time);
		}
		index++;
	}
	canPlay := st.config.canPlay(getAvg(score1, team1),team1LowestTimeBucket, getAvg(score2,team2),team2LowestTimeBucket);
	if(!canPlay){
		//putback all entries to queue
		msvc.putBackAllToQueue(team1,team2);
	}
	return team1,team2,true;
}

func (msvc*AllocatorSvc)  putBackAllToQueue(team1[] PlayerReq, team2[] PlayerReq){
	fmt.Println("put back called");
	msvc.addBackToQ(team1)
	msvc.addBackToQ(team2)
}

func (msvc *AllocatorSvc) addBackToQ(team []PlayerReq) {
	for _, v := range team {
		msvc.st.addRequest(v);
	}
}

func getAvg(score1 int, team1 []PlayerReq) int {
	return score1 / len(team1)
}

func abs(s1 int , s2 int) int  {
	if(s1 > s2){
		return s1-s2;
	}
	return s2-s1;
}
