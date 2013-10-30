#include "cpuwin.h"

PDH_HQUERY CPUWin::cpuQuery;
PDH_HCOUNTER CPUWin::cpuTotal;

CPUWin::CPUWin() {
	init();
	refresh = boost::chrono::milliseconds(500);
}

CPUWin::CPUWin( unsigned int ref ) {
	init();
	refresh = boost::chrono::milliseconds(ref);
}

CPUWin::CPUWin( boost::chrono::milliseconds ref ) {
	init();
	refresh = ref;
}

CPUWin::~CPUWin() {

}

void CPUWin::init() {
	PdhOpenQuery(NULL, NULL, &cpuQuery);
	PdhAddCounter(cpuQuery, L"\\Processor(_Total)\\% Processor Time", NULL, &cpuTotal);
	PdhCollectQueryData(cpuQuery);
}

void CPUWin::readLoop() {
	while (run) {
		PDH_FMT_COUNTERVALUE counterVal;
		PdhCollectQueryData(cpuQuery);
		PdhGetFormattedCounterValue(cpuTotal, PDH_FMT_DOUBLE, NULL, &counterVal);
		updateUsage(counterVal.doubleValue);
		boost::this_thread::sleep_for(refresh);
	}
}