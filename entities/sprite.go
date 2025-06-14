// * Pacote entities que contém todas as entidades do jogo
package entities

//* Importação da biblioteca Ebiten para manipulação de imagens e sprites
import "github.com/hajimehoshi/ebiten/v2"

// * Estrutura base para todas as entidades que são renderizadas no jogo
// * Esta é a estrutura pai que será incorporada em outras entidades (Player, Enemy, Potion)
type Sprite struct {
	Img  *ebiten.Image //* Ponteiro para a imagem da sprite carregada na memória
	X, Y float64       //* Posição X e Y da sprite no mundo do jogo usando números decimais para movimento suave
}
