use dialoguer::Select;

#[derive(Debug)]
pub struct UI {}

impl UI {
    pub fn new() -> Self {
        Self {}
    }

    pub fn run(&self) -> String {
        let programming_language = vec!["Rust", "Python", "C++"];

        let selection = Select::new()
            .with_prompt("What project do you wnat to create?")
            .items(&programming_language)
            .default(0)
            .interact()
            .unwrap();

        programming_language[selection].to_string()
    }
}
