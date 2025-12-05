package repositories

type IsolationSegmentRepo struct {
	klient Klient
}

func NewIsolationSegmentRepo(
	klient Klient,
) *IsolationSegmentRepo {
	return &IsolationSegmentRepo{
		klient: klient,
	}
}
