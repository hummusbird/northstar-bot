package commands

import "github.com/bwmarrin/discordgo"

var (
	ListOpts = []*discordgo.ApplicationCommandOption{
		{
			Name:        "type",
			Description: "List to send",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
			Choices: []*discordgo.ApplicationCommandOptionChoice{
				{
					Name:  "Modes",
					Value: "modes",
				},
				{
					Name:  "Maps",
					Value: "maps",
				},
				{
					Name:  "Playlist Variables",
					Value: "playlistvars",
				},
			},
		},
	}
	AllModesString = "```Diff\n" +
		"+ Vanilla\n" +
		"aitdm               -Attrition\n" +
		"at                  -Bounty Hunt\n" +
		"coliseum            -Coliseum\n" +
		"cp                  -Amped Hardpoint\n" +
		"ctf                 -Capture the Flag\n" +
		"fd_easy             -Frontier Defense Easy\n" +
		"fd_normal           -Frontier Defense Regular\n" +
		"fd_hard             -Frontier Defense Hard\n" +
		"fd_insane           -Frontier Defense Insane\n" +
		"fd_master           -Frontier Defense Master\n" +
		"lts                 -Last Titan Standing\n" +
		"mfd                 -Marked For Death\n" +
		"ps                  -Pilots vs. Pilots\n" +
		"solo                -Campaign\n" +
		"tdm                 -Skirmish\n" +
		"ttdm                -Titan Brawl\n" +
		"lf                  -Live Fire\n" +
		"+ Vanilla (Featured)\n" +
		"alts                -Aegis Last Titan Standing\n" +
		"attdm               -Aegis Titan Brawl\n" +
		"ffa                 -Free For All\n" +
		"fra                 -Free Agents\n" +
		"holopilot_lf        -The Great Bamboozle\n" +
		"rocket_lf           -Rocket Arena\n" +
		"turbo_lts           -Turbo Last Titan Standing\n" +
		"turbo_ttdm          -Turbo Titan Brawl\n" +
		"+ Northstar.Custom\n" +
		"chamber             -One in the Chamber\n" +
		"ctf_comp            -Competitive CTF\n" +
		"fastball            -Fastball\n" +
		"gg                  -Gun Game\n" +
		"hidden              -The Hidden\n" +
		"hs                  -Hide and Seek\n" +
		"inf                 -Infection\n" +
		"kr                  -Amped Killrace\n" +
		"sns                 -Sticks and Stones\n" +
		"tffa                -Titan FFA\n" +
		"tt                  -Titan Tag\n" +
		"fw                  -Frontier War\n" +
		"+ Northstar.Coop\n" +
		"sp_coop             -Singleplayer Coop```"
	AllMapsString = "```Diff\n" +
		"+ Multiplayer Maps\n" +
		"mp_angel_city        - Angel City\n" +
		"mp_black_water_canal - Black Water Canal\n" +
		"mp_grave             - Boomtown\n" +
		"mp_colony02          - Colony\n" +
		"mp_complex3          - Complex\n" +
		"mp_crashsite3        - Crashsite\n" +
		"mp_drydock           - DryDock\n" +
		"mp_eden              - Eden\n" +
		"mp_thaw              - Exoplanet\n" +
		"mp_forwardbase_kodai - Forward Base Kodai\n" +
		"mp_glitch            - Glitch\n" +
		"mp_homestead         - Homestead\n" +
		"mp_relic02           - Relic\n" +
		"mp_rise              - Rise\n" +
		"mp_wargames          - Wargames\n" +
		"mp_lobby             - Lobby\n" +
		"mp_lf_deck           - Deck\n" +
		"mp_lf_meadow         - Meadow\n" +
		"mp_lf_stacks         - Stacks\n" +
		"mp_lf_township       - Township\n" +
		"mp_lf_traffic        - Traffic\n" +
		"mp_lf_uma            - UMA\n" +
		"mp_coliseum          - The Coliseum\n" +
		"mp_coliseum_column   - Pillars\n" +
		"mp_box               - Box\n" +
		"+ Singleplayer Maps\n" +
		"sp_training          -The Pilot's Gauntlet\n" +
		"sp_crashsite         -BT-7274\n" +
		"sp_sewers1           -Blood and Rust\n" +
		"sp_boomtown_start    -Into the Abyss - Part 1\n" +
		"sp_boomtown          -Into the Abyss - Part 2\n" +
		"sp_boomtown_end      -Into the Abyss - Part 2\n" +
		"sp_hub_timeshift     -Effect and Cause - Part 1 or 3\n" +
		"sp_timeshift_spoke02 -Effect and Cause - Part 2\n" +
		"sp_beacon            -The Beacon - Part 1 or 3\n" +
		"sp_beacon_spoke0     -The Beacon - Part 2\n" +
		"sp_tday              -Trial by Fire\n" +
		"sp_s2s               -The Ark\n" +
		"sp_skyway_v1         -The Fold Weapon```\n"
	AllPlaylistVarString = "```Diff\n" +
		"+ Playlist Vars\n" +
		"custom_air_accel_pilot\n" +
		"pilot_health_multiplier\n" +
		"run_epilogue\n" +
		"respawn_delay\n" +
		"boosts_enabled\n" +
		"earn_meter_pilot_overdrive\n" +
		"earn_meter_pilot_multiplier\n" +
		"earn_meter_titan_multiplier\n" +
		"aegis_upgrades\n" +
		"infinite_doomed_state\n" +
		"titan_shield_regen\n" +
		"scorelimit\n" +
		"roundscorelimit\n" +
		"timelimit\n" +
		"oob_timer_enabled\n" +
		"roundtimelimit\n" +
		"classic_rodeo\n" +
		"classic_mp\n" +
		"fp_embark_enabled\n" +
		"promode_enable\n" +
		"riff_floorislava\n" +
		"featured_mode_all_holopilot\n" +
		"featured_mode_all_grapple\n" +
		"featured_mode_all_phase\n" +
		"featured_mode_all_ticks\n" +
		"featured_mode_tactikill\n" +
		"featured_mode_amped_tacticals\n" +
		"featured_mode_rocket_arena\n" +
		"featured_mode_shotguns_snipers\n" +
		"iron_rules\n" +
		"riff_player_bleedout\n" +
		"player_bleedout_forceHolster\n" +
		"player_bleedout_forceDeathOnTeamBleedout\n" +
		"player_bleedout_bleedoutTime\n" +
		"player_bleedout_firstAidTime\n" +
		"player_bleedout_firstAidTimeSelf\n" +
		"player_bleedout_firstAidHealPercent\n" +
		"player_bleedout_aiBleedingPlayerMissChance\n" +
		"+ Sticks and Stones Playlist Vars\n" +
		"sns_softball_enabled\n" +
		"sns_softball_kill_value\n" +
		"sns_wme_kill_value\n" +
		"sns_offhand_kill_value\n" +
		"sns_reset_kill_value\n" +
		"sns_melee_kill_value\n" +
		"sns_reset_pulse_blade_cooldown_on_pulse_blade_kill```"
)

func ListCmdHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var sendString string
	o := i.ApplicationCommandData().Options[0].StringValue()

	if o == "maps" {
		sendString = AllMapsString
	} else if o == "modes" {
		sendString = AllModesString
	} else if o == "playlistvars" {
		sendString = AllPlaylistVarString
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: sendString,
		},
	})

}
