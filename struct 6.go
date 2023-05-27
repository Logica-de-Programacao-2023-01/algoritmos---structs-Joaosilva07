type Funcionario struct {
    nome    string
    salario float64
    idade   int
}

func aumentarSalario(funcionario *Funcionario, porcentagem float64) {
    funcionario.salario *= (1 + porcentagem/100)
}

func diminuirSalario(funcionario *Funcionario, porcentagem float64) {
    funcionario.salario *=
