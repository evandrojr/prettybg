package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// Configuração das flags
	style := flag.String("style", "scaled", "Estilo de exibição (none, wallpaper, centered, scaled, stretched, zoom, spanned)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Uso: %s [OPÇÕES] <caminho-da-imagem>\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Opções:")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nExemplos:")
		fmt.Fprintf(os.Stderr, "  %s ~/Imagens/wallpaper.jpg\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -style zoom ~/Imagens/wallpaper.png\n", os.Args[0])
	}
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	imagePath := flag.Arg(0)

	// Verificar se o arquivo existe
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		log.Fatalf("Arquivo não encontrado: %s", imagePath)
	}

	// Obter caminho absoluto
	absPath, err := filepath.Abs(imagePath)
	if err != nil {
		log.Fatalf("Erro ao obter caminho absoluto: %v", err)
	}

	// Verificar se é uma imagem válida (extensão)
	ext := strings.ToLower(filepath.Ext(absPath))
	switch ext {
	case ".jpg", ".jpeg", ".png", ".tif", ".tiff", ".bmp":
		// Extensão válida
	default:
		log.Fatalf("Formato de imagem não suportado: %s", ext)
	}

	// Alterar o plano de fundo usando gsettings
	cmd := exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", fmt.Sprintf("'file://%s'", absPath))
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Erro ao alterar o plano de fundo: %v\nSaída: %s", err, string(output))
	}

	// Definir o estilo de exibição
	cmd = exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-options", *style)
	output, err = cmd.CombinedOutput()
	if err != nil {
		log.Printf("Aviso: não foi possível definir o estilo de exibição: %v\nSaída: %s", err, string(output))
	}

	fmt.Printf("Plano de fundo alterado para: %s\nEstilo aplicado: %s\n", absPath, *style)
}
