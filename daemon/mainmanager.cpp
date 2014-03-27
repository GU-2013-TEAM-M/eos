#include "mainmanager.h"

daemonManager::daemonManager() {
	//Read settings file
	std::fstream settingsFile("config.txt",std::fstream::in);
	std::string serverString, nameString, passString, idString, refSString, refRTString;
	getline(settingsFile, serverString);
	getline(settingsFile, nameString);
	getline(settingsFile, passString);
	getline(settingsFile, idString);
	getline(settingsFile, refSString);
	getline(settingsFile, refRTString);
	settingsFile.close();
	refreshS = boost::chrono::milliseconds(stoull(refSString));
	refreshRT = boost::chrono::milliseconds(stoull(refRTString));
	daemonID = "12345";

	//Initialize watchers based on platform
#ifdef TARGET_OS_MAC
	mainCPUMon = new CPUMac(refreshRT);
	os = "mac";
#elif defined __linux__
	mainCPUMon = new CPUProcStat(refreshRT);
	mainMemMon = new MemNix(refreshRT);
	mainNetMon = new NETNix(refreshRT);
	totalRAM = std::to_string(MemNix::getTotalRAM());
	os = "linux";
#elif defined _WIN32 || defined _WIN64
	mainCPUMon = new CPUWin(refreshRT);
	mainMemMon = new MemWin(refreshRT);
	mainNetMon = new NETWin(refreshRT);
	totalRAM = std::to_string(MemWin::getTotalRAM());
	os = "win";
#else
#error "unknown platform"
#endif

	//Start server connection manager
	connToServer = new outClient();
	connToServer->init(serverString);
	toServerThread = new boost::thread(boost::bind(&outClient::run, connToServer));

	//Wait until the server connection is ready
	while (!connToServer->ready) {
		boost::this_thread::sleep_for(boost::chrono::milliseconds(1));
	}

	//Bind callback message handler
	connToServer->bindMsg(boost::bind(&daemonManager::handleServerMessage,this,::_1));
	//Send daemon login message to server
	connToServer->send("{\"type\":\"login\",\"data\":{\"name\":\""+nameString+"\",\"password\":\""+passString+"\",\"org_id\":\""+idString+"\"}}");
	//Start client connection manager
	connToClient = new serveToClient();
	toClientThread = new boost::thread(boost::bind(&serveToClient::run, connToClient));

	//Start all watchers and set them to active
	mainCPUMon->start();
	watcherStatus["CPU"] = true;
	mainMemMon->start();
	watcherStatus["RAM"] = true;
	mainNetMon->start();
	watcherStatus["NET"] = true;
	run = true;
	
	serverLoop = new boost::thread(boost::bind(&daemonManager::loopS, this));
	clientLoop = new boost::thread(boost::bind(&daemonManager::loopRT, this));

	clientLoop->join();
	serverLoop->join();
}

void daemonManager::loopS() {
	std::string timestamp;
	std::vector<std::string> serverPayloads;
	std::string serverPayloadBase(
		(std::string) "{"							+
		"\"type\": \"monitoring\","	+
		"\"data\": {"				+
		"\"list\": [{"			+
		"\"parameter\": "	);
	while (run) {
		boost::this_thread::sleep_for(refreshS);
		timestamp = std::to_string(std::time(0));

		if (watcherStatus["CPU"]) {
			//Create CPU usage payload for the server
			std::string cpuP(serverPayloadBase + "\"cpu\", "+ 
								"\"values\": {" +
								"\"" + timestamp +"\": " + std::to_string(mainCPUMon->getUsage()) +
								"} }] } }");
			serverPayloads.push_back(cpuP);
		}

		if (watcherStatus["RAM"]) {
			//Create RAM usage payload for the server
			std::string ramP(serverPayloadBase + "\"ram\", "+ 
				"\"values\": {" +
				"\"" + timestamp +"\": " + std::to_string(mainMemMon->getUsage()) +
				"} }] } }");
			serverPayloads.push_back(ramP);
		}

		if (watcherStatus["NET"]) {
			//Create network usage payload for the server
			std::string netP(serverPayloadBase + "\"net\", "+ 
				"\"values\": {" +
				"\"" + timestamp +"\": " + std::to_string(mainNetMon->getUsage()) +
				"} }] } }");
			serverPayloads.push_back(netP);
		}

		//Send all data payloads to the server
		for (std::string payload : serverPayloads) {
			connToServer->send(payload);
		}
		serverPayloads.clear();
	}
}

void daemonManager::loopRT() {
	while (run) {
		std::string payloadBase("{\"type\": \"monitoring\", \"data\": {\"daemon_id\": \""+daemonID+"\", \"data\": {");
		boost::this_thread::sleep_for(refreshRT);

		std::string usagePayload = payloadBase;
		if (watcherStatus["CPU"]) {
			//Append CPU usage to payload
			usagePayload.append("\"cpu\": \"");
			usagePayload.append(std::to_string(mainCPUMon->getUsage()));
			usagePayload.append("\",");
		}

		if (watcherStatus["RAM"]) {
			//Append RAM usage to payload
			usagePayload.append("\"ram\": \"");
			usagePayload.append(std::to_string(mainMemMon->getUsage()));
			usagePayload.append("\",");
		}

		if (watcherStatus["NET"]) {
			//Append network usage to payload
			usagePayload.append("\"net\": \"");
			usagePayload.append(std::to_string(mainNetMon->getUsage()));
			usagePayload.append("\",");
		}
		//Remove final ","
		usagePayload = usagePayload.substr(0,usagePayload.length()-1);
		usagePayload.append("}}}");

		//Send combined data to all clients
		if (connToClient->open) {
			connToClient->send(usagePayload);
		}
	}
}

void daemonManager::handleServerMessage(std::string msg) {
	std::cout<<"Received message on handler"<<std::endl;
	//boost::property_tree::ptree comps;
	if (msg.find("operation")!=msg.npos) {
		bool setTo = false;
		//if start is found, the server wants us to start one or more watchers
		//	so we will set their status to true
		if ((msg.find("start")!=msg.npos)) {
			setTo = true;
		}
		//if CPU is found, set CPU watcher status
		if ((msg.find("CPU")!=msg.npos)) {
			watcherStatus["CPU"] = setTo;
		}
		//if RAM is found, set RAM watcher status
		if ((msg.find("RAM")!=msg.npos)) {
			watcherStatus["RAM"] = setTo;
		}
		//if NET is found, set NET watcher status
		if ((msg.find("NET")!=msg.npos)) {
			watcherStatus["NET"] = setTo;
		}
	} else if (msg.find("\"id\"")!=msg.npos) {
		//Find given daemon ID in response
		int spos = msg.find("\"id\"");
		spos = msg.find("\"", spos+5);
		daemonID = msg.substr(spos+1,msg.find("\"",spos+1)-(spos+1));
		//Send daemon info after receiving ID
		std::string identString("{\"type\":\"daemon\",\"data\":{\"daemon_id\":\""+daemonID+
			"\",\"daemon_platform\": \""+os+"\","+
			"\"daemon_all_parameters\":[\"cpu\",\"ram\",\"net\"],"+
			"\"daemon_monitored_parameters\":[\"cpu\",\"ram\",\"net\"]," +
			"\"ram_total\": \""+totalRAM+"\"}}");
		connToServer->send(identString);
	}
}
