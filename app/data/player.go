package data

type Player struct {
	ID int64 `json:"id"`
	UserID int64 `json:"user_id"`
	VipPlayer bool `json:"vip_player"`
	Username string `json:"username"`
	Avatar string `json:"avatar"`
	Level int64 `json:"level"`
	Experience int64 `json:"experience"`
	ExperienceForNextLevel int64 `json:"experience_for_next_level"`
	Stats *BaseStats // See BaseStats
	Assets *Assets // See Owned
	Exploration *Exploration // See Exploration
}
type BaseStats struct {
	Health int64 `json:"health"`
	MaxHealth int64 `json:"max_health"`
	Strength int64 `json:"strength"`
	CriticalStr int64 `json:"critical"`
	Defense int64 `json:"defense"`
	Agility int64 `json:"agility"`
	Accuracy int64 `json:"accuracy"`
	Dodge int64 `json:"dodge"`
	StatModifiers *StatModifiers // See StatModifiers
}
type StatModifiers struct {
	Buffs []int `json:"buffs"`
	Debuffs []int `json:"debuffs"`
}
type Assets struct {
	VipUnits int64 `json:"vip_units"`
	Units int64 `json:"units"`
	Debt int64 `json:"copper"`
	Inventory *Inventory // See Inventory
}
type Inventory struct {
	Weapons []int `json:"weapons"`
	Armor []int `json:"armor"`
	Junk []int `json:"junk"`
}
type Exploration struct {
	CurrentLocation int64 `json:"current_location"`
	CurrentLocationName string `json:"current_location_name"`
	CurrentLocationDescription string `json:"current_location_description"`
	CurrentLocationImage string `json:"current_location_image"`
	StepsToNextLocation int64 `json:"steps_to_next_location"`
	Steps int64 `json:"steps"`
	StepsMax int64 `json:"steps_max"`
	StepModifier float64 `json:"step_modifier"`
}

func (U *User) GetPlayer(user_id uint64) (*Player, error) {
	statement := "SELECT * FROM players WHERE user_id=$1"
	row := DB.QueryRow(statement, user_id)
	var player *Player
	if row == nil {
		newPlayer, err := createNewPlayer(user_id)
		if err != nil {
			return &Player{}, err
		}
		return newPlayer, nil
	}
	err := row.Scan(&player.ID, &player.UserID, &player.VipPlayer, &player.Username, &player.Avatar, &player.Level, &player.Experience, &player.ExperienceForNextLevel)
	if err != nil {
		return &Player{}, err
	}
	playerStats := "SELECT * FROM player_stats WHERE player_id=$1"
	statsRow := DB.QueryRow(playerStats, player.ID)
	if statsRow == nil {
		return &Player{}, err
	}
	err = statsRow.Scan(&player.Stats.Health, &player.Stats.MaxHealth, &player.Stats.Strength, &player.Stats.CriticalStr, &player.Stats.Defense, &player.Stats.Agility, &player.Stats.Accuracy, &player.Stats.Dodge, &player.Stats.StatModifiers.Buffs, &player.Stats.StatModifiers.Debuffs)
	if err != nil {
		return &Player{}, err
	}
	playerAssets := "SELECT * FROM player_assets WHERE player_id=$1"
	assetsRow := DB.QueryRow(playerAssets, player.ID)
	if assetsRow == nil {
		return &Player{}, err
	}
	err = assetsRow.Scan(&player.Assets.VipUnits, &player.Assets.Units, &player.Assets.Debt, &player.Assets.Inventory.Weapons, &player.Assets.Inventory.Armor, &player.Assets.Inventory.Junk)
	if err != nil {
		return &Player{}, err
	}
	playerExploration := "SELECT * FROM player_exploration WHERE player_id=$1"
	explorationRow := DB.QueryRow(playerExploration, player.ID)
	if explorationRow == nil {
		return &Player{}, err
	}
	err = explorationRow.Scan(&player.Exploration.CurrentLocation, &player.Exploration.CurrentLocationName, &player.Exploration.CurrentLocationDescription, &player.Exploration.CurrentLocationImage, &player.Exploration.StepsToNextLocation, &player.Exploration.Steps, &player.Exploration.StepsMax, &player.Exploration.StepModifier)
	if err != nil {
		return &Player{}, err
	}
	return player, nil
}

var newPlayerData = Player {
	UserID: 0,
	VipPlayer: false,
	Username: "",
	Avatar: "",
	Level: 1,
	Experience: 0,
	ExperienceForNextLevel: 100,
	Stats: &BaseStats {
		Health: 100,
		MaxHealth: 100,
		Strength: 5,
		CriticalStr: 8,
		Defense: 5,
		Agility: 5,
		Accuracy: 5,
		Dodge: 5,
		StatModifiers: &StatModifiers{
			Buffs: nil,
			Debuffs: nil,
		},
	},
	Assets: &Assets{
		VipUnits: 25,
		Units: 1000,
		Debt: 0,
		Inventory: &Inventory{
			Weapons: nil,
			Armor: nil,
			Junk: nil,
		},
	},
	Exploration: &Exploration {
		// TODO: Get exploration data from database
		CurrentLocation: 0,
		CurrentLocationName: "Dungeon Entrance",
		CurrentLocationDescription: "You arrive at a collosal entrance, easily 4 times your height and wide enough to accomdate  to something... This large crystalline structured archway houses what appears to be a large weathered wooden door. its vast ancient bodywork, is littered with oddly shaped recesses, as though at one time, maybe artifacts or a collection of strange tools had once adorned these voids. The door sits slightly ajar and you hear it rocking and creaking on its corroded hinges as the heavy winter wind batters it. Faint shuffling sounds echo beyond the door...",
		CurrentLocationImage: "/assets/images/locations/0.png",
		StepsToNextLocation: 1250,
		Steps: 250,
		StepsMax: 250,
		StepModifier: 1.0,
	},
}
func createNewPlayer(id uint64) (*Player, error) {

	playerCreate := "INSERT INTO players (user_id, vip_player, username, avatar, level, experience, experience_for_next_level) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	newPlayerRow, err := DB.Exec(playerCreate, &newPlayerData.UserID, &newPlayerData.VipPlayer, &newPlayerData.Username, &newPlayerData.Avatar, &newPlayerData.Level, &newPlayerData.Experience, &newPlayerData.ExperienceForNextLevel)
	if err != nil {
		
		return &Player{}, err
	}
	newPlayerID, err := newPlayerRow.LastInsertId()
	if err != nil {
		return &Player{}, err
	}
	playerStats := "INSERT INTO player_stats (player_id, health, max_health, strength, criticalstr, defense, agility, accuracy, dodge, buffs, debuffs) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"

	_, err = DB.Exec(playerStats, &newPlayerID, &newPlayerData.Stats.Health, &newPlayerData.Stats.MaxHealth, &newPlayerData.Stats.Strength, &newPlayerData.Stats.CriticalStr, &newPlayerData.Stats.Defense, &newPlayerData.Stats.Agility, &newPlayerData.Stats.Accuracy, &newPlayerData.Stats.Dodge, &newPlayerData.Stats.StatModifiers.Buffs, &newPlayerData.Stats.StatModifiers.Debuffs)
	if err != nil {
		return &Player{}, err
	}
	playerAssets := "INSERT INTO player_assets (player_id, vip_units, units, debt, weapons, armor, junk) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	_, err = DB.Exec(playerAssets, &newPlayerID, &newPlayerData.Assets.VipUnits, &newPlayerData.Assets.Units, &newPlayerData.Assets.Debt, &newPlayerData.Assets.Inventory.Weapons, &newPlayerData.Assets.Inventory.Armor, &newPlayerData.Assets.Inventory.Junk)
	if err != nil {
		return &Player{}, err
	}

	playerExploration := "INSERT INTO player_exploration (player_id, current_location, location_name, location_description, location_image, steps_next_location, steps, steps_max, step_modifier) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	_, err = DB.Exec(playerExploration, &newPlayerID, &newPlayerData.Exploration.CurrentLocation, &newPlayerData.Exploration.CurrentLocationName, &newPlayerData.Exploration.CurrentLocationDescription, &newPlayerData.Exploration.CurrentLocationImage, &newPlayerData.Exploration.StepsToNextLocation, &newPlayerData.Exploration.Steps, &newPlayerData.Exploration.StepsMax, &newPlayerData.Exploration.StepModifier)
	if err != nil {
		return &Player{}, err
	}
	var newPlayer = &Player{
		ID: newPlayerID,
		UserID: newPlayerData.UserID,
		VipPlayer: newPlayerData.VipPlayer,
		Username: newPlayerData.Username,
		Avatar: newPlayerData.Avatar,
		Level: newPlayerData.Level,
		Experience: newPlayerData.Experience,
		ExperienceForNextLevel: newPlayerData.ExperienceForNextLevel,
		Stats: &BaseStats{
			Health: newPlayerData.Stats.Health,
			MaxHealth: newPlayerData.Stats.MaxHealth,
			Strength: newPlayerData.Stats.Strength,
			CriticalStr: newPlayerData.Stats.CriticalStr,
			Defense: newPlayerData.Stats.Defense,
			Agility: newPlayerData.Stats.Agility,
			Accuracy: newPlayerData.Stats.Accuracy,
			Dodge: newPlayerData.Stats.Dodge,
			StatModifiers: &StatModifiers{
				Buffs: newPlayerData.Stats.StatModifiers.Buffs,
				Debuffs: newPlayerData.Stats.StatModifiers.Debuffs,
			},
		},
		Assets: &Assets{
			VipUnits: newPlayerData.Assets.VipUnits,
			Units: newPlayerData.Assets.Units,
			Debt: newPlayerData.Assets.Debt,
			Inventory: &Inventory{
				Weapons: newPlayerData.Assets.Inventory.Weapons,
				Armor: newPlayerData.Assets.Inventory.Armor,
				Junk: newPlayerData.Assets.Inventory.Junk,
			},
		},
		Exploration: &Exploration{
			CurrentLocation: newPlayerData.Exploration.CurrentLocation,
			CurrentLocationName: newPlayerData.Exploration.CurrentLocationName,
			CurrentLocationDescription: newPlayerData.Exploration.CurrentLocationDescription,
			CurrentLocationImage: newPlayerData.Exploration.CurrentLocationImage,
			StepsToNextLocation: newPlayerData.Exploration.StepsToNextLocation,
			Steps: newPlayerData.Exploration.Steps,
			StepsMax: newPlayerData.Exploration.StepsMax,
			StepModifier: newPlayerData.Exploration.StepModifier,
		},
	}
	return newPlayer, nil
}