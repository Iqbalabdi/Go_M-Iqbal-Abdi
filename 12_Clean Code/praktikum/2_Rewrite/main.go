package main

type Mobil struct {
	totalRoda      int
	kecepatanPerJam int
}

func (m *Mobil) berjalan() {
	kecepatanBaru := 10
	m.kecepatanPerJam = m.kecepatanPerJam + kecepatanBaru
}

func main(){
   mobilCepat := Mobil{}
   mobilCepat.berjalan()
   mobilCepat.berjalan()
   mobilCepat.berjalan()
   mobilLamban := Mobil{}
   mobilLamban.berjalan()
}

