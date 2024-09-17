package modelos

import (
	"errors"
	"strings"
	"time"
)

// Publicacao é uma estrutura que representa uma publicação feita pelo usuário
type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autorId,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadaEm  time.Time `json:"criadaEm,omitempty"`
}

// Preparar prepara a publicação e checa se o seu conteúdo está OK para envio
func (publicacao *Publicacao) Preparar() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}
	publicacao.formatar()
	return nil

}

// falidar checa se o titulo ou conteudo estão vazios
func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("a publicação deve possuir um titulo")
	}

	if publicacao.Conteudo == "" {
		return errors.New("a publicação deve possuir um conteudo")
	}

	return nil
}

// formatar tira todos os espaços em branco desnecessários no começo e inicio do titulo e do conteudo
func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)

}
