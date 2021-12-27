const Discord = require('discord.js');
const fetch = require('node-fetch')
const client = new Discord.Client();
require('dotenv').config();

client.login(process.env.BOT_TOKEN)

client.on('ready', async () => {
    console.log(`Logged in as ${client.user.tag}!`);
    getServers();
});

const prefix = ','

async function getServers() {

        let res = await fetch("https://northstar.tf/client/servers")
        let parsed = await res.json()
        if (parsed && Object.keys(parsed).length === 0 && parsed.constructor === Object){ //API error
            throw("api error")
        }
        else{ //API success
            try{
                console.log(parsed)
            }
            catch{
                console.log("invalid api call")
            }
        }

}

client.on('message', async msg => {
    if (!msg.guild || msg.author.bot) return;
    if (!msg.content.startsWith(prefix)) return;

    var args = msg.content.split(" ");
    args[0] = args[0].substring(1, args[0].length)

    switch (args[0]){

        case "help":

            break;

        case "test":
            console.log("sdf");
            break;
    }
})