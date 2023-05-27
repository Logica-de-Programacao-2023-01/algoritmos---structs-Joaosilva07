type Filme struct {
    título     string
    diretor    string
    ano        int
    avaliações []int
}

func adicionarAvaliação(filme *Filme, avaliação int) {
    filme.avaliações = append(filme.avaliações, avaliação)
}

func removerAvaliação(filme *Filme, indice int) {
    filme.avaliações = append(filme.avaliações[:indice], filme.avaliações[indice+1:]...)
}

func calcularMédiaAvaliações(filme Filme) float64 {
    total := 0
    for _, avaliação := range filme.avaliações {
        total += avaliação
    }
    return float64(total) / float64(len(filme.avaliações))
}

func imprimirInformaçõesFilme(filme Filme) {
    fmt.Println("Título:", filme.título)
    fmt.Println("Diretor:", filme.diretor)
    fmt.Println("Ano:", filme.ano)
    fmt.Println("Média de Avaliações:", calcularMédiaAvaliações(filme))
}
