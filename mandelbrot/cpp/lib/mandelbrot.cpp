#include "mandelbrot.h"

std::optional<int> ComputeMandelbrot(const int maxIterations, const double c0Real, const double c0Im)
{
    double x2 {0};
    double y2 {0};
    int iteration {0};
    double x {0};
    double y {0};

    for(; iteration < maxIterations && x2 + y2 < 4; ++iteration)
    {
        y = 2 * x * y + c0Im;
        x = x2 - y2 + c0Real;
        x2 = x * x;
        y2 = y * y;
    }
    return (maxIterations == iteration) ? std::optional<int>{}: iteration;
}

std::vector<std::vector<char>> ComputeAsciiMandelbrot(const MandelbrotParameters params)
{
    const double realPixelSize {(params.realMax - params.realMin)/(params.sizeX - 1)};
    const double imaginaryPixelSize {(params.imMax - params.imMin)/(params.sizeY - 1)};
    std::vector<std::vector<char>> result {params.sizeY};
    for(size_t yIndex {0}; yIndex < params.sizeY; ++yIndex)
    {
        const double imaginary {params.imMin + yIndex * imaginaryPixelSize };

        for(size_t xIndex{0}; xIndex < params.sizeX; ++xIndex)
        {
            const double real {params.realMin + xIndex * realPixelSize};
            auto pointResult {ComputeMandelbrot(params.maxIterations, real, imaginary)};
            if(pointResult.has_value())
            {
                result[yIndex].emplace_back('x');
            }
            else
            {
                result[yIndex].emplace_back(' ');
            }
        }
    }
    return result;
}