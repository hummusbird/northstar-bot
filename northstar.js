const Discord = require('discord.js');
const fetch = require('node-fetch')
const client = new Discord.Client();
require('dotenv').config();

client.login(process.env.BOT_TOKEN)

client.on('ready', async () => {
    console.log(`Logged in as ${client.user.tag}!`);
});

const url = "https://northstar.tf/client/servers"
const urlR = /^(?:([A-Za-z]+):)?(\/{0,3})([0-9.\-A-Za-z]+)(?::(\d+))?(?:\/([^?#]*))?(?:\?([^#]*))?(?:#(.*))?$/;
const prefix = ','
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
    "hs": "Hide and Seek"
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
    if (!msg.content.startsWith(prefix)) return;

    var args = msg.content.split(" ");
    args[0] = args[0].substring(1, args[0].length)

    switch (args[0]) {

        case "help":
            msg.channel.send(`\`\`\`diff
+ Here are a list of all available commands:
${prefix}status                 - a general overview of northstar.tf
${prefix}search title [string]  - searches server titles
${prefix}search mode [gamemode] - searches all servers running that mode
${prefix}search map [map]       - searches all servers running that map
${prefix}convars                - lists all available convars
${prefix}modes                  - lists all Titanfall 2 gamemodes
${prefix}maps                   - lists all Titanfall 2 maps
${prefix}host                   - links hummusbird's tutorial
\`\`\``)

            break;

        case "status":
            var data = await getServers(url)

            if (typeof data == typeof "string") {
                msg.channel.send(`\`\`\`diff\n- ${data}\`\`\``)
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
                msg.channel.send(`\`\`\`diff\n
## NORTHSTAR.TF STATUS: ##\n
+ Servers Online: ${data.length}\n
- Password Protected Servers: ${protectedLobbies}\n
+ Players in-game: ${playersOnline}/${playerSlots} (${Math.round((playersOnline / playerSlots) * 100)}%)
\`\`\``)

            }

            break;

        case "search":
            var data = await getServers(url)

            if (typeof data == typeof "string") {
                msg.channel.send(`\`\`\`diff\n- ${data}\`\`\``)
            }
            else if (!args[1]) { msg.channel.send(`\`\`\`diff\n- Please specify title, map or mode.\`\`\``) }
            else if (!args[2]) { msg.channel.send(`\`\`\`diff\n- Please specify a search term.\`\`\``) }
            else {
                var parameter;
                if (args[1] == "title" || args[1] == "name") {
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
                else { return msg.channel.send(`\`\`\`diff\n- Please specify title, map or mode.\`\`\``) }

                var lobbies = [];
                for (i = 0; i < data.length; i++) {
                    if (data[i][parameter].toLowerCase().includes(args[2].toLowerCase())) {
                        lobbies.push(data[i])
                    }
                }

                if (lobbies.length == 0) {
                    msg.channel.send(`\`\`\`diff\n- No servers were found.\`\`\``)
                }
                else {
                    var searchstring = `\`\`\`diff\n+ ${lobbies.length} servers were found${lobbies.length > 9 ? " - displaying first 10 results" : "."}\n`
                    try {
                        for (i = 0; i < (lobbies.length < 10 ? lobbies.length : 10); i++) {
                            searchstring += `
${lobbies[i]["name"]}
${lobbies[i]["playerCount"] == lobbies[i]["maxPlayers"] ? "-" : "+"} ${lobbies[i]["playerCount"]}/${lobbies[i]["maxPlayers"]} players connected
${lobbies[i]["map"] == "mp_lobby" ? "- Currently in the lobby\n" : `+ Playing ${getGamemode(lobbies[i]["playlist"])} on ${getMapName(lobbies[i]["map"])}${lobbies[i]["hasPassword"] ? `\n- PASSWORD PROTECTED!` : ""}
`}`

                        }
                    }
                    catch {
                        searchstring = "```diff\n- Search failed. Please try again"

                    }
                    msg.channel.send(searchstring + "```")
                }
            }
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
hs        - Hide and Seek\`\`\``)
            break;

        case "host":
        case "birb":
        case "vid":
            msg.channel.send("https://youtu.be/EZ3w2Nl9SZo")
            break;

        case "bio":
            if (msg.author.id == 375671695240855553 /* hummusbird */) {
                var status = args[1]
                var statusType = args[2]
                var words = msg.content.split(statusType)[1].trim()

                if ((status == "online" || status == "idle" || status == "dnd" || status == "invisible") && (statusType == "STREAMING" || statusType == "LISTENING" || statusType == "PLAYING" || statusType == "WATCHING" || statusType == "COMPETING")) {

                    if (statusType == "STREAMING" || statusType == "WATCHING") {
                        client.user.setPresence({
                            status: status,
                            activity: {
                                name: words,
                                url: "https://www.twitch.tv/hummusbirb",
                                type: statusType
                            }
                        })
                    } 
                    else {
                        client.user.setPresence({
                            status: status,
                            activity: {
                                name: words,
                                type: statusType
                            }
                        })
                    }

                    msg.channel.send("```diff\n+ status set```")
                    console.log(`${msg.author.username} set status to ${status}, ${statusType}, ${words}`)
                } 
                else {
                    msg.channel.send("```diff\n- invalid lol >:)```")
                }
            }

            break;
    }
})