#include "cpucycles.h"

CPUCycles::CPUCycles() {
	totalCycles = 0;
	workCycles = 0;
}

CPUCycles & CPUCycles::operator=( const CPUCycles &other ) {
	if (this != &other) {
		totalCycles = other.totalCycles;
		workCycles = other.workCycles;
	}
	return *this;
}

void CPUCycles::reset() {
	totalCycles = 0;
	workCycles = 0;
}

long long unsigned int CPUCycles::getWorkCycles() const {
	return workCycles;
}

void CPUCycles::setWorkCycles( long long unsigned int val ) {
	workCycles = val;
}

long long unsigned int CPUCycles::getTotalCycles() const {
	return totalCycles;
}

void CPUCycles::setTotalCycles( long long unsigned int val ) {
	totalCycles = val;
}
