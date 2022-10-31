#ifndef CUSTOM_HELPERS_H
#define CUSTOM_HELPERS_H

#include <DateTime.h>
#include <SGP4.h>
#include <CoordGeodetic.h>

#include <cmath>
#include <iostream>
#include <vector>

struct PassDetails
{
    DateTime aos;
    DateTime los;
    double max_elevation;
};

double FindMaxElevation(const CoordGeodetic &user_geo, SGP4 &sgp4, const DateTime &aos, const DateTime &los);
DateTime FindCrossingPoint(const CoordGeodetic &user_geo, SGP4 &sgp4, const DateTime &initial_time1, const DateTime &initial_time2, bool finding_aos);
std::vector<struct PassDetails> GeneratePassList(const CoordGeodetic &user_geo, SGP4 &sgp4, const DateTime &start_time, const DateTime &end_time, const int time_step);

#endif
