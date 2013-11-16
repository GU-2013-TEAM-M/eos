#include "memnix.h"

MemNix::MemNix() {
	// initialize with default (0.5s) refresh rate
	refresh = boost::chrono::milliseconds(500);
}


MemNix::MemNix( unsigned int ref ) {
	refresh = boost::chrono::milliseconds(ref);
}

MemNix::MemNix( boost::chrono::milliseconds ref ) {
	refresh = ref;
}

MemNix::~MemNix() {

}

void MemNix::readLoop() {
	struct sysinfo memInfo;
	while (run) {
		sysinfo (&memInfo);
		unsigned long long int physMemUsed = memInfo.totalram - memInfo.freeram;
		updateUsage(physMemUsed*memInfo.mem_unit);
		boost::this_thread::sleep_for(refresh);
	}
}