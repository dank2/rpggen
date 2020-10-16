package npc

type Npc struct {
	Name string 
}

func GenerateNpc() (*Npc, error) {
	return &Npc{
		Name: "Ashton \"Smokey\" Gint Onic av Klan Sveppes",
	}, nil
}