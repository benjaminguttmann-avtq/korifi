package repositories

type OrgQuotaRepo struct {
	klient Klient
}

func NewOrgQuotaRepo(
	klient Klient,
) *OrgQuotaRepo {
	return &OrgQuotaRepo{
		klient: klient,
	}
}
