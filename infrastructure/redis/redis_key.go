package redis

type CacheKey string

const (
	LoginUserName CacheKey = "login_user_name"
	SessionID     CacheKey = "session_id"
)
