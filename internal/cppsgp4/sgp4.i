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
    #include "SolarPosition.h"
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
    std::string Name() const;
    std::string Line1() const;
    std::string Line2() const;
    unsigned int NoradNumber() const;
    std::string IntDesignator() const;
    DateTime Epoch() const;
    double MeanMotionDt2() const;
    double MeanMotionDdt6() const;
    double BStar() const;
    double Inclination(bool in_degrees) const;
    double RightAscendingNode(const bool in_degrees) const;
    double Eccentricity() const;
    double ArgumentPerigee(const bool in_degrees) const;
    double MeanAnomaly(const bool in_degrees) const;
    double MeanMotion() const;
    unsigned int OrbitNumber() const;
    std::string ToString() const;
};

class Observer {
public:
    Observer(const CoordGeodetic &geo);
    CoordGeodetic GetLocation() const;
    CoordTopocentric GetLookAngle(const Eci &eci);
};

class Eci {
public:
    Eci(const DateTime& dt, const CoordGeodetic& geo);
    Vector Velocity() const;
    DateTime GetDateTime() const;
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
    std::string ToString() const;
};

struct CoordTopocentric {
public:
    double azimuth;
    double elevation;
    double range;
    double range_rate;
    std::string ToString() const;
};

class SolarPosition {
public:
    SolarPosition();
    Eci FindPosition(const DateTime& dt);
};
