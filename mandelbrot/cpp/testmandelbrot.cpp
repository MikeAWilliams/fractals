#include "mandelbrot.h"

#define CATCH_CONFIG_MAIN 
#include "catch2/catch.hpp"

TEST_CASE("first call")
{
    auto result01 {ComputeMandelbrot(30, 0, 1)};
    REQUIRE_FALSE(result01.has_value());

    auto result10 {ComputeMandelbrot(30, 1, 0)};
    REQUIRE(result10.has_value());
    REQUIRE(2 == result10.value());
}