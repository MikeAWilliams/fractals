#include "../lib/mandelbrot.h"

#include <iostream>

#define CATCH_CONFIG_ENABLE_BENCHMARKING
#include "catch2/catch.hpp"

TEST_CASE("CalculateMandelbrot")
{
    auto result01 {ComputeMandelbrot(30, 0, 1)};
    REQUIRE_FALSE(result01.has_value());

    auto resultNeg10 {ComputeMandelbrot(30, -1, 0)};
    REQUIRE_FALSE(resultNeg10.has_value());

    auto result05 {ComputeMandelbrot(30, 0, 0.5)};
    REQUIRE_FALSE(result05.has_value());

    auto result10 {ComputeMandelbrot(30, 1, 0)};
    REQUIRE(result10.has_value());
    REQUIRE(2 == result10.value());

    auto result02 {ComputeMandelbrot(30, 0, 2)};
    REQUIRE(result02.has_value());
    REQUIRE(1 == result02.value());

    auto result55 {ComputeMandelbrot(30, 0.5, 0.5)};
    REQUIRE(result55.has_value());
    REQUIRE(5 == result55.value());
}

TEST_CASE("DrawMandelbrot")
{
    MandelbrotParameters params {30, -2.0, -1.0, 1.0, 1.0, 70, 30};
    auto asciiResult {ComputeMandelbrot<char>(params, 
        []()
        {
            return ' ';
        },
        [](int nIterations)
        {
            return 'x';
        }
        )};
    for(const auto& row : asciiResult)
    {
        for(const auto c : row)
        {
            std::cout << c;
        }
        std::cout << std::endl;
    }
}

TEST_CASE("Performance")
{
    BENCHMARK("Mandelbrot") {
        MandelbrotParameters params{30, -2.0, -1.0, 1.0, 1.0, 70, 30};
        auto asciiResult{ComputeMandelbrot<char>(
            params,
            []() {
                return ' ';
            },
            [](int nIterations) {
                return 'x';
            })};
    };
}