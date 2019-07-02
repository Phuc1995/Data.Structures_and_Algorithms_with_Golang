package main

import "fmt"

type song struct {
	name     string
	artist   string
	previous *song
	next     *song
}

type playlist struct {
	name       string
	head       *song
	tail       *song
	nowPlaying *song
}

func createPlaylist(name string) *playlist {
	return &playlist{
		name: name,
	}
}

func (p *playlist) addSong(name, author string) error {
	s := &song{
		name:   name,
		artist: author,
	}
	if p.head == nil{
		p.head = s
	}else {
		currentNode := p.tail
		currentNode.next = s
		s.previous = p.tail
	}
	p.tail = s
	return nil
}

func main()  {
	playlistName := "myplaylist"
	myPlaylist := createPlaylist(playlistName)
	fmt.Println("Created playlist")
	fmt.Println()

	fmt.Print("Adding songs to the playlist...\n\n")
	myPlaylist.addSong("Ophelia", "The Lumineers")
	myPlaylist.addSong("Shape of you", "Ed Sheeran")
	myPlaylist.addSong("Stubborn Love", "The Lumineers")
	myPlaylist.addSong("Feels", "Calvin Harris")
	fmt.Println("Showing all songs in playlist...")
}