// * Pacote entities que contém todas as entidades do jogo
package entities

// * Estrutura que representa um inimigo no jogo
type Enemy struct {
	*Sprite            //* Incorpora a estrutura Sprite, herdando suas propriedades (Img, X, Y)
	FollowsPlayer bool //* Flag que determina se este inimigo persegue o jogador (IA básica)
}
