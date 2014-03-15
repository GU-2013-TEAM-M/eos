#ifndef mainmanager_h__
#define mainmanager_h__

#if defined _WIN32 || defined _WIN64
#define WIN32_LEAN_AND_MEAN
#endif

#include <boost/thread.hpp>
#include <boost/chrono.hpp>
#include <boost/property_tree/ptree.hpp>
#include <boost/property_tree/json_parser.hpp>
#include <unordered_map>

#include "cpuwatcher.h"
#include "memwatcher.h"
#include "netwatcher.h"

#include "outclient.h"
#include "inserver.h"

#ifdef TARGET_OS_MAC
#include "cpumac.h"

#elif defined __linux__
#include "cpuprocstat.h"
#include "memnix.h"

#elif defined _WIN32 || defined _WIN64
#define WIN32_LEAN_AND_MEAN
#include "cpuwin.h"
#include "memwin.h"
#include "netwin.h"

#else
#error "unknown platform"
#endif

class daemonManager {
private:
	CPUWatcher * mainCPUMon;
	MemWatcher * mainMemMon;
	NETWatcher * mainNetMon;
	std::unordered_map<std::string,bool> watcherStatus;
	outClient * connToServer;
	serveToClient * connToClient;
	boost::thread * toClientThread, * toServerThread;
	boost::chrono::milliseconds refresh;
	std::string daemonID, os, totalRAM;
	bool run;

	void handleServerMessage(std::string);
public:
	daemonManager();
	~daemonManager();
	void loop();
	
};

#endif // mainmanager_h__