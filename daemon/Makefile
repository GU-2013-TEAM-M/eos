SOURCES = protoDaemon.cpp mainmanager.cpp outclient.cpp cpuwatcher.cpp cpuwin.cpp cpuprocstat.cpp cpucycles.cpp
OBJECTS=$(SOURCES:.cpp=.o)
CC = g++
DEBUG = -g -W -Wall
CFLAGS = -c -std=c++11 $(DEBUG)
WSPPI = -I/usr/include/c++/websocketpp
BOOSTI = -lboost_system-mt -lboost_thread-mt
LFLAGS = $(DEBUG)
EXEC = daemon

all: $(SOURCES) $(EXEC)
	
$(EXEC): $(OBJECTS) 
	$(CC) $(LDFLAGS) $(OBJECTS) $(BOOSTI) -o $@

.cpp.o:
	$(CC) $(CFLAGS) $(WSPPI) $(BOOSTI) $< -o $@

clean:
	\rm *.o *~ $(EXEC)