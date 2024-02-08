package core

type Server struct {
	host   string
	apiKey string
}

func (s *Server) SetHost(host string) {
	s.host = host
}

func (s *Server) GetHost() string {
	return s.host
}

func (s *Server) SetAPIKey(key string) {
	s.apiKey = key
}

func (s *Server) GetAPIKey() string {
	return s.apiKey
}
