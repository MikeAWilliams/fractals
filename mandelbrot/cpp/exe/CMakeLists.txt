cmake_minimum_required (VERSION 3.0)

project (mandelbrot)

add_executable(mandelbrot mandelbrot_main.cpp)
set_property(TARGET mandelbrot PROPERTY CXX_STANDARD 17)
SET(CMAKE_CXX_FLAGS ${CMAKE_CXX_FLAGS} "-pthread")

include_directories(${CMAKE_SOURCE_DIR}/lib)

target_link_libraries(mandelbrot PRIVATE mandelbrot_lib)
target_link_libraries(mandelbrot PRIVATE PNGwriter::PNGwriter)
target_link_libraries(mandelbrot PRIVATE tbb)