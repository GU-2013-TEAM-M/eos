#ifndef memwin_h__
#define memwin_h__

#include <windows.h>
#include <tchar.h>

#include "memwatcher.h"

class MemWin : public MemWatcher {

public:
	MemWin();
	MemWin(unsigned int);
	MemWin( boost::chrono::milliseconds );
	~MemWin();
	void readLoop();
};

#endif // memwin_h__