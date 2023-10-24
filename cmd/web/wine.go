package main

type wine struct {
	Country string
	Name    string
	Price   float32
	Year    int16
}

func fetchWines() []wine {
	wine := []wine{
		{
			Country: "France",
			Name:    "Château Reynon Blanc Bordeaux",
			Year:    2021,
			Price:   76.95,
		},
		{
			Country: "France",
			Name:    "Madame de Beaucaillou Haut Medoc",
			Year:    2021,
			Price:   149.95,
		},
		{
			Country: "France",
			Name:    "Château Fonseche Haut-Médoc",
			Year:    2022,
			Price:   69.95,
		},
	}

	return wine
}
