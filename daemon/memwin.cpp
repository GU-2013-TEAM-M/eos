#include "memwin.h"

MemWin::MemWin() {
	// initialize with default (0.5s) refresh rate
	refresh = boost::chrono::milliseconds(500);
}


MemWin::MemWin( unsigned int ref ) {
	refresh = boost::chrono::milliseconds(ref);
}

MemWin::MemWin( boost::chrono::milliseconds ref ) {
	refresh = ref;
}

MemWin::~MemWin() {

}

void MemWin::readLoop() {
	while (run) {
		MEMORYSTATUSEX memInfo;
		memInfo.dwLength = sizeof(MEMORYSTATUSEX);
		GlobalMemoryStatusEx(&memInfo);
		updateUsage(memInfo.ullTotalPhys - memInfo.ullAvailPhys);
		boost::this_thread::sleep_for(refresh);
	}
}

unsigned long long MemWin::getTotalRAM() {
	MEMORYSTATUSEX memInfo;
	memInfo.dwLength = sizeof(MEMORYSTATUSEX);
	GlobalMemoryStatusEx(&memInfo);
	return memInfo.ullTotalPhys;
}