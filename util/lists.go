package util

type R2List struct {
	Name         string
	ReadableName string
}

var (
	Maps = []R2List{
		//mp maps
		{"mp_angel_city", "Angel City"},
		{"mp_black_water_canal", "Black Water Canal"},
		{"mp_grave", "Boomtown"},
		{"mp_colony02", "Colony"},
		{"mp_complex3", "Complex"},
		{"mp_crashsite3", "Crashsite"},
		{"mp_drydock", "DryDock"},
		{"mp_eden", "Eden"},
		{"mp_thaw", "Exoplanet"},
		{"mp_forwardbase_kodai", "Forward Base Kodai"},
		{"mp_glitch", "Glitch"},
		{"mp_homestead", "Homestead"},
		{"mp_relic02", "Relic"},
		{"mp_rise", "Rise"},
		{"mp_wargames", "Wargames"},
		{"mp_lobby", "Lobby"},
		{"mp_box", "Box"},
		//lf maps
		{"mp_lf_deck", "Deck"},
		{"mp_lf_meadow", "Meadow"},
		{"mp_lf_stacks", "Stacks"},
		{"mp_lf_township", "Township"},
		{"mp_lf_traffic", "Traffic"},
		{"mp_lf_uma", "UMA"},
		//coliseum
		{"mp_coliseum", "The Coliseum"},
		{"mp_coliseum_column", "Pillars"},
		//campaign
		{"sp_training", "The Pilot's Gauntlet"},
		{"sp_crashsite", "BT-7274"},
		{"sp_sewers1", "Blood and Rust"},
		{"sp_boomtown_start", "Into the Abyss - Part 1"},
		{"sp_boomtown", "Into the Abyss - Part 2"},
		{"sp_boomtown_end", "Into the Abyss - Part 2"},
		{"sp_hub_timeshift", "Effect and Cause - Part 1 or 3"},
		{"sp_timeshift_spoke02", "Effect and Cause - Part 2"},
		{"sp_beacon", "The Beacon - Part 1 or 3"},
		{"sp_beacon_spoke0", "The Beacon - Part 2"},
		{"sp_tday", "Trial by Fire"},
		{"sp_s2s", "The Ark"},
		{"sp_skyway_v1", "The Fold Weapon"},
	}

	Playlists = []R2List{
		{"private_match", "Private Match"},
		//vanilla
		{"aitdm", "Attrition"},
		{"at", "Bounty Hunt"},
		{"coliseum", "Coliseum"},
		{"cp", "Amped Hardpoint"},
		{"ctf", "Capture the Flag"},
		{"fd", "Frontier Defense"},
		{"fd_easy", "Frontier Defense {Easy}"},
		{"fd_normal", "Frontier Defense {Regular}"},
		{"fd_hard", "Frontier Defense {Hard}"},
		{"fd_insane", "Frontier Defense {Insane}"},
		{"fd_master", "Frontier Defense {Master}"},
		{"lts", "Last Titan Standing"},
		{"mfd", "Marked For Death"},
		{"ps", "Pilots vs. Pilots"},
		{"solo", "Campaign"},
		{"tdm", "Skirmish"},
		{"ttdm", "Titan Brawl"},
		{"lf", "Live Fire"},
		//vanilla featured
		{"alts", "Aegis Last Titan Standing"},
		{"attdm", "Aegis Titan Brawl"},
		{"ffa", "Free For All"},
		{"fra", "Free Agents"},
		{"holopilot_lf", "The Great Bamboozle"},
		{"rocket_lf", "Rocket Arena"},
		{"turbo_lts", "Turbo Last Titan Standing"},
		{"turbo_ttdm", "Turbo Titan Brawl"},
		//northstar custom
		{"chamber", "One in the Chamber"},
		{"ctf_comp", "Competitive CTF"},
		{"fastball", "Fastball"},
		{"gg", "Gun Game"},
		{"hidden", "The Hidden"},
		{"hs", "Hide and Seek"},
		{"inf", "Infection"},
		{"kr", "Amped Killrace"},
		{"sns", "Sticks and Stones"},
		{"tffa", "Titan FFA"},
		{"tt", "Titan Tag"},
		{"fw", "Frontier War"},
		//northstar coop
		{"sp_coop", "Singleplayer Coop"},
	}
)
