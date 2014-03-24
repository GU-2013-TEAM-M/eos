SOURCES = protoDaemon.cpp mainmanager.cpp toServerConn.cpp toClientsConn.cpp cpuwatcher.cpp cpucycles.cpp cpuprocstat.cpp memwatcher.cpp memnix.cpp netwatcher.cpp netnix.cpp
OBJECTS=$(SOURCES:.cpp=.o)
CC = g++
DEBUG = -g -W -Wall
CFLAGS = -c -std=c++11
WSPPI = -D_WEBSOCKETPP_CPP11_STL_ -I/usr/include/c++/websocketpp
BOOSTI = -lboost_system -lboost_thread -pthread
LFLAGS = $(DEBUG)
EXEC = daemon

all: $(SOURCES) $(EXEC)
	
$(EXEC): $(OBJECTS) 
	$(CC) $(OBJECTS) $(BOOSTI) -o $@

.cpp.o:
	$(CC) $(CFLAGS) $(WSPPI) $(BOOSTI) $< -o $@

clean:
	\rm *.o *~ $(EXEC)