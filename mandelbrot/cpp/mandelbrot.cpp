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