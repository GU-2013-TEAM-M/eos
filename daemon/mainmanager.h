#ifndef mainmanager_h__
#define mainmanager_h__

#include <boost/thread.hpp>
#include <boost/chrono.hpp>

#include "cpuwatcher.h"
#include "outclient.h"

#ifdef TARGET_OS_MAC
#include "cpumac.h"
#elif defined __linux__
#include "cpuprocstat.h"
#elif defined _WIN32 || defined _WIN64
#include "cpuwin.h"
#else
#error "unknown platform"
#endif

class daemonManager {
private:
	CPUWatcher * mainCPUMon;
	outClient * connToServer;
	boost::chrono::milliseconds refresh;
	bool run;
public:
	daemonManager();
	~daemonManager();
	void loop();
};

#endif // mainmanager_h__