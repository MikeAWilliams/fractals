cmake_minimum_required (VERSION 3.0)

project (mandelbrot_test)

add_executable(mandelbrot_test mandelbrot.cpp testmandelbrot.cpp)
set_property(TARGET mandelbrot_test PROPERTY CXX_STANDARD 17)
SET(CMAKE_CXX_FLAGS ${CMAKE_CXX_FLAGS} "-pthread")

find_package(Catch2 CONFIG REQUIRED)

target_link_libraries(mandelbrot_test PRIVATE Catch2::Catch2)