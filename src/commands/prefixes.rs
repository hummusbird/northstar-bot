use rusqlite::{Connection, Error, Result};
use serenity::framework::standard::macros::command;
use serenity::framework::standard::CommandResult;
use serenity::model::prelude::GuildId;
use serenity::model::prelude::*;
use serenity::prelude::*;

use crate::DEFAULTPREFIX;

#[derive(Debug)]
struct Server {
    id: u64,
    prefix: String,
}
#[command]
#[required_permissions(ADMINISTRATOR)]
async fn prefix(ctx: &Context, msg: &Message) -> CommandResult {
    let args: Vec<&str> = msg.content.split(" ").collect();
    let conn = Connection::open("./prefix.db").expect("db fucked");
    let prefix = args[1].chars().nth(0).unwrap();
    conn.execute(
        &("DELETE FROM Servers WHERE id = ".to_owned() + &msg.guild_id.unwrap().to_string()),
        (),
    )?;

    let new_server = Server {
        id: *msg.guild_id.unwrap().as_u64(),
        prefix: prefix.to_string(),
    };

    conn.execute(
        "INSERT INTO Servers (id, prefix) VALUES (?1, ?2)",
        (&new_server.id, &new_server.prefix),
    )?;

    msg.channel_id
        .say(
            ctx,
            "```diff
+ Prefix changed to "
                .to_owned()
                + &prefix.to_string()
                + " ```",
        )
        .await?;

    Ok(())
}

pub fn check_db_prefix(guild_id: Option<GuildId>) -> Option<String> {
    let conn = Connection::open("./prefix.db").expect("db fucked");

    match conn
        .prepare(
            &("SELECT prefix FROM Servers WHERE id LIKE ".to_owned()
                + &guild_id.unwrap().as_u64().to_string()),
        )
        .expect("fail")
    {
        mut stmt => {
            let mut rows = stmt.query(()).unwrap();
            let mut names: Vec<String> = Vec::new();

            while let Some(row) = rows.next().unwrap() {
                names.push(row.get(0).unwrap());
            }
            return Some(names[0].to_string());
        }
    }
}

pub async fn new_server_reg(guild_id: u64) -> Result<(), Error> {
    let conn = Connection::open("./prefix.db").expect("db fucked");

    conn.execute(
        "CREATE TABLE IF NOT EXISTS Servers (
             id integer null unique,
             prefix text not null 
         )",
        (),
    )?;

    match conn
        .prepare(&("SELECT id FROM Servers WHERE id LIKE ".to_owned() + &guild_id.to_string()))
        .expect("fail")
    {
        mut stmt => {
            let mut rows = stmt.query(()).unwrap();
            let mut names: Vec<usize> = Vec::new();

            while let Some(row) = rows.next().unwrap() {
                names.push(row.get(0).unwrap());
            }

            if names.is_empty() == true {
                let new_server = Server {
                    id: guild_id,
                    prefix: DEFAULTPREFIX.to_string(),
                };

                println!("New Server added");
                conn.execute(
                    "INSERT INTO Servers (id, prefix) VALUES (?1, ?2)",
                    (&new_server.id, &new_server.prefix),
                )?;
            }
        }
    }

    Ok(())
}
