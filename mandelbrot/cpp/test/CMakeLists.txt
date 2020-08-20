cmake_minimum_required (VERSION 3.0)

project (mandelbrot_test)

add_executable(mandelbrot_test testmandelbrot.cpp)
set_property(TARGET mandelbrot_test PROPERTY CXX_STANDARD 17)
SET(CMAKE_CXX_FLAGS ${CMAKE_CXX_FLAGS} "-pthread")

include_directories(${CMAKE_SOURCE_DIR}/lib)

target_link_libraries(mandelbrot_test PRIVATE Catch2::Catch2)
target_link_libraries(mandelbrot_test PRIVATE mandelbrot_lib)