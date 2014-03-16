#ifndef netwin_h__
#define netwin_h__

#include <Pdh.h>
#include <tchar.h>

#include "netwatcher.h"

class NETWin : public NETWatcher{
private:
	static PDH_HQUERY netQuery;
	static PDH_HCOUNTER netTotal;

	void init();

public:
	NETWin();
	NETWin(unsigned int);
	NETWin( boost::chrono::milliseconds );
	~NETWin();
	void readLoop();
};
#endif // netwin_h__