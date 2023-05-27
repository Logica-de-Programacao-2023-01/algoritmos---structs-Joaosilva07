func buscarPlaylistsPorTitulo(tituloMusica string, playlists []Playlist) []Playlist {
    playlistsEncontradas := []Playlist{}
    for _, playlist := range playlists {
        for _, musica := range playlist.musicas {
            if musica.titulo == tituloMusica {
                playlistsEncontradas = append(playlistsEncontradas, playlist)
                break
            }
        }
    }
    return playlistsEncontradas
}
