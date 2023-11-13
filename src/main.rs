mod ui;
mod file_creation;
use clap::Parser;
use ui::UI;
use file_creation::FileCreation;

#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
struct Args {
    #[arg(short, long)]
    language: Option<String>,

    #[arg(short, long, default_value_t = 1)]
    count: u8,
}

fn main() {
    let _args = Args::parse();


    if let Some(language) = _args.language {
        FileCreation::new(language).create();
        return;
    }
    let selected = UI::new().run();

    FileCreation::new(selected).create();


}
