package config

var allowedOrigins = []string{
	"http://localhost:5173/",
	"backend-pemrog3-production.up.railway.app",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
