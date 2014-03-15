#ifndef memwin_h__
#define memwin_h__

#include "memwatcher.h"

#include <windows.h>
#include <tchar.h>

class MemWin : public MemWatcher {

public:
	MemWin();
	MemWin(unsigned int);
	MemWin( boost::chrono::milliseconds );
	~MemWin();
	static unsigned long long getTotalRAM();
	void readLoop();
};

#endif // memwin_h__