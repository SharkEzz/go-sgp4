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

#define SWIGMODULE cppsgp4

#ifdef __cplusplus
/* SwigValueWrapper is described in swig.swg */
template<typename T> class SwigValueWrapper {
  struct SwigMovePointer {
    T *ptr;
    SwigMovePointer(T *p) : ptr(p) { }
    ~SwigMovePointer() { delete ptr; }
    SwigMovePointer& operator=(SwigMovePointer& rhs) { T* oldptr = ptr; ptr = 0; delete oldptr; ptr = rhs.ptr; rhs.ptr = 0; return *this; }
  } pointer;
  SwigValueWrapper& operator=(const SwigValueWrapper<T>& rhs);
  SwigValueWrapper(const SwigValueWrapper<T>& rhs);
public:
  SwigValueWrapper() : pointer(0) { }
  SwigValueWrapper& operator=(const T& t) { SwigMovePointer tmp(new T(t)); pointer = tmp; return *this; }
  operator T&() const { return *pointer.ptr; }
  T *operator&() { return pointer.ptr; }
};

template <typename T> T SwigValueInit() {
  return T();
}
#endif

/* -----------------------------------------------------------------------------
 *  This section contains generic SWIG labels for method/variable
 *  declarations/attributes, and other compiler dependent labels.
 * ----------------------------------------------------------------------------- */

/* template workaround for compilers that cannot correctly implement the C++ standard */
#ifndef SWIGTEMPLATEDISAMBIGUATOR
# if defined(__SUNPRO_CC) && (__SUNPRO_CC <= 0x560)
#  define SWIGTEMPLATEDISAMBIGUATOR template
# elif defined(__HP_aCC)
/* Needed even with `aCC -AA' when `aCC -V' reports HP ANSI C++ B3910B A.03.55 */
/* If we find a maximum version that requires this, the test would be __HP_aCC <= 35500 for A.03.55 */
#  define SWIGTEMPLATEDISAMBIGUATOR template
# else
#  define SWIGTEMPLATEDISAMBIGUATOR
# endif
#endif

/* inline attribute */
#ifndef SWIGINLINE
# if defined(__cplusplus) || (defined(__GNUC__) && !defined(__STRICT_ANSI__))
#   define SWIGINLINE inline
# else
#   define SWIGINLINE
# endif
#endif

/* attribute recognised by some compilers to avoid 'unused' warnings */
#ifndef SWIGUNUSED
# if defined(__GNUC__)
#   if !(defined(__cplusplus)) || (__GNUC__ > 3 || (__GNUC__ == 3 && __GNUC_MINOR__ >= 4))
#     define SWIGUNUSED __attribute__ ((__unused__))
#   else
#     define SWIGUNUSED
#   endif
# elif defined(__ICC)
#   define SWIGUNUSED __attribute__ ((__unused__))
# else
#   define SWIGUNUSED
# endif
#endif

#ifndef SWIG_MSC_UNSUPPRESS_4505
# if defined(_MSC_VER)
#   pragma warning(disable : 4505) /* unreferenced local function has been removed */
# endif
#endif

#ifndef SWIGUNUSEDPARM
# ifdef __cplusplus
#   define SWIGUNUSEDPARM(p)
# else
#   define SWIGUNUSEDPARM(p) p SWIGUNUSED
# endif
#endif

/* internal SWIG method */
#ifndef SWIGINTERN
# define SWIGINTERN static SWIGUNUSED
#endif

/* internal inline SWIG method */
#ifndef SWIGINTERNINLINE
# define SWIGINTERNINLINE SWIGINTERN SWIGINLINE
#endif

/* exporting methods */
#if defined(__GNUC__)
#  if (__GNUC__ >= 4) || (__GNUC__ == 3 && __GNUC_MINOR__ >= 4)
#    ifndef GCC_HASCLASSVISIBILITY
#      define GCC_HASCLASSVISIBILITY
#    endif
#  endif
#endif

#ifndef SWIGEXPORT
# if defined(_WIN32) || defined(__WIN32__) || defined(__CYGWIN__)
#   if defined(STATIC_LINKED)
#     define SWIGEXPORT
#   else
#     define SWIGEXPORT __declspec(dllexport)
#   endif
# else
#   if defined(__GNUC__) && defined(GCC_HASCLASSVISIBILITY)
#     define SWIGEXPORT __attribute__ ((visibility("default")))
#   else
#     define SWIGEXPORT
#   endif
# endif
#endif

/* calling conventions for Windows */
#ifndef SWIGSTDCALL
# if defined(_WIN32) || defined(__WIN32__) || defined(__CYGWIN__)
#   define SWIGSTDCALL __stdcall
# else
#   define SWIGSTDCALL
# endif
#endif

/* Deal with Microsoft's attempt at deprecating C standard runtime functions */
#if !defined(SWIG_NO_CRT_SECURE_NO_DEPRECATE) && defined(_MSC_VER) && !defined(_CRT_SECURE_NO_DEPRECATE)
# define _CRT_SECURE_NO_DEPRECATE
#endif

/* Deal with Microsoft's attempt at deprecating methods in the standard C++ library */
#if !defined(SWIG_NO_SCL_SECURE_NO_DEPRECATE) && defined(_MSC_VER) && !defined(_SCL_SECURE_NO_DEPRECATE)
# define _SCL_SECURE_NO_DEPRECATE
#endif

/* Deal with Apple's deprecated 'AssertMacros.h' from Carbon-framework */
#if defined(__APPLE__) && !defined(__ASSERT_MACROS_DEFINE_VERSIONS_WITHOUT_UNDERSCORES)
# define __ASSERT_MACROS_DEFINE_VERSIONS_WITHOUT_UNDERSCORES 0
#endif

/* Intel's compiler complains if a variable which was never initialised is
 * cast to void, which is a common idiom which we use to indicate that we
 * are aware a variable isn't used.  So we just silence that warning.
 * See: https://github.com/swig/swig/issues/192 for more discussion.
 */
#ifdef __INTEL_COMPILER
# pragma warning disable 592
#endif


#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>



typedef long long intgo;
typedef unsigned long long uintgo;


# if !defined(__clang__) && (defined(__i386__) || defined(__x86_64__))
#   define SWIGSTRUCTPACKED __attribute__((__packed__, __gcc_struct__))
# else
#   define SWIGSTRUCTPACKED __attribute__((__packed__))
# endif



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;




#define swiggo_size_assert_eq(x, y, name) typedef char name[(x-y)*(x-y)*-2+1];
#define swiggo_size_assert(t, n) swiggo_size_assert_eq(sizeof(t), n, swiggo_sizeof_##t##_is_not_##n)

swiggo_size_assert(char, 1)
swiggo_size_assert(short, 2)
swiggo_size_assert(int, 4)
typedef long long swiggo_long_long;
swiggo_size_assert(swiggo_long_long, 8)
swiggo_size_assert(float, 4)
swiggo_size_assert(double, 8)

#ifdef __cplusplus
extern "C" {
#endif
extern void crosscall2(void (*fn)(void *, int), void *, int);
extern char* _cgo_topofstack(void) __attribute__ ((weak));
extern void _cgo_allocate(void *, int);
extern void _cgo_panic(void *, int);
#ifdef __cplusplus
}
#endif

static char *_swig_topofstack() {
  if (_cgo_topofstack) {
    return _cgo_topofstack();
  } else {
    return 0;
  }
}

static void _swig_gopanic(const char *p) {
  struct {
    const char *p;
  } SWIGSTRUCTPACKED a;
  a.p = p;
  crosscall2(_cgo_panic, &a, (int) sizeof a);
}




#define SWIG_contract_assert(expr, msg) \
  if (!(expr)) { _swig_gopanic(msg); } else


#define SWIG_exception(code, msg) _swig_gopanic(msg)


static _gostring_ Swig_AllocateString(const char *p, size_t l) {
  _gostring_ ret;
  ret.p = (char*)malloc(l);
  memcpy(ret.p, p, l);
  ret.n = l;
  return ret;
}


static void Swig_free(void* p) {
  free(p);
}

static void* Swig_malloc(int c) {
  return malloc(c);
}


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


#include <string>


#include <vector>
#include <stdexcept>

#ifdef __cplusplus
extern "C" {
#endif

void _wrap_Swig_free_cppsgp4_ae7bdd03b5e00c24(void *_swig_go_0) {
  void *arg1 = (void *) 0 ;
  
  arg1 = *(void **)&_swig_go_0; 
  
  Swig_free(arg1);
  
}


void *_wrap_Swig_malloc_cppsgp4_ae7bdd03b5e00c24(intgo _swig_go_0) {
  int arg1 ;
  void *result = 0 ;
  void *_swig_go_result;
  
  arg1 = (int)_swig_go_0; 
  
  result = (void *)Swig_malloc(arg1);
  *(void **)&_swig_go_result = (void *)result; 
  return _swig_go_result;
}


SGP4 *_wrap_new_SGP4_cppsgp4_ae7bdd03b5e00c24(Tle *_swig_go_0) {
  Tle *arg1 = 0 ;
  SGP4 *result = 0 ;
  SGP4 *_swig_go_result;
  
  arg1 = *(Tle **)&_swig_go_0; 
  
  {
    try {
      result = (SGP4 *)new SGP4((Tle const &)*arg1);;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  *(SGP4 **)&_swig_go_result = (SGP4 *)result; 
  return _swig_go_result;
}


Eci *_wrap_SGP4_FindPosition_cppsgp4_ae7bdd03b5e00c24(SGP4 *_swig_go_0, DateTime *_swig_go_1) {
  SGP4 *arg1 = (SGP4 *) 0 ;
  DateTime *arg2 = 0 ;
  SwigValueWrapper< Eci > result;
  Eci *_swig_go_result;
  
  arg1 = *(SGP4 **)&_swig_go_0; 
  arg2 = *(DateTime **)&_swig_go_1; 
  
  {
    try {
      result = ((SGP4 const *)arg1)->FindPosition((DateTime const &)*arg2);;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  *(Eci **)&_swig_go_result = new Eci(result); 
  return _swig_go_result;
}


void _wrap_delete_SGP4_cppsgp4_ae7bdd03b5e00c24(SGP4 *_swig_go_0) {
  SGP4 *arg1 = (SGP4 *) 0 ;
  
  arg1 = *(SGP4 **)&_swig_go_0; 
  
  {
    try {
      delete arg1;;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  
}


Tle *_wrap_new_Tle_cppsgp4_ae7bdd03b5e00c24(_gostring_ _swig_go_0, _gostring_ _swig_go_1, _gostring_ _swig_go_2) {
  std::string *arg1 = 0 ;
  std::string *arg2 = 0 ;
  std::string *arg3 = 0 ;
  Tle *result = 0 ;
  Tle *_swig_go_result;
  
  
  std::string arg1_str(_swig_go_0.p, _swig_go_0.n);
  arg1 = &arg1_str;
  
  
  std::string arg2_str(_swig_go_1.p, _swig_go_1.n);
  arg2 = &arg2_str;
  
  
  std::string arg3_str(_swig_go_2.p, _swig_go_2.n);
  arg3 = &arg3_str;
  
  
  {
    try {
      result = (Tle *)new Tle((std::string const &)*arg1,(std::string const &)*arg2,(std::string const &)*arg3);;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  *(Tle **)&_swig_go_result = (Tle *)result; 
  return _swig_go_result;
}


_gostring_ _wrap_Tle_Name_cppsgp4_ae7bdd03b5e00c24(Tle *_swig_go_0) {
  Tle *arg1 = (Tle *) 0 ;
  std::string result;
  _gostring_ _swig_go_result;
  
  arg1 = *(Tle **)&_swig_go_0; 
  
  {
    try {
      result = ((Tle const *)arg1)->Name();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  _swig_go_result = Swig_AllocateString((&result)->data(), (&result)->length()); 
  return _swig_go_result;
}


_gostring_ _wrap_Tle_Line1_cppsgp4_ae7bdd03b5e00c24(Tle *_swig_go_0) {
  Tle *arg1 = (Tle *) 0 ;
  std::string result;
  _gostring_ _swig_go_result;
  
  arg1 = *(Tle **)&_swig_go_0; 
  
  {
    try {
      result = ((Tle const *)arg1)->Line1();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  _swig_go_result = Swig_AllocateString((&result)->data(), (&result)->length()); 
  return _swig_go_result;
}


_gostring_ _wrap_Tle_Line2_cppsgp4_ae7bdd03b5e00c24(Tle *_swig_go_0) {
  Tle *arg1 = (Tle *) 0 ;
  std::string result;
  _gostring_ _swig_go_result;
  
  arg1 = *(Tle **)&_swig_go_0; 
  
  {
    try {
      result = ((Tle const *)arg1)->Line2();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  _swig_go_result = Swig_AllocateString((&result)->data(), (&result)->length()); 
  return _swig_go_result;
}


intgo _wrap_Tle_NoradNumber_cppsgp4_ae7bdd03b5e00c24(Tle *_swig_go_0) {
  Tle *arg1 = (Tle *) 0 ;
  unsigned int result;
  intgo _swig_go_result;
  
  arg1 = *(Tle **)&_swig_go_0; 
  
  {
    try {
      result = (unsigned int)((Tle const *)arg1)->NoradNumber();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  _swig_go_result = result; 
  return _swig_go_result;
}


_gostring_ _wrap_Tle_IntDesignator_cppsgp4_ae7bdd03b5e00c24(Tle *_swig_go_0) {
  Tle *arg1 = (Tle *) 0 ;
  std::string result;
  _gostring_ _swig_go_result;
  
  arg1 = *(Tle **)&_swig_go_0; 
  
  {
    try {
      result = ((Tle const *)arg1)->IntDesignator();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  _swig_go_result = Swig_AllocateString((&result)->data(), (&result)->length()); 
  return _swig_go_result;
}


DateTime *_wrap_Tle_Epoch_cppsgp4_ae7bdd03b5e00c24(Tle *_swig_go_0) {
  Tle *arg1 = (Tle *) 0 ;
  SwigValueWrapper< DateTime > result;
  DateTime *_swig_go_result;
  
  arg1 = *(Tle **)&_swig_go_0; 
  
  {
    try {
      result = ((Tle const *)arg1)->Epoch();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  *(DateTime **)&_swig_go_result = new DateTime(result); 
  return _swig_go_result;
}


double _wrap_Tle_MeanMotionDt2_cppsgp4_ae7bdd03b5e00c24(Tle *_swig_go_0) {
  Tle *arg1 = (Tle *) 0 ;
  double result;
  double _swig_go_result;
  
  arg1 = *(Tle **)&_swig_go_0; 
  
  {
    try {
      result = (double)((Tle const *)arg1)->MeanMotionDt2();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  _swig_go_result = result; 
  return _swig_go_result;
}


double _wrap_Tle_MeanMotionDdt6_cppsgp4_ae7bdd03b5e00c24(Tle *_swig_go_0) {
  Tle *arg1 = (Tle *) 0 ;
  double result;
  double _swig_go_result;
  
  arg1 = *(Tle **)&_swig_go_0; 
  
  {
    try {
      result = (double)((Tle const *)arg1)->MeanMotionDdt6();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  _swig_go_result = result; 
  return _swig_go_result;
}


double _wrap_Tle_BStar_cppsgp4_ae7bdd03b5e00c24(Tle *_swig_go_0) {
  Tle *arg1 = (Tle *) 0 ;
  double result;
  double _swig_go_result;
  
  arg1 = *(Tle **)&_swig_go_0; 
  
  {
    try {
      result = (double)((Tle const *)arg1)->BStar();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  _swig_go_result = result; 
  return _swig_go_result;
}


double _wrap_Tle_Inclination_cppsgp4_ae7bdd03b5e00c24(Tle *_swig_go_0, bool _swig_go_1) {
  Tle *arg1 = (Tle *) 0 ;
  bool arg2 ;
  double result;
  double _swig_go_result;
  
  arg1 = *(Tle **)&_swig_go_0; 
  arg2 = (bool)_swig_go_1; 
  
  {
    try {
      result = (double)((Tle const *)arg1)->Inclination(arg2);;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  _swig_go_result = result; 
  return _swig_go_result;
}


double _wrap_Tle_RightAscendingNode_cppsgp4_ae7bdd03b5e00c24(Tle *_swig_go_0, bool _swig_go_1) {
  Tle *arg1 = (Tle *) 0 ;
  bool arg2 ;
  double result;
  double _swig_go_result;
  
  arg1 = *(Tle **)&_swig_go_0; 
  arg2 = (bool)_swig_go_1; 
  
  {
    try {
      result = (double)((Tle const *)arg1)->RightAscendingNode(arg2);;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  _swig_go_result = result; 
  return _swig_go_result;
}


double _wrap_Tle_Eccentricity_cppsgp4_ae7bdd03b5e00c24(Tle *_swig_go_0) {
  Tle *arg1 = (Tle *) 0 ;
  double result;
  double _swig_go_result;
  
  arg1 = *(Tle **)&_swig_go_0; 
  
  {
    try {
      result = (double)((Tle const *)arg1)->Eccentricity();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  _swig_go_result = result; 
  return _swig_go_result;
}


double _wrap_Tle_ArgumentPerigee_cppsgp4_ae7bdd03b5e00c24(Tle *_swig_go_0, bool _swig_go_1) {
  Tle *arg1 = (Tle *) 0 ;
  bool arg2 ;
  double result;
  double _swig_go_result;
  
  arg1 = *(Tle **)&_swig_go_0; 
  arg2 = (bool)_swig_go_1; 
  
  {
    try {
      result = (double)((Tle const *)arg1)->ArgumentPerigee(arg2);;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  _swig_go_result = result; 
  return _swig_go_result;
}


double _wrap_Tle_MeanAnomaly_cppsgp4_ae7bdd03b5e00c24(Tle *_swig_go_0, bool _swig_go_1) {
  Tle *arg1 = (Tle *) 0 ;
  bool arg2 ;
  double result;
  double _swig_go_result;
  
  arg1 = *(Tle **)&_swig_go_0; 
  arg2 = (bool)_swig_go_1; 
  
  {
    try {
      result = (double)((Tle const *)arg1)->MeanAnomaly(arg2);;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  _swig_go_result = result; 
  return _swig_go_result;
}


double _wrap_Tle_MeanMotion_cppsgp4_ae7bdd03b5e00c24(Tle *_swig_go_0) {
  Tle *arg1 = (Tle *) 0 ;
  double result;
  double _swig_go_result;
  
  arg1 = *(Tle **)&_swig_go_0; 
  
  {
    try {
      result = (double)((Tle const *)arg1)->MeanMotion();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  _swig_go_result = result; 
  return _swig_go_result;
}


intgo _wrap_Tle_OrbitNumber_cppsgp4_ae7bdd03b5e00c24(Tle *_swig_go_0) {
  Tle *arg1 = (Tle *) 0 ;
  unsigned int result;
  intgo _swig_go_result;
  
  arg1 = *(Tle **)&_swig_go_0; 
  
  {
    try {
      result = (unsigned int)((Tle const *)arg1)->OrbitNumber();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  _swig_go_result = result; 
  return _swig_go_result;
}


_gostring_ _wrap_Tle_ToString_cppsgp4_ae7bdd03b5e00c24(Tle *_swig_go_0) {
  Tle *arg1 = (Tle *) 0 ;
  std::string result;
  _gostring_ _swig_go_result;
  
  arg1 = *(Tle **)&_swig_go_0; 
  
  {
    try {
      result = ((Tle const *)arg1)->ToString();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  _swig_go_result = Swig_AllocateString((&result)->data(), (&result)->length()); 
  return _swig_go_result;
}


void _wrap_delete_Tle_cppsgp4_ae7bdd03b5e00c24(Tle *_swig_go_0) {
  Tle *arg1 = (Tle *) 0 ;
  
  arg1 = *(Tle **)&_swig_go_0; 
  
  {
    try {
      delete arg1;;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  
}


Observer *_wrap_new_Observer_cppsgp4_ae7bdd03b5e00c24(CoordGeodetic *_swig_go_0) {
  CoordGeodetic *arg1 = 0 ;
  Observer *result = 0 ;
  Observer *_swig_go_result;
  
  arg1 = *(CoordGeodetic **)&_swig_go_0; 
  
  {
    try {
      result = (Observer *)new Observer((CoordGeodetic const &)*arg1);;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  *(Observer **)&_swig_go_result = (Observer *)result; 
  return _swig_go_result;
}


CoordGeodetic *_wrap_Observer_GetLocation_cppsgp4_ae7bdd03b5e00c24(Observer *_swig_go_0) {
  Observer *arg1 = (Observer *) 0 ;
  CoordGeodetic result;
  CoordGeodetic *_swig_go_result;
  
  arg1 = *(Observer **)&_swig_go_0; 
  
  {
    try {
      result = ((Observer const *)arg1)->GetLocation();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  *(CoordGeodetic **)&_swig_go_result = new CoordGeodetic(result); 
  return _swig_go_result;
}


CoordTopocentric *_wrap_Observer_GetLookAngle_cppsgp4_ae7bdd03b5e00c24(Observer *_swig_go_0, Eci *_swig_go_1) {
  Observer *arg1 = (Observer *) 0 ;
  Eci *arg2 = 0 ;
  CoordTopocentric result;
  CoordTopocentric *_swig_go_result;
  
  arg1 = *(Observer **)&_swig_go_0; 
  arg2 = *(Eci **)&_swig_go_1; 
  
  {
    try {
      result = (arg1)->GetLookAngle((Eci const &)*arg2);;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  *(CoordTopocentric **)&_swig_go_result = new CoordTopocentric(result); 
  return _swig_go_result;
}


void _wrap_delete_Observer_cppsgp4_ae7bdd03b5e00c24(Observer *_swig_go_0) {
  Observer *arg1 = (Observer *) 0 ;
  
  arg1 = *(Observer **)&_swig_go_0; 
  
  {
    try {
      delete arg1;;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  
}


Eci *_wrap_new_Eci_cppsgp4_ae7bdd03b5e00c24(DateTime *_swig_go_0, CoordGeodetic *_swig_go_1) {
  DateTime *arg1 = 0 ;
  CoordGeodetic *arg2 = 0 ;
  Eci *result = 0 ;
  Eci *_swig_go_result;
  
  arg1 = *(DateTime **)&_swig_go_0; 
  arg2 = *(CoordGeodetic **)&_swig_go_1; 
  
  {
    try {
      result = (Eci *)new Eci((DateTime const &)*arg1,(CoordGeodetic const &)*arg2);;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  *(Eci **)&_swig_go_result = (Eci *)result; 
  return _swig_go_result;
}


Vector *_wrap_Eci_Velocity_cppsgp4_ae7bdd03b5e00c24(Eci *_swig_go_0) {
  Eci *arg1 = (Eci *) 0 ;
  Vector result;
  Vector *_swig_go_result;
  
  arg1 = *(Eci **)&_swig_go_0; 
  
  {
    try {
      result = ((Eci const *)arg1)->Velocity();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  *(Vector **)&_swig_go_result = new Vector(result); 
  return _swig_go_result;
}


DateTime *_wrap_Eci_GetDateTime_cppsgp4_ae7bdd03b5e00c24(Eci *_swig_go_0) {
  Eci *arg1 = (Eci *) 0 ;
  SwigValueWrapper< DateTime > result;
  DateTime *_swig_go_result;
  
  arg1 = *(Eci **)&_swig_go_0; 
  
  {
    try {
      result = ((Eci const *)arg1)->GetDateTime();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  *(DateTime **)&_swig_go_result = new DateTime(result); 
  return _swig_go_result;
}


CoordGeodetic *_wrap_Eci_ToGeodetic_cppsgp4_ae7bdd03b5e00c24(Eci *_swig_go_0) {
  Eci *arg1 = (Eci *) 0 ;
  CoordGeodetic result;
  CoordGeodetic *_swig_go_result;
  
  arg1 = *(Eci **)&_swig_go_0; 
  
  {
    try {
      result = ((Eci const *)arg1)->ToGeodetic();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  *(CoordGeodetic **)&_swig_go_result = new CoordGeodetic(result); 
  return _swig_go_result;
}


void _wrap_delete_Eci_cppsgp4_ae7bdd03b5e00c24(Eci *_swig_go_0) {
  Eci *arg1 = (Eci *) 0 ;
  
  arg1 = *(Eci **)&_swig_go_0; 
  
  {
    try {
      delete arg1;;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  
}


DateTime *_wrap_new_DateTime_cppsgp4_ae7bdd03b5e00c24(intgo _swig_go_0, intgo _swig_go_1, intgo _swig_go_2, intgo _swig_go_3, intgo _swig_go_4, intgo _swig_go_5) {
  int arg1 ;
  int arg2 ;
  int arg3 ;
  int arg4 ;
  int arg5 ;
  int arg6 ;
  DateTime *result = 0 ;
  DateTime *_swig_go_result;
  
  arg1 = (int)_swig_go_0; 
  arg2 = (int)_swig_go_1; 
  arg3 = (int)_swig_go_2; 
  arg4 = (int)_swig_go_3; 
  arg5 = (int)_swig_go_4; 
  arg6 = (int)_swig_go_5; 
  
  {
    try {
      result = (DateTime *)new DateTime(arg1,arg2,arg3,arg4,arg5,arg6);;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  *(DateTime **)&_swig_go_result = (DateTime *)result; 
  return _swig_go_result;
}


DateTime *_wrap_DateTime_Now__SWIG_0_cppsgp4_ae7bdd03b5e00c24(bool _swig_go_0) {
  bool arg1 ;
  SwigValueWrapper< DateTime > result;
  DateTime *_swig_go_result;
  
  arg1 = (bool)_swig_go_0; 
  
  {
    try {
      result = DateTime::Now(arg1);;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  *(DateTime **)&_swig_go_result = new DateTime(result); 
  return _swig_go_result;
}


DateTime *_wrap_DateTime_Now__SWIG_1_cppsgp4_ae7bdd03b5e00c24() {
  SwigValueWrapper< DateTime > result;
  DateTime *_swig_go_result;
  
  
  {
    try {
      result = DateTime::Now();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  *(DateTime **)&_swig_go_result = new DateTime(result); 
  return _swig_go_result;
}


double _wrap_DateTime_ToJulian_cppsgp4_ae7bdd03b5e00c24(DateTime *_swig_go_0) {
  DateTime *arg1 = (DateTime *) 0 ;
  double result;
  double _swig_go_result;
  
  arg1 = *(DateTime **)&_swig_go_0; 
  
  {
    try {
      result = (double)((DateTime const *)arg1)->ToJulian();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  _swig_go_result = result; 
  return _swig_go_result;
}


intgo _wrap_DateTime_Year_cppsgp4_ae7bdd03b5e00c24(DateTime *_swig_go_0) {
  DateTime *arg1 = (DateTime *) 0 ;
  int result;
  intgo _swig_go_result;
  
  arg1 = *(DateTime **)&_swig_go_0; 
  
  {
    try {
      result = (int)((DateTime const *)arg1)->Year();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  _swig_go_result = result; 
  return _swig_go_result;
}


intgo _wrap_DateTime_Month_cppsgp4_ae7bdd03b5e00c24(DateTime *_swig_go_0) {
  DateTime *arg1 = (DateTime *) 0 ;
  int result;
  intgo _swig_go_result;
  
  arg1 = *(DateTime **)&_swig_go_0; 
  
  {
    try {
      result = (int)((DateTime const *)arg1)->Month();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  _swig_go_result = result; 
  return _swig_go_result;
}


intgo _wrap_DateTime_Day_cppsgp4_ae7bdd03b5e00c24(DateTime *_swig_go_0) {
  DateTime *arg1 = (DateTime *) 0 ;
  int result;
  intgo _swig_go_result;
  
  arg1 = *(DateTime **)&_swig_go_0; 
  
  {
    try {
      result = (int)((DateTime const *)arg1)->Day();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  _swig_go_result = result; 
  return _swig_go_result;
}


intgo _wrap_DateTime_Hour_cppsgp4_ae7bdd03b5e00c24(DateTime *_swig_go_0) {
  DateTime *arg1 = (DateTime *) 0 ;
  int result;
  intgo _swig_go_result;
  
  arg1 = *(DateTime **)&_swig_go_0; 
  
  {
    try {
      result = (int)((DateTime const *)arg1)->Hour();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  _swig_go_result = result; 
  return _swig_go_result;
}


intgo _wrap_DateTime_Minute_cppsgp4_ae7bdd03b5e00c24(DateTime *_swig_go_0) {
  DateTime *arg1 = (DateTime *) 0 ;
  int result;
  intgo _swig_go_result;
  
  arg1 = *(DateTime **)&_swig_go_0; 
  
  {
    try {
      result = (int)((DateTime const *)arg1)->Minute();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  _swig_go_result = result; 
  return _swig_go_result;
}


intgo _wrap_DateTime_Second_cppsgp4_ae7bdd03b5e00c24(DateTime *_swig_go_0) {
  DateTime *arg1 = (DateTime *) 0 ;
  int result;
  intgo _swig_go_result;
  
  arg1 = *(DateTime **)&_swig_go_0; 
  
  {
    try {
      result = (int)((DateTime const *)arg1)->Second();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  _swig_go_result = result; 
  return _swig_go_result;
}


void _wrap_delete_DateTime_cppsgp4_ae7bdd03b5e00c24(DateTime *_swig_go_0) {
  DateTime *arg1 = (DateTime *) 0 ;
  
  arg1 = *(DateTime **)&_swig_go_0; 
  
  {
    try {
      delete arg1;;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  
}


void _wrap_CoordGeodetic_latitude_set_cppsgp4_ae7bdd03b5e00c24(CoordGeodetic *_swig_go_0, double _swig_go_1) {
  CoordGeodetic *arg1 = (CoordGeodetic *) 0 ;
  double arg2 ;
  
  arg1 = *(CoordGeodetic **)&_swig_go_0; 
  arg2 = (double)_swig_go_1; 
  
  if (arg1) (arg1)->latitude = arg2;
  
}


double _wrap_CoordGeodetic_latitude_get_cppsgp4_ae7bdd03b5e00c24(CoordGeodetic *_swig_go_0) {
  CoordGeodetic *arg1 = (CoordGeodetic *) 0 ;
  double result;
  double _swig_go_result;
  
  arg1 = *(CoordGeodetic **)&_swig_go_0; 
  
  result = (double) ((arg1)->latitude);
  _swig_go_result = result; 
  return _swig_go_result;
}


void _wrap_CoordGeodetic_longitude_set_cppsgp4_ae7bdd03b5e00c24(CoordGeodetic *_swig_go_0, double _swig_go_1) {
  CoordGeodetic *arg1 = (CoordGeodetic *) 0 ;
  double arg2 ;
  
  arg1 = *(CoordGeodetic **)&_swig_go_0; 
  arg2 = (double)_swig_go_1; 
  
  if (arg1) (arg1)->longitude = arg2;
  
}


double _wrap_CoordGeodetic_longitude_get_cppsgp4_ae7bdd03b5e00c24(CoordGeodetic *_swig_go_0) {
  CoordGeodetic *arg1 = (CoordGeodetic *) 0 ;
  double result;
  double _swig_go_result;
  
  arg1 = *(CoordGeodetic **)&_swig_go_0; 
  
  result = (double) ((arg1)->longitude);
  _swig_go_result = result; 
  return _swig_go_result;
}


void _wrap_CoordGeodetic_altitude_set_cppsgp4_ae7bdd03b5e00c24(CoordGeodetic *_swig_go_0, double _swig_go_1) {
  CoordGeodetic *arg1 = (CoordGeodetic *) 0 ;
  double arg2 ;
  
  arg1 = *(CoordGeodetic **)&_swig_go_0; 
  arg2 = (double)_swig_go_1; 
  
  if (arg1) (arg1)->altitude = arg2;
  
}


double _wrap_CoordGeodetic_altitude_get_cppsgp4_ae7bdd03b5e00c24(CoordGeodetic *_swig_go_0) {
  CoordGeodetic *arg1 = (CoordGeodetic *) 0 ;
  double result;
  double _swig_go_result;
  
  arg1 = *(CoordGeodetic **)&_swig_go_0; 
  
  result = (double) ((arg1)->altitude);
  _swig_go_result = result; 
  return _swig_go_result;
}


_gostring_ _wrap_CoordGeodetic_ToString_cppsgp4_ae7bdd03b5e00c24(CoordGeodetic *_swig_go_0) {
  CoordGeodetic *arg1 = (CoordGeodetic *) 0 ;
  std::string result;
  _gostring_ _swig_go_result;
  
  arg1 = *(CoordGeodetic **)&_swig_go_0; 
  
  {
    try {
      result = ((CoordGeodetic const *)arg1)->ToString();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  _swig_go_result = Swig_AllocateString((&result)->data(), (&result)->length()); 
  return _swig_go_result;
}


CoordGeodetic *_wrap_new_CoordGeodetic_cppsgp4_ae7bdd03b5e00c24() {
  CoordGeodetic *result = 0 ;
  CoordGeodetic *_swig_go_result;
  
  
  {
    try {
      result = (CoordGeodetic *)new CoordGeodetic();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  *(CoordGeodetic **)&_swig_go_result = (CoordGeodetic *)result; 
  return _swig_go_result;
}


void _wrap_delete_CoordGeodetic_cppsgp4_ae7bdd03b5e00c24(CoordGeodetic *_swig_go_0) {
  CoordGeodetic *arg1 = (CoordGeodetic *) 0 ;
  
  arg1 = *(CoordGeodetic **)&_swig_go_0; 
  
  {
    try {
      delete arg1;;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  
}


void _wrap_CoordTopocentric_azimuth_set_cppsgp4_ae7bdd03b5e00c24(CoordTopocentric *_swig_go_0, double _swig_go_1) {
  CoordTopocentric *arg1 = (CoordTopocentric *) 0 ;
  double arg2 ;
  
  arg1 = *(CoordTopocentric **)&_swig_go_0; 
  arg2 = (double)_swig_go_1; 
  
  if (arg1) (arg1)->azimuth = arg2;
  
}


double _wrap_CoordTopocentric_azimuth_get_cppsgp4_ae7bdd03b5e00c24(CoordTopocentric *_swig_go_0) {
  CoordTopocentric *arg1 = (CoordTopocentric *) 0 ;
  double result;
  double _swig_go_result;
  
  arg1 = *(CoordTopocentric **)&_swig_go_0; 
  
  result = (double) ((arg1)->azimuth);
  _swig_go_result = result; 
  return _swig_go_result;
}


void _wrap_CoordTopocentric_elevation_set_cppsgp4_ae7bdd03b5e00c24(CoordTopocentric *_swig_go_0, double _swig_go_1) {
  CoordTopocentric *arg1 = (CoordTopocentric *) 0 ;
  double arg2 ;
  
  arg1 = *(CoordTopocentric **)&_swig_go_0; 
  arg2 = (double)_swig_go_1; 
  
  if (arg1) (arg1)->elevation = arg2;
  
}


double _wrap_CoordTopocentric_elevation_get_cppsgp4_ae7bdd03b5e00c24(CoordTopocentric *_swig_go_0) {
  CoordTopocentric *arg1 = (CoordTopocentric *) 0 ;
  double result;
  double _swig_go_result;
  
  arg1 = *(CoordTopocentric **)&_swig_go_0; 
  
  result = (double) ((arg1)->elevation);
  _swig_go_result = result; 
  return _swig_go_result;
}


void _wrap_CoordTopocentric_Xrange_set_cppsgp4_ae7bdd03b5e00c24(CoordTopocentric *_swig_go_0, double _swig_go_1) {
  CoordTopocentric *arg1 = (CoordTopocentric *) 0 ;
  double arg2 ;
  
  arg1 = *(CoordTopocentric **)&_swig_go_0; 
  arg2 = (double)_swig_go_1; 
  
  if (arg1) (arg1)->range = arg2;
  
}


double _wrap_CoordTopocentric_Xrange_get_cppsgp4_ae7bdd03b5e00c24(CoordTopocentric *_swig_go_0) {
  CoordTopocentric *arg1 = (CoordTopocentric *) 0 ;
  double result;
  double _swig_go_result;
  
  arg1 = *(CoordTopocentric **)&_swig_go_0; 
  
  result = (double) ((arg1)->range);
  _swig_go_result = result; 
  return _swig_go_result;
}


void _wrap_CoordTopocentric_range_rate_set_cppsgp4_ae7bdd03b5e00c24(CoordTopocentric *_swig_go_0, double _swig_go_1) {
  CoordTopocentric *arg1 = (CoordTopocentric *) 0 ;
  double arg2 ;
  
  arg1 = *(CoordTopocentric **)&_swig_go_0; 
  arg2 = (double)_swig_go_1; 
  
  if (arg1) (arg1)->range_rate = arg2;
  
}


double _wrap_CoordTopocentric_range_rate_get_cppsgp4_ae7bdd03b5e00c24(CoordTopocentric *_swig_go_0) {
  CoordTopocentric *arg1 = (CoordTopocentric *) 0 ;
  double result;
  double _swig_go_result;
  
  arg1 = *(CoordTopocentric **)&_swig_go_0; 
  
  result = (double) ((arg1)->range_rate);
  _swig_go_result = result; 
  return _swig_go_result;
}


_gostring_ _wrap_CoordTopocentric_ToString_cppsgp4_ae7bdd03b5e00c24(CoordTopocentric *_swig_go_0) {
  CoordTopocentric *arg1 = (CoordTopocentric *) 0 ;
  std::string result;
  _gostring_ _swig_go_result;
  
  arg1 = *(CoordTopocentric **)&_swig_go_0; 
  
  {
    try {
      result = ((CoordTopocentric const *)arg1)->ToString();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  _swig_go_result = Swig_AllocateString((&result)->data(), (&result)->length()); 
  return _swig_go_result;
}


CoordTopocentric *_wrap_new_CoordTopocentric_cppsgp4_ae7bdd03b5e00c24() {
  CoordTopocentric *result = 0 ;
  CoordTopocentric *_swig_go_result;
  
  
  {
    try {
      result = (CoordTopocentric *)new CoordTopocentric();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  *(CoordTopocentric **)&_swig_go_result = (CoordTopocentric *)result; 
  return _swig_go_result;
}


void _wrap_delete_CoordTopocentric_cppsgp4_ae7bdd03b5e00c24(CoordTopocentric *_swig_go_0) {
  CoordTopocentric *arg1 = (CoordTopocentric *) 0 ;
  
  arg1 = *(CoordTopocentric **)&_swig_go_0; 
  
  {
    try {
      delete arg1;;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  
}


SolarPosition *_wrap_new_SolarPosition_cppsgp4_ae7bdd03b5e00c24() {
  SolarPosition *result = 0 ;
  SolarPosition *_swig_go_result;
  
  
  {
    try {
      result = (SolarPosition *)new SolarPosition();;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  *(SolarPosition **)&_swig_go_result = (SolarPosition *)result; 
  return _swig_go_result;
}


Eci *_wrap_SolarPosition_FindPosition_cppsgp4_ae7bdd03b5e00c24(SolarPosition *_swig_go_0, DateTime *_swig_go_1) {
  SolarPosition *arg1 = (SolarPosition *) 0 ;
  DateTime *arg2 = 0 ;
  SwigValueWrapper< Eci > result;
  Eci *_swig_go_result;
  
  arg1 = *(SolarPosition **)&_swig_go_0; 
  arg2 = *(DateTime **)&_swig_go_1; 
  
  {
    try {
      result = (arg1)->FindPosition((DateTime const &)*arg2);;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  *(Eci **)&_swig_go_result = new Eci(result); 
  return _swig_go_result;
}


void _wrap_delete_SolarPosition_cppsgp4_ae7bdd03b5e00c24(SolarPosition *_swig_go_0) {
  SolarPosition *arg1 = (SolarPosition *) 0 ;
  
  arg1 = *(SolarPosition **)&_swig_go_0; 
  
  {
    try {
      delete arg1;;
    } catch (std::runtime_error &e) {
      _swig_gopanic(e.what());
    }
  }
  
}


#ifdef __cplusplus
}
#endif

