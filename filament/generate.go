package filament

import "fmt"

func GenerateFilaments() ([]Filament, error) {
	profiles := []struct {
		name        string
		inherits    string
		filamentTyp string
	}{
		{name: "Gen - PLA", inherits: "Generic PLA @BBL X1C", filamentTyp: "PLA"},
		{name: "Gen - PETG", inherits: "Generic PETG @BBL X1C", filamentTyp: "PETG"},
		{name: "Gen - ABS", inherits: "Generic ABS @BBL X1C", filamentTyp: "ABS"},
	}

	filaments := make([]Filament, 0, len(profiles))

	for _, p := range profiles {
		f := Filament{
			Name:           p.name,
			From:           "User",
			Inherits:       p.inherits,
			FilamentType:   p.filamentTyp,
			FilamentVendor: []string{"Generic"},
			FilamentId:     p.name,
			SettingId:      p.name,
			Description:    "Generated default filament profile",
			InfoFile:       fmt.Sprintf("sync_info = update\nuser_id = \nsetting_id = \nbase_id = GF001\nupdated_time = 1703950786\n"),
		}

		filaments = append(filaments, f)
	}

	return filaments, nil
}
