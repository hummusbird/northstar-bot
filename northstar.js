const Discord = require('discord.js');
const fetch = require('node-fetch')
const client = new Discord.Client();
require('dotenv').config();

client.login(process.env.BOT_TOKEN)

client.on('ready', async () => {
    console.log(`Logged in as ${client.user.tag}!`);
});

const urlR = /^(?:([A-Za-z]+):)?(\/{0,3})([0-9.\-A-Za-z]+)(?::(\d+))?(?:\/([^?#]*))?(?:\?([^#]*))?(?:#(.*))?$/;
const prefix = ','

function getMapName(name) {
    var maps = {
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
        "mp_coliseum_column": "Pillars",

    }
    return maps[name]
}

function getGamemode(mode) {
    var modes = {
        "tdm" : "Skirmish",
        "cp" : "Amped Hardpoint",
        "ctf" : "Capture the Flag",
        "lts" : "Last Titan Standing",
        "ps" : "Pilots V Pilots",
        "ffa" : "Free For All",
        "speedball" : "Live Fire",
        "mfd" : "Marked for Death",
        "ttdm": "Titan Brawl",
        "fra" : "Free Agents",
        "gg" : "Gun Game",
        "inf" : "Infection",
        "tt" : "Titan Tag",
        "kr" : "Amped Killrace",
        "fastball" : "Fastball",
        "arena" : "1v1 Arena",
        "ctf_comp" : "Capture the Flag",
        "hs" : "Hide and Seek"
    }
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
        console.log("api serror")
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
            msg.channel.send(`\`\`\`diff\n+ Here are a list of all available commands:\n${prefix}status - a general overview of northstar.tf\n${prefix}search [string] - searches server titles\`\`\``)
            break;

        case "status":
            var url = "https://northstar.tf/client/servers"
            var data = await getServers(url)

            if (typeof data == typeof "string") {
                msg.channel.send(`\`\`\`diff\n- ${data}\`\`\``)
            }
            else {
                var playersOnline = 0;
                var protectedLobbies = 0;
                for (i = 0; i < data.length; i++) {
                    playersOnline += data[i]["playerCount"];
                    protectedLobbies += data[i]["hasPassword"] ? 1 : 0;
                }
                msg.channel.send(`\`\`\`diff\n- ## NORTHSTAR.TF STATUS: ##\n+ Servers Online: ${data.length}\n+ Players in-game: ${playersOnline}\n+ Password Protected Servers: ${protectedLobbies}\`\`\``)
            }

            break;

        case "search":
            var url = "https://northstar.tf/client/servers"
            var data = await getServers(url)

            if (typeof data == typeof "string") {
                msg.channel.send(`\`\`\`diff\n- ${data}\`\`\``)
            }
            else {
                var lobbies = [];
                for (i = 0; i < data.length; i++) {
                    if (data[i]["name"].toLowerCase().includes(args[1].toLowerCase())) {
                        lobbies.push(data[i])
                    }
                }

                if (lobbies.length == 0) {
                    msg.channel.send(`\`\`\`diff\n- No servers were found.\`\`\``)
                }
                else {
                    var searchstring = `\`\`\`diff\n+ ${lobbies.length} servers were found!\n`

                    for (i = 0; i < lobbies.length; i++) {
                        searchstring += `\n${lobbies[i]["name"]}\n${lobbies[i]["playerCount"] == lobbies[i]["maxPlayers"] ? "-" : "+"} ${lobbies[i]["playerCount"]}/${lobbies[i]["maxPlayers"]} players connected\n${lobbies[i]["map"] == "mp_lobby" ? "- Currently in the lobby" : `+ Playing ${getGamemode(lobbies[i]["playlist"])} on ${getMapName(lobbies[i]["map"])}`}`
                    }
                    msg.channel.send(searchstring + "```")
                }

                break;
            }
        }
    })