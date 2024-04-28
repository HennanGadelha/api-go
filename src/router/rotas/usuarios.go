package rotas

import (
	"api-go/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{

	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		requerAutenticacao: false,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		requerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{usuariosId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		requerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{usuariosId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		requerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{usuariosId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		requerAutenticacao: false,
	},
}
