#ifndef cpucycles_h__
#define cpucycles_h__


class CPUCycles {
private:
	long long unsigned int totalCycles, workCycles;
public:
	CPUCycles();
	CPUCycles & operator= (const CPUCycles &other);
	void reset();
	long long unsigned int getWorkCycles() const;
	void setWorkCycles(long long unsigned int val);
	long long unsigned int getTotalCycles() const;
	void setTotalCycles(long long unsigned int val);
};

#endif // cpucycles_h__