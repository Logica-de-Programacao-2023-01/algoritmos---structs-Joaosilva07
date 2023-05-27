type Viagem struct {
    origem  string
    destino string
    data    string
    preço   float64
}

func encontrarViagemMaisCara(viagens []Viagem) Viagem {
    viagemMaisCara := viagens[0]
    for _, viagem := range viagens {
        if viagem.preço > viagemMaisCara.preço {
            viagemMaisCara = viagem
        }
    }
    return viagemMaisCara
}
