package modelos

// Senha representa o formato da requisição da alteração de senha feita pelo usuario
type Senha struct {
	Nova string `json:"nova"`
	Atual string `json:"atual"`
}