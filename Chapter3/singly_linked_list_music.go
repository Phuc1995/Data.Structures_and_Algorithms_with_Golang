/*package main

import "fmt"

type song struct {
	name string
	artist string
	next *song
}

type playlist struct {
	name string
	head *song
	nowPlaying *song
}

func createPlaylist(name string) *playlist  {
	return &playlist{
		name: name,
	}
}

func (p *playlist) addSong(name, artist string) error  {

	s := &song{
		name: name,
		artist: artist,
	}
	if p.head == nil{
		p.head = s
	} else {
		currentNode := p.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = s
	}

	return nil
}

func (p *playlist) showAllSong(){
	currentNode := p.head
	if currentNode == nil{
		fmt.Println("emty list")
	}
	fmt.Println( currentNode.name)

	for currentNode.next != nil{
		currentNode = currentNode.next
		fmt.Println(currentNode.name)
	}
}

func (p *playlist) startPlaying() *song{
	p.nowPlaying = p.head
	return p.nowPlaying
}

func (p *playlist) nextSong() *song  {
	p.nowPlaying = p.nowPlaying.next
	return p.nowPlaying
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
	myPlaylist.showAllSong()

	fmt.Println()
	myPlaylist.startPlaying()
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)

	myPlaylist.nextSong()
	fmt.Println("Changing next song...")
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)

	myPlaylist.nextSong()
	fmt.Println("Changing next song...")
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)

}
*/