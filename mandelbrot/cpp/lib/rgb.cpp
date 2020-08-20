#include "rgb.h"

#include <utility>

SingleShadeRGB::SingleShadeRGB(int maxIt, RGB outColor, RGB inColor)
    : maxIterations {maxIt}
    , outOfSetColor{std::move(outColor)}
    , inSetColor{std::move(inColor)}
{

}

RGB SingleShadeRGB::operator()()
{
    return {0, 0, 0};
}

RGB SingleShadeRGB::operator()(int)
{
    return {0, 0, 0};
}