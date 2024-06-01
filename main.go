package main 

import "github.com/hegedustibor/htgo-tts"
import "github.com/hegedustibor/htgo-tts/voices"

func main() {
	speech := htgotts.Speech{
		Folder: "audio", 
		Language: voices.Indonesian,
	}
	speech.Speak("Selamat pagi, Saya Habib dari Indonesia!")

}









