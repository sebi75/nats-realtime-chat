package auth

type AuthService struct {
	repo AuthRepositoryDB
}

func NewAuthService(repo AuthRepositoryDB) AuthService {
	return AuthService{
		repo: repo,
	}
}
