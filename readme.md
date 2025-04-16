# README - `prettybg`

![Wallpaper changer](https://img.icons8.com/color/96/000000/background.png)

**Um programa em Go para alterar o plano de fundo do desktop no Ubuntu Linux de forma fácil e personalizável**

## 📝 Descrição

O `prettybg` é uma ferramenta de linha de comando escrita em Go que permite alterar o plano de fundo do seu desktop no Ubuntu Linux com várias opções de estilo.

## ✨ Funcionalidades

- Altera o wallpaper com um simples comando
- Suporta múltiplos formatos de imagem (JPG, PNG, BMP, TIFF)
- Permite escolher entre diferentes estilos de exibição
- Mensagens de erro claras e úteis
- Interface simples e intuitiva

## 📦 Instalação

### Pré-requisitos

- Sistema Ubuntu Linux (com GNOME)
- Go instalado (apenas para compilação)

### Método 1: Instalar a partir do código fonte

1. Clone o repositório ou baixe o arquivo `prettybg.go`
2. Compile o programa:
   ```bash
   go build -o prettybg prettybg.go