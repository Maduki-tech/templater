use std::{
    fs::{self, File},
    io::Write,
};

#[derive(Debug)]
pub struct FileCreation {
    language: String,
}

impl FileCreation {
    pub fn new(language: String) -> Self {
        Self { language }
    }

    pub fn create(&self) {
        match self.language.as_str() {
            "C++" => {
                self.cpp_project();
            }
            _ => {
                println!("No project created");
            }
        }
    }

    fn cpp_project(&self) {
        let mut file = File::create("CMakeLists.txt").expect("Unable to create file");
        let content = " cmake_minimum_required (VERSION 3.5)

project (main)

# Set C and C++ standards
set(CMAKE_C_STANDARD 11)
set(CMAKE_C_STANDARD_REQUIRED ON)
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)

include_directories(${PROJECT_SOURCE_DIR}/includes)

# Glob source files
file(GLOB_RECURSE C_SOURCES \"./src/*.c\")
file(GLOB_RECURSE CPP_SOURCES \"./src/*.cpp\")
file(GLOB_RECURSE HEADERS \"./includes/*.h\")
file(GLOB_RECURSE MAIN \"./main.cpp\")

# Add all sources to the executable
add_executable(main ${MAIN} ${C_SOURCES} ${CPP_SOURCES} ${HEADERS})

# Find Google Test
find_package(GTest REQUIRED)
include_directories(${GTEST_INCLUDE_DIRS})

# Add test executable
file(GLOB_RECURSE TEST_SOURCES \"./tests/*.cpp\")
add_executable(runUnitTests ${TEST_SOURCES} ${C_SOURCES} ${CPP_SOURCES} ${HEADERS})
target_link_libraries(runUnitTests GTest::gtest_main)


# This is so you can run your tests with CTest
add_test(NAME unitTests COMMAND runUnitTests)

            ";

        file.write_all(content.as_bytes())
            .expect("Unable to write data");

        fs::create_dir("src").expect("Unable to create directory");
        fs::create_dir("includes").expect("Unable to create directory");

        let mut file = File::create("main.cpp").expect("Unable to create file");
        let content = "#include <iostream>

int main(int argc, char *argv[])
{
    std::cout << \"Hello World\" << std::endl;
    return 0;
}";
        file.write_all(content.as_bytes())
            .expect("Unable to write data");

        let mut file = File::create("includes/add.h").expect("Unable to create file");
        let content = "#pragma once

int add(int a, int b);";
        file.write_all(content.as_bytes())
            .expect("Unable to write data");

        let mut file = File::create("src/add.cpp").expect("Unable to create file");
        let content = "#include \"../includes/add.h\"

int add(int a, int b)
{
    return a + b;
}";
        file.write_all(content.as_bytes())
            .expect("Unable to write data");

        fs::create_dir("tests").expect("Unable to create directory");

        let mut file = File::create("tests/add_test.cpp").expect("Unable to create file");
        let content = "#include <gtest/gtest.h>
#include \"../includes/add.h\"

TEST(AddTest, PositiveNumbers)
{
    EXPECT_EQ(add(1, 2), 3);

    EXPECT_EQ(add(10, 20), 30);

    EXPECT_EQ(add(100, 200), 300);

    EXPECT_EQ(add(1000, 2000), 3000);

}

TEST(AddTest, NegativeNumbers)
{
    EXPECT_EQ(add(-1, -2), -3);

    EXPECT_EQ(add(-10, -20), -30);

    EXPECT_EQ(add(-100, -200), -300);

    EXPECT_EQ(add(-1000, -2000), -3000);

}
        ";

        file.write_all(content.as_bytes())
            .expect("Unable to write data");

    }


}
