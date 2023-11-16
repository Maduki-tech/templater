mod ui;
mod file_creation;

use clap::Parser;
use ui::UI;
use file_creation::FileCreation;
use warp::Filter;

#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
struct Args {
    #[arg(short, long)]
    language: Option<String>,

    #[arg(short, long, default_value_t = 1)]
    count: u8,
}

#[tokio::main]
async fn main() {
    let static_files = warp::fs::dir("static"); 

    let routes = warp::path("app").and(static_files);

    warp::serve(routes).run(([127, 0, 0, 1], 3030)).await;
    
    let _args = Args::parse();

    if let Some(language) = _args.language {
        FileCreation::new(language).create();
        return;
    }

    let selected = UI::new().run();

    FileCreation::new(selected).create();
}
