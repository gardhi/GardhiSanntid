CC=gcc
CFLAGS=-c -Wall -std=gnu11
LDFLAGS= -lpthread
SOURCES=chat.c
OBJECTS=$(SOURCES:.c=.o)
EXECUTABLE=chat

all: $(SOURCES) $(EXECUTABLE)
	
$(EXECUTABLE): $(OBJECTS)
	$(CC)  $(OBJECTS) $(LDFLAGS) -o  $@ 

.cpp.o:
	$(CC) $(CFLAGS) $< -o $@ 




