package config

var allowedOrigins = []string{
	"http://localhost:5173/",
	"https://front-end-pemprog3.vercel.app/",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
