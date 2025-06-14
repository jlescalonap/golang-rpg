// * Pacote entities que contém todas as entidades do jogo
package entities

// * Estrutura que representa o jogador no jogo
type Player struct {
	*Sprite      //* Incorpora a estrutura Sprite, herdando suas propriedades (Img, X, Y)
	Health  uint //* Pontos de vida do jogador, usando uint para garantir apenas valores positivos
}
