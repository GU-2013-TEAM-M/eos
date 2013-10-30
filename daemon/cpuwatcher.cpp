#include "cpuwatcher.h"

CPUWatcher::CPUWatcher() {
	
}

CPUWatcher::~CPUWatcher() {

}

void CPUWatcher::start() {
	run = true;
	loopThread = new boost::thread(boost::bind(&CPUWatcher::readLoop, this));
}

double CPUWatcher::getUsage() {
	return usageThisInterval;
}

void CPUWatcher::updateUsage(double usage) {
	usageThisInterval = usage;
}


