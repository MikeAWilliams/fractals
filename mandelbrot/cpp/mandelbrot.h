#include <optional>
#include <vector>

std::optional<int> ComputeMandelbrot(const int maxIterations, const double c0Real, const double c0Im);
std::vector<std::vector<char>> ComputeAsciiMandelbrot(const int maxIterations, const double realMin, const double imMin, 
    const double realMax, const double imMax, const size_t sizeX, const size_t sizeY);