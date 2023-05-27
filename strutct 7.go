type Animal struct {
    nome   string
    espécie string
    idade  int
    som    string
}

func modificarSom(animal *Animal, novoSom string) {
    animal.som = novoSom
}

func imprimirInformacoesAnimal(animal Animal) {
    fmt.Println("Nome:", animal.nome)
    fmt.Println("Espécie:", animal.espécie)
    fmt.Println("Idade:", animal.idade)
    fmt.Println("Som:", animal.som)
}
