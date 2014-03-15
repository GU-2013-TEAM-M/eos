#include "mainmanager.h"

daemonManager::daemonManager() {
	refresh = boost::chrono::milliseconds(1000);
	daemonID = 12345;

#ifdef TARGET_OS_MAC
	mainCPUMon = new CPUMac(refresh);
	os = "mac"
#elif defined __linux__
	mainCPUMon = new CPUProcStat(refresh);
	mainMemMon = new MemNix(refresh);
	mainNetMon = new NETNix(refresh);
	os = "linux"
#elif defined _WIN32 || defined _WIN64
	mainCPUMon = new CPUWin(refresh);
	mainMemMon = new MemWin(refresh);
	mainNetMon = new NETWin(refresh);
	os = "win";
#else
#error "unknown platform"
#endif

	connToServer = new outClient();
	connToServer->init("ws://eos.sytes.net/wsdaemon");
	toServerThread = new boost::thread(boost::bind(&outClient::run, connToServer));

	while (!connToServer->ready) {
		boost::this_thread::sleep_for(boost::chrono::milliseconds(1));
	}
	connToServer->bindMsg(boost::bind(&daemonManager::handleServerMessage,this,::_1));
	connToServer->send("{\"type\":\"login\",\"data\":{\"name\":\"Anekharat\",\"password\":\"derp\",\"org_id\":\"52f2bf8521db4a04d5000001\"}}");
	connToClient = new serveToClient();
	toClientThread = new boost::thread(boost::bind(&serveToClient::run, connToClient));

	mainCPUMon->start();
	watcherStatus["CPU"] = true;
	mainMemMon->start();
	watcherStatus["RAM"] = true;
	mainNetMon->start();
	watcherStatus["NET"] = true;
	run = true;
	loop();
}

void daemonManager::loop() {
	while (run) {
		std::string sendCPUbase("{\"type\": \"monitoring\", \"data\": {\"daemon_id\": \""+daemonID+"\", \"data\": {");
		boost::this_thread::sleep_for(refresh);
		std::string usagePayload = sendCPUbase;
		if (watcherStatus["CPU"]) {
			usagePayload.append("\"cpu\": \"");
			usagePayload.append(std::to_string(mainCPUMon->getUsage()));
			usagePayload.append("\",");
		}
		if (watcherStatus["RAM"]) {
			usagePayload.append("\"ram\": \"");
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
	//	std::cout<<usagePayload<<std::endl;
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
		if ((msg.find("RAM")!=msg.npos)) {
			watcherStatus["RAM"] = setTo;
		}
		//if (comps.get<std::string>(op)=="NET") {
		if ((msg.find("NET")!=msg.npos)) {
			watcherStatus["NET"] = setTo;
		}
	} else if (msg.find("\"id\"")!=msg.npos) {
		int spos = msg.find("\"id\"");
		spos = msg.find("\"", spos+5);
		daemonID = msg.substr(spos+1,msg.find("\"",spos+1)-(spos+1));
		//Send daemon info after receiving ID
		std::string identString = "{\"type\":\"daemon\",\"data\":{\"daemon_id\":\""+daemonID+
			"\",\"daemon_platform\":{os: \""+os+"\", ram_total: \""+std::to_string(mainMemMon->getTotalRAM()) +"\"},"+
			"\"daemon_all_parameters\":[\"cpu\",\"ram\",\"net\"],"+
			"\"daemon_monitored_parameters\":[\"cpu\",\"ram\",\"net\"]}}";
		connToServer->send(identString);
	//	daemonID = stoll(idString);
	}
}