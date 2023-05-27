type Endereco struct {
    rua     string
    numero  int
    cidade  string
    estado  string
}

type Pessoa struct {
    nome     string
    idade    int
    endereco Endereco
}

func imprimirEnderecoCompleto(pessoa Pessoa) {
    endereco := pessoa.endereco
    fmt.Printf("Rua: %s, Número: %d, Cidade: %s, Estado: %s\n", endereco.rua, endereco.numero, endereco.cidade, endereco.estado)
}
