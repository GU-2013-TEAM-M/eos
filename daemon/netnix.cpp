#include "netnix.h"

NETNix::NETNix() {
	refresh = boost::chrono::milliseconds(500);
}

NETNix::NETNix( unsigned int ref ) {
	refresh = boost::chrono::milliseconds(ref);
}

NETNix::NETNix( boost::chrono::milliseconds ref ) {
	refresh = ref;
}

NETNix::~NETNix() {

}

void NETNix::readLoop() {
	while (run) {
		updateUsage(9001);
		boost::this_thread::sleep_for(refresh);
	}
}