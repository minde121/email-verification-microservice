var (
	ErrRedirectNotFound = errs.New("Redirect Not Found")
	ErrRedirectInvalid  = errs.New("Redirect Invalid")
)
type redirectService struct {
	redirectRepo RedisRepository
}

func NewRedisService(redirectRepo RedisRepository) RedirectService {
	return &redirectService{
		redirectRepo,
	}
}

func (r *redirectService) Find(email string) (*Redirect, error) {
	return r.redirectRepo.Find(email)
}

func (r *redirectService) Store(redirect *Redirect) error {
	redirect.Email = find(email string)
	redirect.TimeStamp = time.Now().UTC().Unix()
	return r.redirectRepo.Store(redirect)
}
