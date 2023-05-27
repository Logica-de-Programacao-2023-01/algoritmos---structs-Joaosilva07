type Aluno struct {
    nome  string
    idade int
    notas []float64
}

func adicionarNota(aluno *Aluno, nota float64) {
    aluno.notas = append(aluno.notas, nota)
}

func removerNota(aluno *Aluno, indice int) {
    aluno.notas = append(aluno.notas[:indice], aluno.notas[indice+1:]...)
}

func calcularMediaNotas(aluno Aluno) float64 {
    total := 0.0
    for _, nota := range aluno.notas {
        total += nota
    }
    return total / float64(len(aluno.notas))
}

func imprimirInformacoesAluno(aluno Aluno) {
    fmt.Println("Nome:", aluno.nome)
    fmt.Println("Idade:", aluno.idade)
    fmt.Println("Média de Notas:", calcularMediaNotas(aluno))
}
