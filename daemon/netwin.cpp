#include "netwin.h"

PDH_HQUERY NETWin::netQuery;
PDH_HCOUNTER NETWin::netTotal;

NETWin::NETWin() {
	init();
	refresh = boost::chrono::milliseconds(500);
}

NETWin::NETWin( unsigned int ref ) {
	init();
	refresh = boost::chrono::milliseconds(ref);
}

NETWin::NETWin( boost::chrono::milliseconds ref ) {
	init();
	refresh = ref;
}

NETWin::~NETWin() {

}

void NETWin::init() {
	PdhOpenQuery(NULL, NULL, &netQuery);
	PDH_STATUS pdhStatus;
	DWORD pathSize = 90000;
	LPTSTR buffer = new TCHAR[90000];
	pdhStatus = PdhExpandWildCardPath(NULL,L"\\Network Interface(*)\\Bytes Received/sec",buffer,&pathSize,NULL);
	std::cout<<GetLastError();
	std::wstring test;
	test.append(buffer);
	std::wcout<<test.c_str();
	std::cout<<PdhAddCounter(netQuery, buffer, NULL, &netTotal)<<std::endl;
	PdhCollectQueryData(netQuery);
}

void NETWin::readLoop() {
	while (run) {
		PDH_FMT_COUNTERVALUE counterVal;
		PdhCollectQueryData(netQuery);
		PdhGetFormattedCounterValue(netTotal, PDH_FMT_LARGE, NULL, &counterVal);
		updateUsage(counterVal.largeValue);
		boost::this_thread::sleep_for(refresh);
	}
}