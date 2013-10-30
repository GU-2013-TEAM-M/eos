#ifndef cpumon_h__
#define cpumon_h__

#include <string>
#include <boost/thread.hpp>
#include <boost/chrono.hpp>

class CPUWatcher {
private:
	double usageThisInterval;
	boost::thread *loopThread;
protected:
	boost::chrono::milliseconds refresh;
	bool run;
public:
	CPUWatcher();
	~CPUWatcher();
	double getUsage();
	void start();
	void updateUsage(double);
	virtual void readLoop() = 0;
};

#endif // cpumon_h__