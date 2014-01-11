#include "mainmanager.h"

daemonManager::daemonManager() {
	refresh = boost::chrono::milliseconds(250);
	daemonID = 12345;
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
	toServerThread = new boost::thread(boost::bind(&outClient::run, connToServer));

	while (!connToServer->ready) {
		boost::this_thread::sleep_for(boost::chrono::milliseconds(1));
	}
	connToServer->bindMsg(boost::bind(&daemonManager::handleServerMessage,this,::_1));
	connToServer->send("{\"type\":\"login\", \"data\":{}}");

	connToClient = new serveToClient();
	toClientThread = new boost::thread(boost::bind(&serveToClient::run, connToClient));

	mainCPUMon->start();
	watcherStatus["CPU"] = true;
	mainMemMon->start();
	watcherStatus["MEM"] = true;
	mainNetMon->start();
	watcherStatus["NET"] = true;
	run = true;
	loop();
}

void daemonManager::loop() {
	while (run) {
		std::string sendCPUbase("{\"type\": \"monitoring\", \"data\": {\"daemon_id\": \""+std::to_string(daemonID)+"\", \"data\": {");
		boost::this_thread::sleep_for(refresh);
		std::string usagePayload = sendCPUbase;
		if (watcherStatus["CPU"]) {
			usagePayload.append("\"cpu\": \"");
			usagePayload.append(std::to_string(mainCPUMon->getUsage()));
			usagePayload.append("\",");
		}
		if (watcherStatus["MEM"]) {
			usagePayload.append("\"mem\": \"");
			usagePayload.append(std::to_string(mainMemMon->getUsage()));
			usagePayload.append("\",");
		}
		if (watcherStatus["NET"]) {
			usagePayload.append("\"net\": \"");
			usagePayload.append(std::to_string(mainNetMon->getUsage()));
			usagePayload.append("\",");
		}
		usagePayload = usagePayload.substr(0,usagePayload.length()-1);
		usagePayload.append("}}}");
		std::cout<<usagePayload<<std::endl;
		//connToServer->send(usagePayload);
		if (connToClient->open) {
			connToClient->send(usagePayload);
		}
	}
}

void daemonManager::handleServerMessage(std::string msg) {
	std::cout<<"Received message on handler"<<std::endl;
	//boost::property_tree::ptree comps;
	if (msg.find("operation")!=msg.npos) {
		//std::istringstream inputBuffer(msg.c_str());
		//boost::property_tree::read_json(inputBuffer, comps);
		bool setTo = false;
		//std::string op = comps.get<std::string>("operation");
		if ((msg.find("start")!=msg.npos)) {
			setTo = true;
		}
		//if (comps.get<std::string>(op)=="CPU") {
		if ((msg.find("CPU")!=msg.npos)) {
			watcherStatus["CPU"] = setTo;
		}
		//if (comps.get<std::string>(op)=="RAM") {
		if ((msg.find("MEM")!=msg.npos)) {
			watcherStatus["MEM"] = setTo;
		}
		//if (comps.get<std::string>(op)=="NET") {
		if ((msg.find("NET")!=msg.npos)) {
			watcherStatus["NET"] = setTo;
		}
	} else if (msg.find("\"id\"")!=msg.npos) {
		int spos = msg.find("\"id\"");
		spos = msg.find("\"", spos+5);
		std::string idString = msg.substr(spos+1,msg.find("\"",spos+1)-(spos+1));
		daemonID = stoll(idString);
	}
}