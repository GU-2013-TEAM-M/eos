#include "netnix.h"

PDH_HQUERY NETWin::netQuery;
PDH_HCOUNTER NETWin::netTotal;

NETNix::NETNix() {
	init();
	refresh = boost::chrono::milliseconds(500);
}

NETNix::NETNix( unsigned int ref ) {
	init();
	refresh = boost::chrono::milliseconds(ref);
}

NETNix::NETNix( boost::chrono::milliseconds ref ) {
	init();
	refresh = ref;
}

NETNix::~NETNix() {

}

void NETNix::init() {

}

void NETNix::readLoop() {
	updateUsage((long long int) 9001);
}