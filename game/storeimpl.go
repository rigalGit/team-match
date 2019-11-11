package game

import (
	"fmt"
	"github.com/emirpasic/gods/maps/treemap"
)

type Store struct {
	timeValues[] PriorityQueue; // list of time bucket queues storing requests based on time
	scoreValues[] *treemap.Map; // list of time bucket balanced bst maps stroing entries based on score
	priorityCount int;
	config GameConfig;
	total int;
}

func newStore(priorities int,cfg GameConfig) (*Store){
	s := &Store{
		timeValues:  make([]PriorityQueue,priorities,priorities),
		scoreValues: make([]*treemap.Map,priorities,priorities),
		priorityCount : priorities,
		config : cfg,
		total:0,
	}
	for i:=0;i<priorities;i++{
		s.scoreValues[i] = treemap.NewWithIntComparator();
	}
	return s;
}
// add new requests to store
func (st *Store) addRequest(req PlayerReq)  {
	id := req.Id;
	score := req.score;

	item := &Item{
		value:id,
		priority: int(req.time),
		expiryTime: st.config.getExpiryTime(req.time),
		joinReq: req,
	}
	priorityList := st.config.getPriority(req.time)
	st.timeValues[priorityList].Push(item);
	m := st.scoreValues[priorityList];
	fmt.Println("size ",m.Size());
	var pMap map[string]string;
	v,ok := m.Get(score);
	if(!ok){
		pMap = make(map[string]string);
		m.Put(score,pMap);
	}else {
		pMap = v.(map[string]string);
	}
	pMap[id] = id;
	st.total++;

}
// take out most suitable candidate for match from store across all buckets
func (st *Store) take() (PlayerReq,bool) {
	var req PlayerReq;
	if(st.total == 0){
		return req,false;
	}
	for i := 0; i<st.priorityCount; i++ {
		if(st.timeValues[i].Len() == 0){
			continue;
		}
		pop := st.timeValues[i].Pop();
		fmt.Println("op",pop);
		topReq := pop.(*Item)
		return topReq.joinReq,true;
	}
	return req,false;
}
// find most suitable candidate against given score
func (st *Store) findClosest(score int) (PlayerReq,bool) {
	var req PlayerReq;
	if(st.total == 0){
		return req,false;
	}
	reqList := make([]PlayerReq,0,st.priorityCount);
	for i := 0; i<st.priorityCount; i++ {
		if(st.timeValues[i].Len() == 0){
			continue;
		}
		topReq := st.timeValues[i].Pop().(*Item)
		reqList = append(reqList,topReq.joinReq)
	}
	if(len(reqList) == 0){
		return req,false;
	}
	var bestMatch PlayerReq = reqList[0];
	bestAbs := abs(bestMatch.score,score);
	for _,r := range reqList{
		if(abs(score,r.score) < bestAbs){
			bestAbs = abs(score,r.score);
			bestMatch = r;
		}
	}
	for _,r := range reqList{
		if(r.Id != bestMatch.Id){
			st.addRequest(r);
		}
	}
	return bestMatch,true;
}



