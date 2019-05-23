

// func Error(w ResponseWriter, error string, code int)
// json := `{"status":"error", "msg":"method not supported, only POST"}`
// http.Error(w, json, http.StatusUnauthorized)

const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH" // RFC 5789
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)

const (
	StatusContinue           = 100 // RFC 7231, 6.2.1
	StatusSwitchingProtocols = 101 // RFC 7231, 6.2.2
	StatusProcessing         = 102 // RFC 2518, 10.1

	StatusOK       = 200 // RFC 7231, 6.3.1
	StatusCreated  = 201 // RFC 7231, 6.3.2
	StatusAccepted = 202 // RFC 7231, 6.3.3

)

// const DefaultMaxHeaderBytes = 1 << 20 // 1 MB

// O ServeMux é um multiplexador de solicitação HTTP.
// Ele corresponde a URL de cada solicitação recebida a uma lista de padrões
// registrados e chama o manipulador para o padrão que mais se aproxima do URL.

// O ServeMux também cuida da limpeza do caminho de solicitação de URL e do cabeçalho do Host,
// retirando o número da porta e redirecionando qualquer solicitação que contenha

type ServeMux struct {
	// contains filtered or unexported fields
}

func NewServeMux() *ServeMux
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
