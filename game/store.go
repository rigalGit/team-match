package game

type IRequestStore interface {
	addRequest(req PlayerReq);
	take() (PlayerReq,bool);
	findClosest(score int) (PlayerReq,bool) ;
}