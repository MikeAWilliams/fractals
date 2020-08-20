cmake_minimum_required (VERSION 3.0)

project (mandelbrot_lib)

add_library(mandelbrot_lib mandelbrot.cpp rgb.cpp)
set_property(TARGET mandelbrot_lib PROPERTY CXX_STANDARD 17)
SET(CMAKE_CXX_FLAGS ${CMAKE_CXX_FLAGS} "-pthread")