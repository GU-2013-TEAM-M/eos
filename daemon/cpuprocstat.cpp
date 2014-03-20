#include "cpuprocstat.h"

CPUProcStat::CPUProcStat() : PROC ("/proc/stat") {
	procNumsFirst = new CPUCycles();
	procNumsSecond = new CPUCycles();
}

CPUProcStat::CPUProcStat( unsigned int ref ) : PROC ("/proc/stat") {
	procNumsFirst = new CPUCycles();
	procNumsSecond = new CPUCycles();
	refresh = boost::chrono::milliseconds(ref);
}

CPUProcStat::CPUProcStat( boost::chrono::milliseconds ref ) : PROC ("/proc/stat") {
	procNumsFirst = new CPUCycles();
	procNumsSecond = new CPUCycles();
	refresh = ref;
}

CPUProcStat::~CPUProcStat() {
	delete procNumsFirst;
	delete procNumsSecond;
}

CPUCycles * CPUProcStat::procCalcTotal(std::vector<long long unsigned int> cycles) {
	CPUCycles *current = new CPUCycles();
	//cout<<"Size: "<<cycles.size()<<endl;
	for (unsigned int currCycleIndex = 0; currCycleIndex < cycles.size(); currCycleIndex++) {
		if (currCycleIndex<3) {
			current->setWorkCycles(current->getWorkCycles()+cycles.at(currCycleIndex));
		}
		current->setTotalCycles(current->getTotalCycles()+cycles.at(currCycleIndex));
		//cout<<"Current assigned cycle num: "<<cycles.at(currCycleIndex)<<endl;
	}
	//cout<<"Current calculation: "<<current.workCycles<<" "<<current.totalCycles<<endl;
	return current;
}

void CPUProcStat::procParse(std::string procLine) {
	std::vector<long long unsigned int> procCycles;
	std::stringstream procLineStream;
	std::string currCycles;
	procLine = procLine.substr(procLine.find(' ')+1);
	procLineStream.str(procLine);
	//cout<<procLine<<endl;
	while (getline(procLineStream,currCycles,' ')) {
		//cout<<"Current cycles to intify: "<<currCycles<<endl;
		long long unsigned int currCyclesLLUI;
		try {
			currCyclesLLUI = stoull(currCycles);
		} catch (std::invalid_argument e) {
	//		std::cout<<"Cycles string \""<<currCycles<<"\" invalid. Continuing."<<std::endl;
			continue;
		}
		//cout<<"Pushing "<<currCyclesLLUI<<"; current size "<<procCycles.size()<<endl;
		procCycles.push_back(currCyclesLLUI);
	}
	//cout<<"First cycles before check: "<<procNumsFirst.totalCycles<<endl;
	//cout<<"Size before check: "<<procCycles.size()<<endl;
	if (procNumsFirst->getTotalCycles() == 0) {
		//cout<<"calculating first cycles"<<endl;
		*procNumsFirst = *procCalcTotal(procCycles);
		//cout<<"First cycles after calc "<<procNumsFirst.totalCycles<<endl;
	} else if (procNumsSecond->getTotalCycles() == 0) {
		//cout<<"calculating second cycles"<<endl;
		*procNumsSecond = *procCalcTotal(procCycles);
		long long unsigned int work = procNumsSecond->getWorkCycles() - procNumsFirst->getWorkCycles();
		long long unsigned int total = procNumsSecond->getTotalCycles() - procNumsFirst->getTotalCycles();
		updateUsage((double)work/total*100.0);
	} else {
		*procNumsFirst = *procNumsSecond;
		*procNumsSecond = *procCalcTotal(procCycles);
		long long unsigned int work = procNumsSecond->getWorkCycles() - procNumsFirst->getWorkCycles();
		long long unsigned int total = procNumsSecond->getTotalCycles() - procNumsFirst->getTotalCycles();
		updateUsage((double)work/total*100.0);
	}
	//std::cout<<getUsage()<<std::endl;
}

void CPUProcStat::readLoop() {
	std::string procLine;
	std::fstream procFile;

	const bool DEBUG = false;

	while (run) {
		if (!DEBUG) {
			procFile.open(PROC,std::fstream::in);
			getline(procFile, procLine);
			procFile.close();
			//std::cout<<procLine<<std::endl;
			procParse(procLine);
			procLine.clear();
			boost::this_thread::sleep_for(refresh);
		} else {
			procFile.open(PROC,std::fstream::in);
			while (std::getline(procFile,procLine)) {
				std::cout<<"Proc line: "<<procLine<<std::endl;
				procParse(procLine);
				procLine.clear();
				boost::this_thread::sleep_for(refresh);
			}
			procFile.close();
			run=false;
		}
	}
}
