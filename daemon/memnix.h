#ifndef memnix_h__
#define memnix_h__

#include "sys/types.h"
#include "sys/sysinfo.h"

#include "memwatcher.h"

class MemNix : public MemWatcher {

public:
	MemNix();
	MemNix(unsigned int);
	MemNix( boost::chrono::milliseconds );
	~MemNix();
	void readLoop();
};

#endif // memnix_h__