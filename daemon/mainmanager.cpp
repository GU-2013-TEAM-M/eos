#include "mainmanager.h"

daemonManager::daemonManager() {
	refresh = boost::chrono::milliseconds(250);

#ifdef TARGET_OS_MAC
	mainCPUMon = new CPUMac(refresh);
#elif defined __linux__
	mainCPUMon = new CPUProcStat(refresh);
	mainMemMon = new MemNix(refresh);
#elif defined _WIN32 || defined _WIN64
	mainCPUMon = new CPUWin(refresh);
	mainMemMon = new MemWin(refresh);
	mainNetMon = new NETWin(refresh);
#else
#error "unknown platform"
#endif

	connToServer = new outClient();
	connToServer->init("ws://shacron.twilightparadox.com:8080/wsdaemon");
	connToClient = new serveToClient();
	toClientThread = new boost::thread(boost::bind(&serveToClient::run, connToClient));
	mainCPUMon->start();
	mainMemMon->start();
	mainNetMon->start();
	run = true;
	loop();
}

void daemonManager::loop() {
	std::string sendCPUbase("{\"type\": \"monitoring\", \"data\": {\"daemon_id\": \"12345\", \"data\": {\"cpu\": [");
	while (run) {
		boost::this_thread::sleep_for(refresh);
		std::string usagePayload = sendCPUbase;
		std::cout<<mainNetMon->getUsage()<<std::endl;
		usagePayload.append(std::to_string(mainCPUMon->getUsage()));
		usagePayload.append("]}}}");
		//connToServer->send(usagePayload);
		if (connToClient->open) {
			connToClient->send(usagePayload);
		}
	}
}

