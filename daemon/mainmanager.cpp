#include "mainmanager.h"

daemonManager::daemonManager() {
	refresh = boost::chrono::milliseconds(250);

#ifdef TARGET_OS_MAC
	mainCPUMon = new CPUMac(refresh);
#elif defined __linux__
	mainCPUMon = new CPUProcStat(refresh);
#elif defined _WIN32 || defined _WIN64
	mainCPUMon = new CPUWin(refresh);
#else
#error "unknown platform"
#endif

//	connToServer = new outClient;
//	connToServer->init("ws://localhost:8080");
	mainCPUMon->start();
	run = true;
	loop();
}

void daemonManager::loop() {
	while (run) {
		boost::this_thread::sleep_for(refresh);
		std::string usagePayload;
		std::cout<<mainCPUMon->getUsage()<<std::endl;
		//connToServer->send("");
	}
}

