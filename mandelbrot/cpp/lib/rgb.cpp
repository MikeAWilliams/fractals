#include "rgb.h"

#include <utility>
#include <iostream>

SingleShadeRGB::SingleShadeRGB(int maxIt, RGB outColor, RGB inColor)
    : maxIterations {maxIt}
    , inSetColor{std::move(inColor)}
    , colorDelta{outColor.r - inSetColor.r, outColor.g - inSetColor.g, outColor.b - inSetColor.b}
{

}

RGB SingleShadeRGB::operator()() const
{
    return inSetColor;
}

RGB SingleShadeRGB::operator()(int iter) const
{
    double maxFraction {static_cast<double>(iter) / static_cast<double>(maxIterations)};
    //std::cout << iter << " " << maxFraction << std::endl;
    return {static_cast<int>(colorDelta.r*maxFraction), 
            static_cast<int>(colorDelta.g*maxFraction), 
            static_cast<int>(colorDelta.b*maxFraction)};
}