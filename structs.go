package leagueliveapi

type APIError struct {
	ErrorCode             string      `json:"errorCode"`
	HTTPStatus            int         `json:"httpStatus"`
	ImplementationDetails interface{} `json:"implementationDetails"`
	Message               string      `json:"message"`
}

type AllGameData struct {
	ActivePlayer ActivePlayer `json:"activePlayer"`
	AllPlayers   []Player     `json:"allPlayers"`
	Events       Events       `json:"events"`
	GameData     GameData     `json:"gameData"`
}

type ActivePlayer struct {
	Abilities     Abilities     `json:"abilities"`
	ChampionStats ChampionStats `json:"championStats"`
	CurrentGold   float64       `json:"currentGold"`
	FullRunes     FullRunes     `json:"fullRunes"`
	Level         int           `json:"level"`
	SummonerName  string        `json:"summonerName"`
}

type Abilities struct {
	E       Ability `json:"E"`
	Passive Ability `json:"Passive"`
	Q       Ability `json:"Q"`
	R       Ability `json:"R"`
	W       Ability `json:"W"`
}

type Ability struct {
	AbilityLevel   int    `json:"abilityLevel,omitempty"`
	DisplayName    string `json:"displayName"`
	ID             string `json:"id"`
	RawDescription string `json:"rawDescription"`
	RawDisplayName string `json:"rawDisplayName"`
}

type ChampionStats struct {
	AbilityHaste                 float64 `json:"abilityHaste"`
	AbilityPower                 float64 `json:"abilityPower"`
	Armor                        float64 `json:"armor"`
	ArmorPenetrationFlat         float64 `json:"armorPenetrationFlat"`
	ArmorPenetrationPercent      float64 `json:"armorPenetrationPercent"`
	AttackDamage                 float64 `json:"attackDamage"`
	AttackRange                  float64 `json:"attackRange"`
	AttackSpeed                  float64 `json:"attackSpeed"`
	BonusArmorPenetrationPercent float64 `json:"bonusArmorPenetrationPercent"`
	BonusMagicPenetrationPercent float64 `json:"bonusMagicPenetrationPercent"`
	CritChance                   float64 `json:"critChance"`
	CritDamage                   float64 `json:"critDamage"`
	CurrentHealth                float64 `json:"currentHealth"`
	HealthRegenRate              float64 `json:"healthRegenRate"`
	LifeSteal                    float64 `json:"lifeSteal"`
	MagicLethality               float64 `json:"magicLethality"`
	MagicPenetrationFlat         float64 `json:"magicPenetrationFlat"`
	MagicPenetrationPercent      float64 `json:"magicPenetrationPercent"`
	MagicResist                  float64 `json:"magicResist"`
	MaxHealth                    float64 `json:"maxHealth"`
	MoveSpeed                    float64 `json:"moveSpeed"`
	PhysicalLethality            float64 `json:"physicalLethality"`
	ResourceMax                  float64 `json:"resourceMax"`
	ResourceRegenRate            float64 `json:"resourceRegenRate"`
	ResourceType                 string  `json:"resourceType"`
	ResourceValue                float64 `json:"resourceValue"`
	SpellVamp                    float64 `json:"spellVamp"`
	Tenacity                     float64 `json:"tenacity"`
}

type Rune struct {
	DisplayName    string `json:"displayName,omitempty"`
	ID             int    `json:"id"`
	RawDescription string `json:"rawDescription"`
	RawDisplayName string `json:"rawDisplayName,omitempty"`
}

type FullRunes struct {
	GeneralRunes      []Rune `json:"generalRunes"`
	Keystone          Rune   `json:"keystone"`
	PrimaryRuneTree   Rune   `json:"primaryRuneTree"`
	SecondaryRuneTree Rune   `json:"secondaryRuneTree"`
	StatRunes         []Rune `json:"statRunes"`
}

type Player struct {
	ChampionName    string         `json:"championName"`
	IsBot           bool           `json:"isBot"`
	IsDead          bool           `json:"isDead"`
	Items           []Item         `json:"items"`
	Level           int            `json:"level"`
	Position        string         `json:"position"`
	RawChampionName string         `json:"rawChampionName"`
	RespawnTimer    float64        `json:"respawnTimer"`
	Runes           Runes          `json:"runes"`
	Scores          Scores         `json:"scores"`
	SkinID          int            `json:"skinID"`
	SummonerName    string         `json:"summonerName"`
	SummonerSpells  SummonerSpells `json:"summonerSpells"`
	Team            string         `json:"team"`
}

type Item struct {
	CanUse         bool   `json:"canUse"`
	Consumable     bool   `json:"consumable"`
	Count          int    `json:"count"`
	DisplayName    string `json:"displayName"`
	ItemID         int    `json:"itemID"`
	Price          int    `json:"price"`
	RawDescription string `json:"rawDescription"`
	RawDisplayName string `json:"rawDisplayName"`
	Slot           int    `json:"slot"`
}

type Runes struct {
	Keystone          Rune `json:"keystone"`
	PrimaryRuneTree   Rune `json:"primaryRuneTree"`
	SecondaryRuneTree Rune `json:"secondaryRuneTree"`
}

type Scores struct {
	Assists    int     `json:"assists"`
	CreepScore int     `json:"creepScore"`
	Deaths     int     `json:"deaths"`
	Kills      int     `json:"kills"`
	WardScore  float64 `json:"wardScore"`
}

type SummonerSpell struct {
	DisplayName    string `json:"displayName"`
	RawDescription string `json:"rawDescription"`
	RawDisplayName string `json:"rawDisplayName"`
}

type SummonerSpells struct {
	SummonerSpellOne SummonerSpell `json:"summonerSpellOne"`
	SummonerSpellTwo SummonerSpell `json:"summonerSpellTwo"`
}

type Events struct {
	Events []Event `json:"Events"`
}

type Event struct {
	EventID      int      `json:"EventID"`
	EventName    string   `json:"EventName"`
	EventTime    float64  `json:"EventTime"`
	Assisters    []string `json:"Assisters,omitempty"`
	DragonType   string   `json:"DragonType,omitempty"`
	KillerName   string   `json:"KillerName,omitempty"`
	Recipient    string   `json:"Recipient,omitempty"`
	Stolen       string   `json:"Stolen,omitempty"`
	TurretKilled string   `json:"TurretKilled,omitempty"`
	VictimName   string   `json:"VictimName,omitempty"`
}

type GameData struct {
	GameMode   string  `json:"gameMode"`
	GameTime   float64 `json:"gameTime"`
	MapName    string  `json:"mapName"`
	MapNumber  int     `json:"mapNumber"`
	MapTerrain string  `json:"mapTerrain"`
}
