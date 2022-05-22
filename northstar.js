const Discord = require('discord.js');
fs = require('fs')
const fetch = require('node-fetch')
const moment = require('moment')
const client = new Discord.Client();
require('dotenv').config();

let guildSettings = []

client.login(process.env.BOT_TOKEN)

client.on('ready', async () => {
    console.log(`Logged in as ${client.user.tag}!`);
    client.user.setActivity('Northstar.TF', { type: 'PLAYING' })

    let guilds = client.guilds.cache.map(guild => guild);
    console.log(`The bot is in ${guilds.length} guilds`);

    console.log("\x1b[35m%s\x1b[0m", `Loading guild settings:`)
    for(let guild of guilds) {
        let exists = fs.existsSync(`${guild.id}_config.json`)
        await (exists ? loadConfig(guild) : createConfig(guild))
    }
});

client.on('guildCreate', async guild => {
    console.log("\x1b[32m", `Joined new guild: ${guild.name}`)
    let exists = fs.existsSync(`${guild.id}_config.json`)
    await (exists ? loadConfig(guild) : createConfig(guild))
})

client.on('guildDelete', async guild => {
    console.log("\x1b[31m", `Kicked from Guild: ${guild.name}`)
})

// CONFIG FUNCTIONS

async function createConfig(e) {
    console.log("\x1b[33m%s\x1b[0m",`Guild ${e.name} does not have settings file, Creating....`)
    
    var config = {
        guildID: e.id,
        prefix: ",",
    }

    fs.appendFile(`${e.id}_config.json`, JSON.stringify(config), function (err) {
        if (err) throw err;
        console.log('Saved!')
        loadConfig(e)
    })

}

async function loadConfig(e) {
    console.log(`loading ${e.name} config file.`)
    if (guildSettings.find(config => config.guildID == e.id)){
        var index = guildSettings.indexOf(e)
        var data = fs.readFileSync(`${e.id}_config.json`, 'utf8')
        guildSettings[index] = JSON.parse(data)
        console.log("\x1b[33m%s\x1b[0m",`${e.name} reloaded`)
    }
    else {
        try {
            var data = fs.readFileSync(`${e.id}_config.json`, 'utf8')
            guildSettings.push(JSON.parse(data))
            console.log("\x1b[32m%s\x1b[0m",`${e.name} successfully loaded.`)
        } 
        catch(error) {
            console.log(error);
            return -1
        }
    }
}

function writeConfig(config, guild){
    fs.writeFile(`${guild.id}_config.json`, JSON.stringify(config), function (err) {
        if (err) throw err;
        console.log(`Saved new config for ${guild.name}`)
        loadConfig(guild)
    })
}

const url = "https://northstar.tf/client/servers"
const maps = {
    "mp_angel_city": "Angel City",
    "mp_black_water_canal": "Black Water Canal",
    "mp_grave": "Boomtown",
    "mp_colony02": "Colony",
    "mp_complex3": "Complex",
    "mp_crashsite3": "Crashsite",
    "mp_drydock": "DryDock",
    "mp_eden": "Eden",
    "mp_thaw": "Exoplanet",
    "mp_forwardbase_kodai": "Forward Base Kodai",
    "mp_glitch": "Glitch",
    "mp_homestead": "Homestead",
    "mp_relic02": "Relic",
    "mp_rise": "Rise",
    "mp_wargames": "Wargames",
    "mp_lobby": "Lobby",
    "mp_lf_deck": "Deck",
    "mp_lf_meadow": "Meadow",
    "mp_lf_stacks": "Stacks",
    "mp_lf_township": "Township",
    "mp_lf_traffic": "Traffic",
    "mp_lf_uma": "UMA",
    "mp_coliseum": "The Coliseum",
    "mp_coliseum_column": "Pillars"
}

const modes = {
    "tdm": "Skirmish",
    "aitdm": "Attrition",
    "sns": "Sticks and Stones",
    "cp": "Amped Hardpoint",
    "ctf": "Capture the Flag",
    "lts": "Last Titan Standing",
    "ps": "Pilots V Pilots",
    "ffa": "Free For All",
    "speedball": "Live Fire",
    "mfd": "Marked for Death",
    "ttdm": "Titan Brawl",
    "fra": "Free Agents",
    "gg": "Gun Game",
    "inf": "Infection",
    "tt": "Titan Tag",
    "kr": "Amped Killrace",
    "fastball": "Fastball",
    "arena": "1v1 Arena",
    "ctf_comp": "Capture the Flag",
    "hs": "Hide and Seek",
    "hidden": "The Hidden",
    "chamber": "One in the Chamber"
}

function getMapName(name) {
    return maps[name]
}

function getGamemode(mode) {
    return modes[mode]
}

async function getServers(url) {
    try {
        let res = await fetch(url)
        let parsed = await res.json()
        if (parsed && Object.keys(parsed).length === 0 && parsed.constructor === Object) { //API error
            throw ("api error")
        }
        else { //API success
            try {
                return parsed
            }
            catch {
                console.log("invalid api call")
                return `Unable to parse data`
            }
        }
    }
    catch (e) {
        console.log("api error")
        return `${url} is unavailable!`
    }
}

client.on('message', async msg => {
    if (!msg.guild || msg.author.bot) return;

    else if (msg.content.includes("<@!925064195186233344>")){
        client.api.channels[msg.channel.id].messages.post({
            data: {
                content: "what",
                    message_reference: {
                    message_id: msg.id,
                    channel_id: msg.channel.id,
                    guild_id: msg.guild.id
                }
            }
        })
    }

    const config = guildSettings.find(config => config.guildID == msg.guild.id)

    if (!msg.content.startsWith(config.prefix)) return;

    var args = msg.content.split(" ");
    args[0] = args[0].substring(1, args[0].length)

    switch (args[0]) {
        case "help":
            msg.channel.send(`\`\`\`diff
+ Here are a list of all available commands:
${config.prefix}status                 - a general overview of northstar.tf
${config.prefix}search title [string]  - searches server titles
${config.prefix}search mode [gamemode] - searches all servers running that mode
${config.prefix}search map [map]       - searches all servers running that map
${config.prefix}convars                - lists some useful ConVars
${config.prefix}modes                  - lists all Titanfall 2 gamemodes
${config.prefix}maps                   - lists all Titanfall 2 maps
${config.prefix}host                   - links hummusbird's server tutorial
${config.prefix}git                    - links the github
${config.prefix}wiki                   - links the wiki
\`\`\``)

            break;

        case "status":
            var data = await getServers(url)

            if (typeof data == typeof "string") {
                msg.channel.send("https://media.discordapp.net/attachments/891428299430043708/955612322427191296/motivate.gif")
            }
            else {
                var playersOnline = 0;
                var protectedLobbies = 0;
                var playerSlots = 0;
                for (i = 0; i < data.length; i++) {
                    var hasPwd = data[i]["hasPassword"]

                    protectedLobbies += hasPwd ? 1 : 0;
                    playersOnline += hasPwd ? 0 : (data[i]["playerCount"] == undefined ? 0 : data[i]["playerCount"]);
                    playerSlots += hasPwd ? 0 : data[i]["maxPlayers"];
                }
                var percentage = playerSlots > 0 ? Math.round((playersOnline / playerSlots) * 100) : 0;
                msg.channel.send(`\`\`\`diff\n
## NORTHSTAR.TF STATUS: ##\n
+ Servers Online: ${data.length}\n
- Password Protected Servers: ${protectedLobbies}\n
+ Players in-game: ${playersOnline}/${playerSlots} (${percentage}%)
\`\`\``)
            }

            break;

        case "search":
            if (!args[1]) { return msg.channel.send(`\`\`\`diff\n- Please specify title, map or mode.\`\`\``) }
            else {
                var search = args[2];
                var parameter = "name";
                if (args[1] == "title" || args[1] == "name") {
                    if (!args[2]) {return msg.channel.send(`\`\`\`diff\n- Please specify a search term.\`\`\``)}
                    parameter = "name"
                }
                else if (args[1] == "map" || args[1] == "maps") {
                    if (getMapName(args[2]) == undefined) { return msg.channel.send(`\`\`\`diff\n- Please specify a valid map.\`\`\``) }
                    parameter = "map"
                }
                else if (args[1] == "mode" || args[1] == "modes" || args[1] == "gamemode"){
                    if (getGamemode(args[2]) == undefined) { return msg.channel.send(`\`\`\`diff\n- Please specify a valid gamemode.\`\`\``) }
                    parameter = "playlist"
                }

                else { search = args[1] }

                var data = await getServers(url)

                if (typeof data == typeof "string") {
                    return msg.channel.send(`\`\`\`diff\n- ${data}\`\`\``)
                }

                var lobbies = [];
                for (i = 0; i < data.length; i++) {
                    if (data[i][parameter].toLowerCase().includes(search.toLowerCase())) {
                        lobbies.push(data[i])
                    }
                }

                if (lobbies.length == 0) {
                    msg.channel.send(`\`\`\`diff\n- No servers were found.\`\`\``)
                }
                else {
                    var search_playersOnline = 0;
                    var search_playerSlots = 0;
                    var searchstring = `\`\`\`diff\n+ ${lobbies.length} servers were found${lobbies.length > 10 ? " - displaying first 10 results" : "."}\n`
                    try {
                        for (i = 0; i < lobbies.length; i++) {
                            search_playersOnline += lobbies[i]["playerCount"]
                            search_playerSlots += lobbies[i]["maxPlayers"]
                            if (i < 10) {searchstring += `
${lobbies[i]["name"].replace(/`/g,'')}
${lobbies[i]["playerCount"] == lobbies[i]["maxPlayers"] ? "-" : "+"} ${lobbies[i]["playerCount"]}/${lobbies[i]["maxPlayers"]} players connected
${lobbies[i]["map"] == "mp_lobby" ? "- Currently in the lobby\n" : `+ Playing ${getGamemode(lobbies[i]["playlist"])} on ${getMapName(lobbies[i]["map"])}${lobbies[i]["hasPassword"] ? `\n- PASSWORD PROTECTED!` : ""}
`}`

                        }}
                    }
                    catch (e) {
                        searchstring = "```diff\n- Search failed. Please try again"
                        console.log(e)
                    }
                    
                    var ratingstring = `\`\`\`diff\n\nTotal ${search_playersOnline}/${search_playerSlots} players in "${search}" servers, for:\n\n+ ${search_playersOnline}/${playersOnline} or ${Math.round((search_playersOnline / playersOnline)*10000)/100}% of all NS players\n\n- ${lobbies.length}/${data.length} or ${Math.round((lobbies.length / data.length)*10000)/100}% of all NS servers\`\`\``
                    msg.channel.send((args[args.length -1] == "stats") ? ratingstring : searchstring + "```")
                }
            }
            break;

        case "prefix":
            if (!msg.member.hasPermission("ADMINISTRATOR")) { return msg.channel.send("```diff\n- no <3```")}

            if (!args[1]) {msg.channel.send(`\`\`\`Current Prefix is '${config.prefix}'\`\`\``)}
            else {
                config.prefix = args[1]
                writeConfig(config, msg.guild)
                msg.channel.send(`\`\`\`diff\n+ Prefix changed to '${config.prefix}'\`\`\``)
            }
            break;

        case "uptime":
            var time = moment.duration(moment().diff(msg.client.readyAt))
            msg.channel.send(`This bot has been up for ${time.days()} days, ${time.hours()} hours, ${time.minutes()} minutes and ${time.seconds()} seconds`)
            break;
        
        case "maps":
        case "map":
            msg.channel.send(`\`\`\`diff\n+ Titanfall 2 Maps:
mp_angel_city        - Angel City
mp_black_water_canal - Black Water Canal
mp_grave             - Boomtown
mp_colony02          - Colony
mp_complex3          - Complex
mp_crashsite3        - Crashsite
mp_drydock           - DryDock
mp_eden              - Eden
mp_thaw              - Exoplanet
mp_forwardbase_kodai - Forward Base Kodai
mp_glitch            - Glitch
mp_homestead         - Homestead
mp_relic02           - Relic
mp_rise              - Rise
mp_wargames          - Wargames
mp_lobby             - Lobby
mp_lf_deck           - Deck
mp_lf_meadow         - Meadow
mp_lf_stacks         - Stacks
mp_lf_township       - Township
mp_lf_traffic        - Traffic
mp_lf_uma            - UMA
mp_coliseum          - The Coliseum
mp_coliseum_column   - Pillars\`\`\``)

            break;

        case "mode":
        case "modes":
        case "gamemodes":
        case "gamemode":
            msg.channel.send(`\`\`\`diff\n+ Titanfall 2 Gamemodes:
tdm       - Skirmish
aitdm     - Attrition
sns       - Sticks and Stones
cp        - Amped Hardpoint
ctf       - Capture the Flag
lts       - Last Titan Standing
ps        - Pilots V Pilots
ffa       - Free For All
speedball - Live Fire
mfd       - Marked for Death
ttdm      - Titan Brawl
fra       - Free Agents
gg        - Gun Game
inf       - Infection
tt        - Titan Tag
kr        - Amped Killrace
fastball  - Fastball
arena     - 1v1 Arena
ctf_comp  - Capture the Flag
hs        - Hide and Seek
hidden    - The Hidden
chamber   - One in the Chamber\`\`\``)
            break;

        case "host":
        case "birb":
        case "vid":
            msg.channel.send("https://youtu.be/EZ3w2Nl9SZo")
            break;

        case "wiki":
            msg.channel.send("https://r2northstar.gitbook.io/r2northstar-wiki/")
            break;
        
        case "git":
        case "github":
            msg.channel.send("https://github.com/R2Northstar")
            break;

        case "convars":
            msg.channel.send(`\`\`\`diff
custom_air_accel_pilot
pilot_health_multiplier
run_epilogue
respawn_delay

boosts_enabled
earn_meter_pilot_overdrive
earn_meter_pilot_multiplier

earn_meter_titan_multiplier
aegis_upgrades
infinite_doomed_state
titan_shield_regen

scorelimit
roundscorelimit
timelimit
oob_timer_enabled
roundtimelimit

classic_rodeo
classic_mp
fp_embark_enabled
promode_enable

riff_floorislava
featured_mode_all_holopilot
featured_mode_all_grapple
featured_mode_all_phase
featured_mode_all_ticks
featured_mode_tactikill
featured_mode_amped_tacticals
featured_mode_rocket_arena
featured_mode_shotguns_snipers
iron_rules

riff_player_bleedout
player_bleedout_forceHolster
player_bleedout_forceDeathOnTeamBleedout
player_bleedout_bleedoutTime
player_bleedout_firstAidTime
player_bleedout_firstAidTimeSelf
player_bleedout_firstAidHealPercent
player_bleedout_aiBleedingPlayerMissChance
\`\`\``)
        break;
    }
})
