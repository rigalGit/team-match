package game

type MatchSvc struct {
	allocator *AllocatorSvc;
	store *Store;
}

func CreateGame(cfg GameConfig) (*MatchSvc) {
	st := newStore(cfg.TimeBuckets,cfg);

	svc := &MatchSvc{
		allocator: create(st),
		store:st,
	}
	return svc;
}

func (svc *MatchSvc) AddReq(req PlayerReq)  {
	svc.store.addRequest(req);
	// write trigger to a channel
}

func (svc *MatchSvc) FindMatch() ([]PlayerReq, []PlayerReq, bool) {
	// to test we can manually trigger this
	// but
	return svc.allocator.findMatch();
}
