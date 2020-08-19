#include "mandelbrot.h"

#define CATCH_CONFIG_MAIN 
#include "catch2/catch.hpp"

TEST_CASE("first call")
{
    auto result {ComputeMandelbrot(30, 0, 1)};
    REQUIRE_FALSE(result.has_value());
}