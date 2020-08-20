#include "../lib/rgb.h"

#include <utility>

#include "catch2/catch.hpp"

TEST_CASE("singlecolorblackwhite")
{
    RGB outOfSetColor {255, 255, 255};
    RGB inSetColor {0, 0, 0};

    const int maxIterations {10};

    SingleShadeRGB blackToWhite{maxIterations, std::move(outOfSetColor), std::move(inSetColor)};
    const auto inSetResult {blackToWhite()};
    REQUIRE(0 == inSetResult.r);
    REQUIRE(0 == inSetResult.g);
    REQUIRE(0 == inSetResult.b);

    const auto halfOutResult {blackToWhite(5)};
    REQUIRE(127 == halfOutResult.r);
    REQUIRE(127 == halfOutResult.g);
    REQUIRE(127 == halfOutResult.b);
}
