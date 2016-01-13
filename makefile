CC=gcc
CFLAGS=-c -Wall -std=gnu11 -fsanitize=thread -fPIC
LDFLAGS= -fsanitize=thread -pie -lpthread
SOURCES=count.c
OBJECTS=$(SOURCES:.c=.o)
EXECUTABLE=count

all: $(SOURCES) $(EXECUTABLE)
	
$(EXECUTABLE): $(OBJECTS)
	$(CC)  $(OBJECTS) $(LDFLAGS) -o  $@ 

.cpp.o:
	$(CC) $(CFLAGS) $< -o $@ 




