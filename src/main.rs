use serenity::model::prelude::Guild;
use std::env;

use serenity::async_trait;
use serenity::framework::standard::macros::group;
use serenity::framework::standard::StandardFramework;
use serenity::model::channel::Message;
use serenity::model::gateway::Activity;
use serenity::model::gateway::Ready;
use serenity::model::user::OnlineStatus;
use serenity::prelude::*;

mod commands;

use crate::commands::links::*;
use crate::commands::lists::*;
use crate::commands::northstar::*;
use crate::commands::prefixes::*;
use crate::commands::titancoins::*;

static DEFAULTPREFIX: &str = ",";
static MS: &str = "https://northstar.tf";

#[group("GENERAL")]
#[commands(prefix)]
struct General;

#[group("LIST")]
#[commands(maps, modes, playlistvars, help)]
struct List;

#[group("LINKS")]
#[commands(birb, github, wiki, info)]
struct Link;

#[group("NORTHSTAR")]
#[commands(status, search)]
struct Northstar;

struct Handler;
#[async_trait]
impl EventHandler for Handler {
    async fn ready(&self, ctx: Context, ready: Ready) {
        println!("Connected as {}", ready.user.name);
        let guilds = ctx.cache.guilds().len();

        println!("The bot is in {} guilds", guilds);
        set_activity(ctx).await;
    }
    async fn guild_create(&self, _ctx: Context, guild: Guild, _is_new: bool) {
        new_server_reg(*guild.id.as_u64()).await.expect("fuck");
    }
    async fn message(&self, ctx: Context, msg: Message) {
        if msg.content == "/redeem" {
            if let Err(why) = redeem(&ctx, &msg).await {
                println!("Error sending message: {:?}", why);
            }
        }
        if msg.content.contains("<@925064195186233344>") {
            if let Err(why) = msg.reply_ping(ctx, "what").await {
                println!("Error sending message: {:?}", why);
            }
        }
    }
}

async fn set_activity(ctx: Context) {
    let activity = Activity::playing("Northstar.TF");
    let status = OnlineStatus::Online;
    ctx.set_presence(Some(activity), status).await;
}

#[tokio::main]
async fn main() {
    let args: Vec<_> = env::args().collect();
    if args.len() > 1 {
        if args[1] == "-dev" {
            println!("--- Dev mode ---");
            println!("Panic logging occurs now");
        } else {
            std::panic::set_hook(Box::new(|_info| {}));
        }
    }

    let mut framework = StandardFramework::new().configure(|c| {
        c.dynamic_prefix(|_, msg| Box::pin(async move { check_db_prefix(msg.guild_id) }))
    });

    framework.group_add(&GENERAL_GROUP);
    framework.group_add(&LIST_GROUP);
    framework.group_add(&LINK_GROUP);
    framework.group_add(&NORTHSTAR_GROUP);

    dotenv::dotenv().expect("Failed to load .env file");

    let token = env::var("DISCORD_TOKEN").expect("token");
    let intents = GatewayIntents::non_privileged() | GatewayIntents::MESSAGE_CONTENT;
    let mut client = Client::builder(token, intents)
        .event_handler(Handler)
        .framework(framework)
        .await
        .expect("Error creating client");

    if let Err(why) = client.start().await {
        println!("An error occurred while running the client: {:?}", why);
    }
}
