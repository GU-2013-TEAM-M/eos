#ifndef memwatcher_h__
#define memwatcher_h__

#include <string>
#include <boost/thread.hpp>
#include <boost/chrono.hpp>

class MemWatcher {
private:
	unsigned long long int usageThisInterval;
	boost::thread *loopThread;
protected:	
	boost::chrono::milliseconds refresh;
	bool run;
public:
	MemWatcher();
	~MemWatcher();
	unsigned long long int getUsage();
	void start();
	void updateUsage(unsigned long long int);
	static unsigned long long getTotalRAM();
	virtual void readLoop() = 0;
};

#endif // memwatcher_h__