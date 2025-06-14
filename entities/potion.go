// * Pacote entities que contém todas as entidades do jogo
package entities

// * Estrutura que representa uma poção de cura no jogo
type Potion struct {
	*Sprite      //* Incorpora a estrutura Sprite, herdando suas propriedades (Img, X, Y)
	AmtHeal uint //* Quantidade de vida que esta poção recupera quando coletada
}
