CC=gcc
CFLAGS=-Iinclude -Wall -g
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
