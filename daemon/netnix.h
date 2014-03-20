#ifndef netnix_h__
#define netnix_h__

#include "netwatcher.h"

class NETNix : public NETWatcher{
private:
	void procParse(std::string);

public:
	NETNix();
	NETNix(unsigned int);
	NETNix( boost::chrono::milliseconds );
	~NETNix();
	void readLoop();
};
#endif // netnix_h__