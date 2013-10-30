#ifndef cpuprocstat_h__
#define cpuprocstat_h__

#include <iostream>
#include <iomanip>
#include <fstream>
#include <string>
#include <sstream>
#include <vector>
#include <boost/thread.hpp>
#include <chrono>
#include <stdexcept>
#include "cpucycles.h"

#include "cpuwatcher.h"

class CPUProcStat : public CPUWatcher{
private:
	CPUCycles *procNumsFirst, *procNumsSecond;
	const std::string PROC;

	CPUCycles * procCalcTotal(std::vector<long long unsigned int>);
	void procParse(std::string);
	

public:
	CPUProcStat();
	CPUProcStat(unsigned int);
	CPUProcStat( boost::chrono::milliseconds );
	~CPUProcStat();
	void readLoop();
	
};

#endif // cpuprocstat_h__