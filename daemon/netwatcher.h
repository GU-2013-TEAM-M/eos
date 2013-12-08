#ifndef netwatcher_h__
#define netwatcher_h__

#include <string>
#include <boost/thread.hpp>
#include <boost/chrono.hpp>

class NETWatcher {
private:
	long long usageThisInterval;
	boost::thread *loopThread;
protected:
	boost::chrono::milliseconds refresh;
	bool run;
public:
	NETWatcher();
	~NETWatcher();
	long long getUsage();
	void start();
	void updateUsage(long long);
	virtual void readLoop() = 0;
};
#endif // netwatcher_h__