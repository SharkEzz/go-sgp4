/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 4.0.2
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: internal/cppsgp4/sgp4.i

package cppsgp4

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef _gostring_ swig_type_1;
typedef _gostring_ swig_type_2;
typedef _gostring_ swig_type_3;
typedef _gostring_ swig_type_4;
typedef _gostring_ swig_type_5;
typedef _gostring_ swig_type_6;
typedef _gostring_ swig_type_7;
typedef _gostring_ swig_type_8;
typedef _gostring_ swig_type_9;
typedef _gostring_ swig_type_10;
extern void _wrap_Swig_free_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_cppsgp4_5c90d950726e409b(swig_intgo arg1);
extern uintptr_t _wrap_new_SGP4_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern uintptr_t _wrap_SGP4_FindPosition_cppsgp4_5c90d950726e409b(uintptr_t arg1, uintptr_t arg2);
extern void _wrap_delete_SGP4_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern uintptr_t _wrap_new_Tle_cppsgp4_5c90d950726e409b(swig_type_1 arg1, swig_type_2 arg2, swig_type_3 arg3);
extern swig_type_4 _wrap_Tle_Name_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern swig_type_5 _wrap_Tle_Line1_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern swig_type_6 _wrap_Tle_Line2_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern swig_intgo _wrap_Tle_NoradNumber_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern swig_type_7 _wrap_Tle_IntDesignator_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern uintptr_t _wrap_Tle_Epoch_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern double _wrap_Tle_MeanMotionDt2_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern double _wrap_Tle_MeanMotionDdt6_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern double _wrap_Tle_BStar_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern double _wrap_Tle_Inclination_cppsgp4_5c90d950726e409b(uintptr_t arg1, _Bool arg2);
extern double _wrap_Tle_RightAscendingNode_cppsgp4_5c90d950726e409b(uintptr_t arg1, _Bool arg2);
extern double _wrap_Tle_Eccentricity_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern double _wrap_Tle_ArgumentPerigee_cppsgp4_5c90d950726e409b(uintptr_t arg1, _Bool arg2);
extern double _wrap_Tle_MeanAnomaly_cppsgp4_5c90d950726e409b(uintptr_t arg1, _Bool arg2);
extern double _wrap_Tle_MeanMotion_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern swig_intgo _wrap_Tle_OrbitNumber_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern swig_type_8 _wrap_Tle_ToString_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern void _wrap_delete_Tle_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern uintptr_t _wrap_new_Observer_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern uintptr_t _wrap_Observer_GetLocation_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern uintptr_t _wrap_Observer_GetLookAngle_cppsgp4_5c90d950726e409b(uintptr_t arg1, uintptr_t arg2);
extern void _wrap_delete_Observer_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern uintptr_t _wrap_new_Eci_cppsgp4_5c90d950726e409b(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_Eci_GetDateTime_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern uintptr_t _wrap_Eci_ToGeodetic_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern void _wrap_delete_Eci_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern uintptr_t _wrap_new_DateTime_cppsgp4_5c90d950726e409b(swig_intgo arg1, swig_intgo arg2, swig_intgo arg3, swig_intgo arg4, swig_intgo arg5, swig_intgo arg6);
extern uintptr_t _wrap_DateTime_Now__SWIG_0_cppsgp4_5c90d950726e409b(_Bool arg1);
extern uintptr_t _wrap_DateTime_Now__SWIG_1_cppsgp4_5c90d950726e409b(void);
extern double _wrap_DateTime_ToJulian_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern swig_intgo _wrap_DateTime_Year_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern swig_intgo _wrap_DateTime_Month_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern swig_intgo _wrap_DateTime_Day_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern swig_intgo _wrap_DateTime_Hour_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern swig_intgo _wrap_DateTime_Minute_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern swig_intgo _wrap_DateTime_Second_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern void _wrap_delete_DateTime_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern uintptr_t _wrap_new_CoordGeodetic__SWIG_0_cppsgp4_5c90d950726e409b(double arg1, double arg2, double arg3, _Bool arg4);
extern uintptr_t _wrap_new_CoordGeodetic__SWIG_1_cppsgp4_5c90d950726e409b(double arg1, double arg2, double arg3);
extern void _wrap_CoordGeodetic_latitude_set_cppsgp4_5c90d950726e409b(uintptr_t arg1, double arg2);
extern double _wrap_CoordGeodetic_latitude_get_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern void _wrap_CoordGeodetic_longitude_set_cppsgp4_5c90d950726e409b(uintptr_t arg1, double arg2);
extern double _wrap_CoordGeodetic_longitude_get_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern void _wrap_CoordGeodetic_altitude_set_cppsgp4_5c90d950726e409b(uintptr_t arg1, double arg2);
extern double _wrap_CoordGeodetic_altitude_get_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern swig_type_9 _wrap_CoordGeodetic_ToString_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern void _wrap_delete_CoordGeodetic_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern uintptr_t _wrap_new_CoordTopocentric_cppsgp4_5c90d950726e409b(double arg1, double arg2, double arg3, double arg4);
extern void _wrap_CoordTopocentric_azimuth_set_cppsgp4_5c90d950726e409b(uintptr_t arg1, double arg2);
extern double _wrap_CoordTopocentric_azimuth_get_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern void _wrap_CoordTopocentric_elevation_set_cppsgp4_5c90d950726e409b(uintptr_t arg1, double arg2);
extern double _wrap_CoordTopocentric_elevation_get_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern void _wrap_CoordTopocentric_Xrange_set_cppsgp4_5c90d950726e409b(uintptr_t arg1, double arg2);
extern double _wrap_CoordTopocentric_Xrange_get_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern void _wrap_CoordTopocentric_range_rate_set_cppsgp4_5c90d950726e409b(uintptr_t arg1, double arg2);
extern double _wrap_CoordTopocentric_range_rate_get_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern swig_type_10 _wrap_CoordTopocentric_ToString_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern void _wrap_delete_CoordTopocentric_cppsgp4_5c90d950726e409b(uintptr_t arg1);
extern uintptr_t _wrap_new_SolarPosition_cppsgp4_5c90d950726e409b(void);
extern uintptr_t _wrap_SolarPosition_FindPosition_cppsgp4_5c90d950726e409b(uintptr_t arg1, uintptr_t arg2);
extern void _wrap_delete_SolarPosition_cppsgp4_5c90d950726e409b(uintptr_t arg1);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


type _ sync.Mutex


type swig_gostring struct { p uintptr; n int }
func swigCopyString(s string) string {
  p := *(*swig_gostring)(unsafe.Pointer(&s))
  r := string((*[0x7fffffff]byte)(unsafe.Pointer(p.p))[:p.n])
  Swig_free(p.p)
  return r
}

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_cppsgp4_5c90d950726e409b(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrSGP4 uintptr

func (p SwigcptrSGP4) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrSGP4) SwigIsSGP4() {
}

func NewSGP4(arg1 Tle) (_swig_ret SGP4) {
	var swig_r SGP4
	_swig_i_0 := arg1.Swigcptr()
	swig_r = (SGP4)(SwigcptrSGP4(C._wrap_new_SGP4_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrSGP4) FindPosition(arg2 DateTime) (_swig_ret Eci) {
	var swig_r Eci
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (Eci)(SwigcptrEci(C._wrap_SGP4_FindPosition_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

func DeleteSGP4(arg1 SGP4) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_SGP4_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0))
}

type SGP4 interface {
	Swigcptr() uintptr
	SwigIsSGP4()
	FindPosition(arg2 DateTime) (_swig_ret Eci)
}

type SwigcptrTle uintptr

func (p SwigcptrTle) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrTle) SwigIsTle() {
}

func NewTle(arg1 string, arg2 string, arg3 string) (_swig_ret Tle) {
	var swig_r Tle
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (Tle)(SwigcptrTle(C._wrap_new_Tle_cppsgp4_5c90d950726e409b(*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_0)), *(*C.swig_type_2)(unsafe.Pointer(&_swig_i_1)), *(*C.swig_type_3)(unsafe.Pointer(&_swig_i_2)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg3
	}
	return swig_r
}

func (arg1 SwigcptrTle) Name() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_Tle_Name_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrTle) Line1() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_Tle_Line1_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrTle) Line2() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_Tle_Line2_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrTle) NoradNumber() (_swig_ret uint) {
	var swig_r uint
	_swig_i_0 := arg1
	swig_r = (uint)(C._wrap_Tle_NoradNumber_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrTle) IntDesignator() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_Tle_IntDesignator_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrTle) Epoch() (_swig_ret DateTime) {
	var swig_r DateTime
	_swig_i_0 := arg1
	swig_r = (DateTime)(SwigcptrDateTime(C._wrap_Tle_Epoch_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrTle) MeanMotionDt2() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_Tle_MeanMotionDt2_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrTle) MeanMotionDdt6() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_Tle_MeanMotionDdt6_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrTle) BStar() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_Tle_BStar_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrTle) Inclination(arg2 bool) (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (float64)(C._wrap_Tle_Inclination_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0), C._Bool(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrTle) RightAscendingNode(arg2 bool) (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (float64)(C._wrap_Tle_RightAscendingNode_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0), C._Bool(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrTle) Eccentricity() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_Tle_Eccentricity_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrTle) ArgumentPerigee(arg2 bool) (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (float64)(C._wrap_Tle_ArgumentPerigee_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0), C._Bool(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrTle) MeanAnomaly(arg2 bool) (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (float64)(C._wrap_Tle_MeanAnomaly_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0), C._Bool(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrTle) MeanMotion() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_Tle_MeanMotion_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrTle) OrbitNumber() (_swig_ret uint) {
	var swig_r uint
	_swig_i_0 := arg1
	swig_r = (uint)(C._wrap_Tle_OrbitNumber_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrTle) ToString() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_Tle_ToString_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func DeleteTle(arg1 Tle) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_Tle_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0))
}

type Tle interface {
	Swigcptr() uintptr
	SwigIsTle()
	Name() (_swig_ret string)
	Line1() (_swig_ret string)
	Line2() (_swig_ret string)
	NoradNumber() (_swig_ret uint)
	IntDesignator() (_swig_ret string)
	Epoch() (_swig_ret DateTime)
	MeanMotionDt2() (_swig_ret float64)
	MeanMotionDdt6() (_swig_ret float64)
	BStar() (_swig_ret float64)
	Inclination(arg2 bool) (_swig_ret float64)
	RightAscendingNode(arg2 bool) (_swig_ret float64)
	Eccentricity() (_swig_ret float64)
	ArgumentPerigee(arg2 bool) (_swig_ret float64)
	MeanAnomaly(arg2 bool) (_swig_ret float64)
	MeanMotion() (_swig_ret float64)
	OrbitNumber() (_swig_ret uint)
	ToString() (_swig_ret string)
}

type SwigcptrObserver uintptr

func (p SwigcptrObserver) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrObserver) SwigIsObserver() {
}

func NewObserver(arg1 CoordGeodetic) (_swig_ret Observer) {
	var swig_r Observer
	_swig_i_0 := arg1.Swigcptr()
	swig_r = (Observer)(SwigcptrObserver(C._wrap_new_Observer_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrObserver) GetLocation() (_swig_ret CoordGeodetic) {
	var swig_r CoordGeodetic
	_swig_i_0 := arg1
	swig_r = (CoordGeodetic)(SwigcptrCoordGeodetic(C._wrap_Observer_GetLocation_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrObserver) GetLookAngle(arg2 Eci) (_swig_ret CoordTopocentric) {
	var swig_r CoordTopocentric
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (CoordTopocentric)(SwigcptrCoordTopocentric(C._wrap_Observer_GetLookAngle_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

func DeleteObserver(arg1 Observer) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_Observer_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0))
}

type Observer interface {
	Swigcptr() uintptr
	SwigIsObserver()
	GetLocation() (_swig_ret CoordGeodetic)
	GetLookAngle(arg2 Eci) (_swig_ret CoordTopocentric)
}

type SwigcptrEci uintptr

func (p SwigcptrEci) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrEci) SwigIsEci() {
}

func NewEci(arg1 DateTime, arg2 CoordGeodetic) (_swig_ret Eci) {
	var swig_r Eci
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (Eci)(SwigcptrEci(C._wrap_new_Eci_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrEci) GetDateTime() (_swig_ret DateTime) {
	var swig_r DateTime
	_swig_i_0 := arg1
	swig_r = (DateTime)(SwigcptrDateTime(C._wrap_Eci_GetDateTime_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrEci) ToGeodetic() (_swig_ret CoordGeodetic) {
	var swig_r CoordGeodetic
	_swig_i_0 := arg1
	swig_r = (CoordGeodetic)(SwigcptrCoordGeodetic(C._wrap_Eci_ToGeodetic_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func DeleteEci(arg1 Eci) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_Eci_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0))
}

type Eci interface {
	Swigcptr() uintptr
	SwigIsEci()
	GetDateTime() (_swig_ret DateTime)
	ToGeodetic() (_swig_ret CoordGeodetic)
}

type SwigcptrDateTime uintptr

func (p SwigcptrDateTime) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrDateTime) SwigIsDateTime() {
}

func NewDateTime(arg1 int, arg2 int, arg3 int, arg4 int, arg5 int, arg6 int) (_swig_ret DateTime) {
	var swig_r DateTime
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	_swig_i_4 := arg5
	_swig_i_5 := arg6
	swig_r = (DateTime)(SwigcptrDateTime(C._wrap_new_DateTime_cppsgp4_5c90d950726e409b(C.swig_intgo(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2), C.swig_intgo(_swig_i_3), C.swig_intgo(_swig_i_4), C.swig_intgo(_swig_i_5))))
	return swig_r
}

func DateTimeNow__SWIG_0(arg1 bool) (_swig_ret DateTime) {
	var swig_r DateTime
	_swig_i_0 := arg1
	swig_r = (DateTime)(SwigcptrDateTime(C._wrap_DateTime_Now__SWIG_0_cppsgp4_5c90d950726e409b(C._Bool(_swig_i_0))))
	return swig_r
}

func DateTimeNow__SWIG_1() (_swig_ret DateTime) {
	var swig_r DateTime
	swig_r = (DateTime)(SwigcptrDateTime(C._wrap_DateTime_Now__SWIG_1_cppsgp4_5c90d950726e409b()))
	return swig_r
}

func DateTimeNow(a ...interface{}) DateTime {
	argc := len(a)
	if argc == 0 {
		return DateTimeNow__SWIG_1()
	}
	if argc == 1 {
		return DateTimeNow__SWIG_0(a[0].(bool))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrDateTime) ToJulian() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_DateTime_ToJulian_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrDateTime) Year() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_DateTime_Year_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrDateTime) Month() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_DateTime_Month_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrDateTime) Day() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_DateTime_Day_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrDateTime) Hour() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_DateTime_Hour_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrDateTime) Minute() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_DateTime_Minute_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrDateTime) Second() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_DateTime_Second_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func DeleteDateTime(arg1 DateTime) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_DateTime_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0))
}

type DateTime interface {
	Swigcptr() uintptr
	SwigIsDateTime()
	ToJulian() (_swig_ret float64)
	Year() (_swig_ret int)
	Month() (_swig_ret int)
	Day() (_swig_ret int)
	Hour() (_swig_ret int)
	Minute() (_swig_ret int)
	Second() (_swig_ret int)
}

type SwigcptrCoordGeodetic uintptr

func (p SwigcptrCoordGeodetic) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrCoordGeodetic) SwigIsCoordGeodetic() {
}

func NewCoordGeodetic__SWIG_0(arg1 float64, arg2 float64, arg3 float64, arg4 bool) (_swig_ret CoordGeodetic) {
	var swig_r CoordGeodetic
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	swig_r = (CoordGeodetic)(SwigcptrCoordGeodetic(C._wrap_new_CoordGeodetic__SWIG_0_cppsgp4_5c90d950726e409b(C.double(_swig_i_0), C.double(_swig_i_1), C.double(_swig_i_2), C._Bool(_swig_i_3))))
	return swig_r
}

func NewCoordGeodetic__SWIG_1(arg1 float64, arg2 float64, arg3 float64) (_swig_ret CoordGeodetic) {
	var swig_r CoordGeodetic
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (CoordGeodetic)(SwigcptrCoordGeodetic(C._wrap_new_CoordGeodetic__SWIG_1_cppsgp4_5c90d950726e409b(C.double(_swig_i_0), C.double(_swig_i_1), C.double(_swig_i_2))))
	return swig_r
}

func NewCoordGeodetic(a ...interface{}) CoordGeodetic {
	argc := len(a)
	if argc == 3 {
		return NewCoordGeodetic__SWIG_1(a[0].(float64), a[1].(float64), a[2].(float64))
	}
	if argc == 4 {
		return NewCoordGeodetic__SWIG_0(a[0].(float64), a[1].(float64), a[2].(float64), a[3].(bool))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrCoordGeodetic) SetLatitude(arg2 float64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_CoordGeodetic_latitude_set_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0), C.double(_swig_i_1))
}

func (arg1 SwigcptrCoordGeodetic) GetLatitude() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_CoordGeodetic_latitude_get_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrCoordGeodetic) SetLongitude(arg2 float64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_CoordGeodetic_longitude_set_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0), C.double(_swig_i_1))
}

func (arg1 SwigcptrCoordGeodetic) GetLongitude() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_CoordGeodetic_longitude_get_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrCoordGeodetic) SetAltitude(arg2 float64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_CoordGeodetic_altitude_set_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0), C.double(_swig_i_1))
}

func (arg1 SwigcptrCoordGeodetic) GetAltitude() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_CoordGeodetic_altitude_get_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrCoordGeodetic) ToString() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_CoordGeodetic_ToString_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func DeleteCoordGeodetic(arg1 CoordGeodetic) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_CoordGeodetic_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0))
}

type CoordGeodetic interface {
	Swigcptr() uintptr
	SwigIsCoordGeodetic()
	SetLatitude(arg2 float64)
	GetLatitude() (_swig_ret float64)
	SetLongitude(arg2 float64)
	GetLongitude() (_swig_ret float64)
	SetAltitude(arg2 float64)
	GetAltitude() (_swig_ret float64)
	ToString() (_swig_ret string)
}

type SwigcptrCoordTopocentric uintptr

func (p SwigcptrCoordTopocentric) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrCoordTopocentric) SwigIsCoordTopocentric() {
}

func NewCoordTopocentric(arg1 float64, arg2 float64, arg3 float64, arg4 float64) (_swig_ret CoordTopocentric) {
	var swig_r CoordTopocentric
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	swig_r = (CoordTopocentric)(SwigcptrCoordTopocentric(C._wrap_new_CoordTopocentric_cppsgp4_5c90d950726e409b(C.double(_swig_i_0), C.double(_swig_i_1), C.double(_swig_i_2), C.double(_swig_i_3))))
	return swig_r
}

func (arg1 SwigcptrCoordTopocentric) SetAzimuth(arg2 float64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_CoordTopocentric_azimuth_set_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0), C.double(_swig_i_1))
}

func (arg1 SwigcptrCoordTopocentric) GetAzimuth() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_CoordTopocentric_azimuth_get_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrCoordTopocentric) SetElevation(arg2 float64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_CoordTopocentric_elevation_set_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0), C.double(_swig_i_1))
}

func (arg1 SwigcptrCoordTopocentric) GetElevation() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_CoordTopocentric_elevation_get_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrCoordTopocentric) SetXrange(arg2 float64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_CoordTopocentric_Xrange_set_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0), C.double(_swig_i_1))
}

func (arg1 SwigcptrCoordTopocentric) GetXrange() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_CoordTopocentric_Xrange_get_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrCoordTopocentric) SetRange_rate(arg2 float64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_CoordTopocentric_range_rate_set_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0), C.double(_swig_i_1))
}

func (arg1 SwigcptrCoordTopocentric) GetRange_rate() (_swig_ret float64) {
	var swig_r float64
	_swig_i_0 := arg1
	swig_r = (float64)(C._wrap_CoordTopocentric_range_rate_get_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrCoordTopocentric) ToString() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_CoordTopocentric_ToString_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func DeleteCoordTopocentric(arg1 CoordTopocentric) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_CoordTopocentric_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0))
}

type CoordTopocentric interface {
	Swigcptr() uintptr
	SwigIsCoordTopocentric()
	SetAzimuth(arg2 float64)
	GetAzimuth() (_swig_ret float64)
	SetElevation(arg2 float64)
	GetElevation() (_swig_ret float64)
	SetXrange(arg2 float64)
	GetXrange() (_swig_ret float64)
	SetRange_rate(arg2 float64)
	GetRange_rate() (_swig_ret float64)
	ToString() (_swig_ret string)
}

type SwigcptrSolarPosition uintptr

func (p SwigcptrSolarPosition) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrSolarPosition) SwigIsSolarPosition() {
}

func NewSolarPosition() (_swig_ret SolarPosition) {
	var swig_r SolarPosition
	swig_r = (SolarPosition)(SwigcptrSolarPosition(C._wrap_new_SolarPosition_cppsgp4_5c90d950726e409b()))
	return swig_r
}

func (arg1 SwigcptrSolarPosition) FindPosition(arg2 DateTime) (_swig_ret Eci) {
	var swig_r Eci
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (Eci)(SwigcptrEci(C._wrap_SolarPosition_FindPosition_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

func DeleteSolarPosition(arg1 SolarPosition) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_SolarPosition_cppsgp4_5c90d950726e409b(C.uintptr_t(_swig_i_0))
}

type SolarPosition interface {
	Swigcptr() uintptr
	SwigIsSolarPosition()
	FindPosition(arg2 DateTime) (_swig_ret Eci)
}


