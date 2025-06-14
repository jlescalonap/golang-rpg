// * Pacote principal que contém a implementação da câmera
package main

//* Importação do pacote math para operações matemáticas
import "math"

// * Estrutura da câmera que mantém a posição atual no mundo do jogo
type Camera struct {
	X, Y float64 //* Coordenadas X e Y da câmera no mundo do jogo
}

// * Construtor da câmera que inicializa uma nova instância na posição especificada
func NewCamera(x, y float64) *Camera {
	return &Camera{
		X: x,
		Y: y,
	}
}

// * Método que faz a câmera seguir um alvo (geralmente o jogador)
// * targetX, targetY: posição do alvo
// * screenWidth, screenHeight: dimensões da tela para centralizar a visualização
func (c *Camera) FollowTarget(targetX, targetY, screenWidth, screenHeight float64) {
	//* Calcula o offset necessário para manter o alvo no centro da tela
	c.X = -targetX + screenWidth/2.0
	c.Y = -targetY + screenHeight/2.0
}

// * Método que restringe o movimento da câmera aos limites do mapa
// * Evita que a câmera mostre áreas fora do mapa
func (c *Camera) Constrain(tilemapWidthInPixels, tilemapHeightInPixels, screenWidthInPixels, screenHeightInPixels float64) {
	//* Impede que a câmera mostre área vazia à esquerda e acima do mapa
	c.X = math.Min(c.X, 0.0)
	c.Y = math.Min(c.Y, 0.0)

	//* Impede que a câmera mostre área vazia à direita e abaixo do mapa
	c.X = math.Max(c.X, -tilemapWidthInPixels+screenWidthInPixels)
	c.Y = math.Max(c.Y, -tilemapHeightInPixels+screenHeightInPixels)
}
