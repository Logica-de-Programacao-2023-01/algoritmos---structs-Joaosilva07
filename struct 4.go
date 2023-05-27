type Musica struct {
    titulo   string
    artista  string
    duracao  string
}

type Playlist struct {
    nome    string
    musicas []Musica
}

func imprimirInformacoesPlaylist(playlist Playlist) {
    fmt.Println("Nome da playlist:", playlist.nome)
    tempoTotal := 0
    for _, musica := range playlist.musicas {
        fmt.Printf("Música: %s, Artista: %s, Duração: %s\n", musica.titulo, musica.artista, musica.duracao)
        tempoTotal += converterDuracaoParaSegundos(musica.duracao)
    }
    fmt.Println("Tempo total da playlist:", converterSegundosParaDuracao(tempoTotal))
}

func converterDuracaoParaSegundos(duracao string) int {
    
}

func converterSegundosParaDuracao(segundos int) string {
}
