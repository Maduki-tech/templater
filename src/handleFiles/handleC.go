package handleFiles

import (
	"os"
)

/*
CreateCFile creates a C file with a makefile or CMakeLists.txt

Input:
	"makefile"
	"cmake"

@param complier string - The complier you want to use for the Project
*/
func CreateCFile(complier string) {
	if complier == "makefile" {
		writeMakefile()
	} else if complier == "cmake" {
		writeCMakeLists()
	}
	writeMain()
	createDirs()
}

func writeCMakeLists() {
	panic("unimplemented")
}

func writeMakefile() {
	os.WriteFile("Makefile",
		[]byte(`CC=gcc
CFLAGS=-Iinclude -Wall -g
# here you can add the DEPS you want to use 
# Example: DEPS = libs/myheader.h
DEPS = 
OBJ = main.o

# Pattern rule for object files
%.o: %.c $(DEPS)
	$(CC) $(CFLAGS) -c -o $@ $< 

# Rule for building the final executable
main: $(OBJ)
	$(CC) $(CFLAGS) -o $@ $^

.PHONY: clean compile_commands

clean:
	rm -f *.o main  # Updated to remove the 'src/' as your object file is not in 'src/'

compile_commands:
	bear -- make
	`), 0644)
}

func writeMain() {
	os.WriteFile("main.c",
		[]byte(`#include <stdio.h>

int main() {
	printf("Hello World!");
	return 0;
}`), 0644)
}

func createDirs() {
	os.Mkdir("src", 0755)
	os.Mkdir("libs", 0755)
}
