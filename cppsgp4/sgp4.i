%module cppsgp4

%{
    #include "SGP4.h"
    #include "Observer.h"
    #include "CoordTopocentric.h"
    #include "CoordGeodetic.h"
    #include "OrbitalElements.h"
    #include "Eci.h"
    #include "Tle.h"
    #include "TleException.h"
    #include "DateTime.h"
    #include "Helpers.h"
%}

%include <typemaps.i>
%include "std_string.i"
%include "std_vector.i"
%include "exception.i"

%exception {
    try {
        $action;
    } catch (std::runtime_error &e) {
        _swig_gopanic(e.what());
    }
}

class SGP4 {
public:
    SGP4(const Tle& tle);
    Eci FindPosition(const DateTime& date) const;
};

class Tle {
public:
    Tle(const std::string& name, const std::string& line_one, const std::string& line_two);
    std::string Line1() const;
    std::string Line2() const;
    std::string Name() const;
    unsigned int NoradNumber() const;
    DateTime Epoch() const;
};

class Observer {
public:
    Observer(const CoordGeodetic &geo);
    CoordTopocentric GetLookAngle(const Eci &eci);
    CoordGeodetic GetLocation() const;
};

class Eci {
public:
    Eci(const DateTime& dt, const CoordGeodetic& geo);
    CoordGeodetic ToGeodetic() const;
};

class DateTime {
public:
    DateTime(int year, int month, int day, int hour, int minute, int second);
    static DateTime Now(bool useMicroseconds = false);
    double ToJulian() const;
};

struct CoordGeodetic {
public:
    double latitude;
    double longitude;
    double altitude;
};

struct CoordTopocentric {
public:
    double azimuth;
    double elevation;
    double range;
    double range_rate;
};

// Helpers part

namespace std {
   %template(PassDetailsVector) vector<PassDetails>;
}

struct PassDetails
{
public:
    DateTime aos;
    DateTime los;
    double max_elevation;
};

std::vector<struct PassDetails> GeneratePassList(const CoordGeodetic &user_geo, SGP4 &sgp4, const DateTime &start_time, const DateTime &end_time, const int time_step);
