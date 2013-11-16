#ifndef cpuwin_h__
#define cpuwin_h__

#include <Pdh.h>
#include <tchar.h>

#include "cpuwatcher.h"

class CPUWin : public CPUWatcher{
private:
	static PDH_HQUERY cpuQuery;
	static PDH_HCOUNTER cpuTotal;

	void procParse(std::string);
	void init();

public:
	CPUWin();
	CPUWin(unsigned int);
	CPUWin( boost::chrono::milliseconds );
	~CPUWin();
	void readLoop();
};

#endif // cpuwin_h__