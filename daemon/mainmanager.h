#ifndef mainmanager_h__
#define mainmanager_h__

#include <boost/thread.hpp>
#include <boost/chrono.hpp>

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

	outClient * connToServer;
	serveToClient * connToClient;
	boost::thread * toClientThread;
	boost::chrono::milliseconds refresh;
	bool run;
public:
	daemonManager();
	~daemonManager();
	void loop();
};

#endif // mainmanager_h__