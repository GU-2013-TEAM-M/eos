#include "netwatcher.h"

NETWatcher::NETWatcher() {

}

NETWatcher::~NETWatcher() {

}

void NETWatcher::start() {
	run = true;
	loopThread = new boost::thread(boost::bind(&NETWatcher::readLoop, this));
}

long long NETWatcher::getUsage() {
	return usageThisInterval;
}

void NETWatcher::updateUsage(long long usage) {
	usageThisInterval = usage;
}


