#include "memwatcher.h"

MemWatcher::MemWatcher() {

}

MemWatcher::~MemWatcher() {

}

void MemWatcher::start() {
	run = true;
	loopThread = new boost::thread(boost::bind(&MemWatcher::readLoop, this));
}

unsigned long long int MemWatcher::getUsage() {
	return usageThisInterval;
}

void MemWatcher::updateUsage(unsigned long long int usage) {
	usageThisInterval = usage;
}


