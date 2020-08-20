cmake_minimum_required(VERSION 3.2.2)
project(mandalbrotMaster)

find_package(PNGwriter CONFIG REQUIRED)
find_package(Catch2 CONFIG REQUIRED)

add_subdirectory(lib)
add_subdirectory(test)