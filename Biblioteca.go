package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// Estrutura do livro
type Livro struct {
	ID     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
	Ano    int    `json:"ano"`
}

// Estrutura para representar um periódico ou revista
type Periodico struct {
	Nome   string `json:"nome"`
	Tipo   string `json:"tipo"`
	Volume int    `json:"volume"`
}

// Slice para armazenar os livros
var Livros []Livro

// Slice para armazenar os periódicos
var Periodicos []Periodico

// Mutex para proteger os slices em concorrência
var mutex sync.Mutex

// Inicializa os dados a partir do arquivo
func InicializarDados(arquivo string) error {
	file, err := os.Open(arquivo)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Split(line, "|")
		if len(parts) < 4 {
			continue
		}

		id, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}

		ano, err := strconv.Atoi(parts[3])
		if err != nil {
			continue
		}

		novoLivro := Livro{
			ID:     id,
			Titulo: parts[1],
			Autor:  parts[2],
			Ano:    ano,
		}
		Livros = append(Livros, novoLivro)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// Função para listar os livros carregados do arquivo
func ListarLivros() {
	fmt.Println("Lista de Livros:")
	for _, livro := range Livros {
		fmt.Printf("ID: %d, Título: %s, Autor: %s, Ano: %d\n", livro.ID, livro.Titulo, livro.Autor, livro.Ano)
	}
	fmt.Println()
}

// Função para buscar um livro pelo título
func BuscarLivroPorTitulo(titulo string) {
	encontrado := false
	for _, livro := range Livros {
		if strings.Contains(strings.ToLower(livro.Titulo), strings.ToLower(titulo)) {
			fmt.Printf("Livro encontrado: ID: %d, Título: %s, Autor: %s, Ano: %d\n", livro.ID, livro.Titulo, livro.Autor, livro.Ano)
			encontrado = true
		}
	}
	if !encontrado {
		fmt.Println("Nenhum livro encontrado com esse título.")
	}
	fmt.Println()
}

// Função para buscar livros por autor
func BuscarLivrosPorAutor(autor string) {
	encontrados := false
	for _, livro := range Livros {
		if strings.Contains(strings.ToLower(livro.Autor), strings.ToLower(autor)) {
			fmt.Printf("ID: %d, Título: %s, Autor: %s, Ano: %d\n", livro.ID, livro.Titulo, livro.Autor, livro.Ano)
			encontrados = true
		}
	}
	if !encontrados {
		fmt.Println("Nenhum livro encontrado desse autor.")
	}
	fmt.Println()
}

// Função para remover um livro
func RemoverLivro(id int) {
	mutex.Lock()
	defer mutex.Unlock()

	indice := -1
	for i, livro := range Livros {
		if livro.ID == id {
			indice = i
			break
		}
	}

	if indice == -1 {
		fmt.Println("Livro não encontrado.")
	} else {
		Livros = append(Livros[:indice], Livros[indice+1:]...)
		fmt.Println("Livro removido com sucesso.")
	}
	fmt.Println()
}

// Função para listar os periódicos
func ListarPeriodicos() {
	fmt.Println("Lista de Periódicos:")
	for _, peri := range Periodicos {
		fmt.Printf("Nome: %s, Tipo: %s, Volume: %d\n", peri.Nome, peri.Tipo, peri.Volume)
	}
	fmt.Println()
}

// Função para buscar periódicos por nome
func BuscarPeriodicoPorNome(nome string) {
	encontrado := false
	for _, peri := range Periodicos {
		if strings.Contains(strings.ToLower(peri.Nome), strings.ToLower(nome)) {
			fmt.Printf("Periódico encontrado: Nome: %s, Tipo: %s, Volume: %d\n", peri.Nome, peri.Tipo, peri.Volume)
			encontrado = true
		}
	}
	if !encontrado {
		fmt.Println("Nenhum periódico encontrado com esse nome.")
	}
	fmt.Println()
}

// Função para adicionar um livro
func AdicionarLivro(titulo, autor string, ano int) {
	mutex.Lock()
	defer mutex.Unlock()

	novoLivro := Livro{
		ID:     len(Livros) + 1,
		Titulo: titulo,
		Autor:  autor,
		Ano:    ano,
	}
	Livros = append(Livros, novoLivro)

	fmt.Println("Novo livro adicionado com sucesso!")
	fmt.Println()
}
func AdicionarLivroNoArquivo(livro Livro) error {
	file, err := os.OpenFile("F:trabalho/arquivo.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%d|%s|%s|%d\n", livro.ID, livro.Titulo, livro.Autor, livro.Ano)
	if err != nil {
		return err
	}

	return nil
}
// Função para adicionar um periódico
func AdicionarPeriodico(nome, tipo string, volume int) {
	mutex.Lock()
	defer mutex.Unlock()

	novoPeriodico := Periodico{
		Nome:   nome,
		Tipo:   tipo,
		Volume: volume,
	}
	Periodicos = append(Periodicos, novoPeriodico)

	fmt.Println("Novo periódico adicionado com sucesso!")
	fmt.Println()
}

// Função principal
func main() {
	if err := InicializarDados("C:/Users/lucas/OneDrive/bibliotGit/ler/arquivo.txt"); err != nil {
		fmt.Println("Erro ao inicializar dados:", err)
		return
	}

	// Menu de opções
	fmt.Println("Bem-vindo à Biblioteca Virtual")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Opções:")
		fmt.Println("1. Listar Livros")
		fmt.Println("2. Buscar Livro por Título")
		fmt.Println("3. Buscar Livros por Autor")
		fmt.Println("4. Remover Livro")
		fmt.Println("5. Listar Periódicos")
		fmt.Println("6. Buscar Periódico por Nome")
		fmt.Println("7. Adicionar Livro")
		fmt.Println("8. Adicionar Periódico")
		fmt.Println("0. Sair")

		fmt.Print("Escolha uma opção: ")
		opcaoStr, _ := reader.ReadString('\n')
		opcaoStr = strings.TrimSpace(opcaoStr)
		opcao, err := strconv.Atoi(opcaoStr)
		if err != nil {
			fmt.Println("Opção inválida. Tente novamente.")
			continue
		}

		switch opcao {
		case 1:
			ListarLivros()
		case 2:
			fmt.Print("Digite o título do livro: ")
			titulo, _ := reader.ReadString('\n')
			titulo = strings.TrimSpace(titulo)
			BuscarLivroPorTitulo(titulo)
		case 3:
			fmt.Print("Digite o nome do autor: ")
			autor, _ := reader.ReadString('\n')
			autor = strings.TrimSpace(autor)
			BuscarLivrosPorAutor(autor)
		case 4:
			fmt.Print("Digite o ID do livro a ser removido: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, _ := strconv.Atoi(idStr)
			RemoverLivro(id)
		case 5:
			ListarPeriodicos()
		case 6:
			fmt.Print("Digite o nome do periódico: ")
			nome, _ := reader.ReadString('\n')
			nome = strings.TrimSpace(nome)
			BuscarPeriodicoPorNome(nome)
		case 7:
			fmt.Print("Título do livro: ")
			titulo, _ := reader.ReadString('\n')
			titulo = strings.TrimSpace(titulo)

			fmt.Print("Autor do livro: ")
			autor, _ := reader.ReadString('\n')
			autor = strings.TrimSpace(autor)

			fmt.Print("Ano do livro: ")
			anoStr, _ := reader.ReadString('\n')
			ano, _ := strconv.Atoi(strings.TrimSpace(anoStr))

			AdicionarLivro(titulo, autor, ano)
		case 8:
			fmt.Print("Nome do periódico: ")
			nome, _ := reader.ReadString('\n')
			nome = strings.TrimSpace(nome)

			fmt.Print("Tipo do periódico: ")
			tipo, _ := reader.ReadString('\n')
			tipo = strings.TrimSpace(tipo)

			fmt.Print("Volume do periódico: ")
			volumeStr, _ := reader.ReadString('\n')
			volume, _ := strconv.Atoi(strings.TrimSpace(volumeStr))

			AdicionarPeriodico(nome, tipo, volume)
		case 0:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}
