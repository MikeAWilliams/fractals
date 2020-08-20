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

std::vector<std::vector<char>> ComputeAsciiMandelbrot(const int maxIterations, const double realMin, const double imMin, 
    const double realMax, const double imMax, const size_t sizeX, const size_t sizeY)
{
    const double realPixelSize {(realMax - realMin)/(sizeX - 1)};
    const double imaginaryPixelSize {(imMax - imMin)/(sizeY - 1)};
    std::vector<std::vector<char>> result {sizeY};
    for(size_t yIndex {0}; yIndex < sizeY; ++yIndex)
    {
        const double imaginary {imMin + yIndex * imaginaryPixelSize };

        for(size_t xIndex{0}; xIndex < sizeX; ++xIndex)
        {
            const double real {realMin + xIndex * realPixelSize};
            auto pointResult {ComputeMandelbrot(maxIterations, real, imaginary)};
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