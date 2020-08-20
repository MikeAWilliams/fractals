#include <functional>
#include <optional>
#include <vector>

std::optional<int> ComputeMandelbrot(const int maxIterations, const double c0Real, const double c0Im);

struct MandelbrotParameters
{
    int maxIterations;
    double realMin, imMin;
    double realMax,  imMax;
    size_t sizeX, sizeY;
};

template<typename pixelType>
std::vector<std::vector<pixelType>> ComputeMandelbrot(
    const MandelbrotParameters params, 
    std::function<pixelType()> getInSetPixel, 
    std::function<pixelType(int)> getOutOfSetPixel)
{
    const double realPixelSize {(params.realMax - params.realMin)/(params.sizeX - 1)};
    const double imaginaryPixelSize {(params.imMax - params.imMin)/(params.sizeY - 1)};
    std::vector<std::vector<pixelType>> result {params.sizeY};
    for(size_t yIndex {0}; yIndex < params.sizeY; ++yIndex)
    {
        const double imaginary {params.imMin + yIndex * imaginaryPixelSize };

        for(size_t xIndex{0}; xIndex < params.sizeX; ++xIndex)
        {
            const double real {params.realMin + xIndex * realPixelSize};
            const auto pointResult {ComputeMandelbrot(params.maxIterations, real, imaginary)};
            if(pointResult.has_value())
            {
                result[yIndex].emplace_back(getOutOfSetPixel(pointResult.value()));
            }
            else
            {
                result[yIndex].emplace_back(getInSetPixel());
            }
        }
    }
    return result;
}