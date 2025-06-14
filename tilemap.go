// * Pacote principal que contém a implementação do sistema de tilemap
package main

//* Importações necessárias para carregar e processar o arquivo JSON do mapa
import (
	"encoding/json" //* Pacote para decodificação do JSON
	"os"            //* Pacote para operações com arquivos
	"path"
)

// * Estrutura que representa uma camada do tilemap no formato JSON do Tiled
type TilemapLayerJSON struct {
	Data   []int  `json:"data"`   //* Array de IDs dos tiles, cada ID corresponde a um tile no tileset
	Width  int    `json:"width"`  //* Largura da camada em número de tiles
	Height int    `json:"height"` //* Altura da camada em número de tiles
	Name   string `json:"name"`   //* Nome da camada
}

// * Estrutura principal que representa o mapa completo
type TilemapJSON struct {
	Layers   []TilemapLayerJSON `json:"layers"`   //* Array de camadas do mapa
	Tilesets []map[string]any   `json:"tilesets"` //* Array de tilesets do mapa
}

func (t *TilemapJSON) GenTilesets() ([]Tileset, error) {

	tilesets := make([]Tileset, 0)

	for _, tilesetData := range t.Tilesets {
		tilesetPath := path.Join("assets/maps/", tilesetData["source"].(string))
		tileset, err := NewTileset(tilesetPath, int(tilesetData["firstgid"].(float64)))
		if err != nil {
			return nil, err
		}
		tilesets = append(tilesets, tileset)
	}

	return tilesets, nil
}

// * Função que carrega e decodifica um arquivo de mapa no formato JSON do Tiled
// * filepath: caminho para o arquivo .tmj
// * Retorna: ponteiro para a estrutura TilemapJSON e possível erro
func NewTilemapJSON(filepath string) (*TilemapJSON, error) {
	//* Lê todo o conteúdo do arquivo
	contents, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	//* Decodifica o JSON para a estrutura Go
	var tilemapJSON TilemapJSON
	err = json.Unmarshal(contents, &tilemapJSON)
	if err != nil {
		return nil, err
	}
	return &tilemapJSON, nil
}
