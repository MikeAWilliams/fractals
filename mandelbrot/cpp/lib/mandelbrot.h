#include <optional>
#include <vector>

std::optional<int> ComputeMandelbrot(const int maxIterations, const double c0Real, const double c0Im);
std::vector<std::vector<char>> ComputeAsciiMandelbrot(const int maxIterations, const double realMin, const double imMin, 
    const double realMax, const double imMax, const size_t sizeX, const size_t sizeY);

struct RGB
{
    int R;
    int G;
    int B;
};

class ColorMap
{
public:
    virtual ~ColorMap() = default;
    virtual RGB GetEscapedColor(int iterations) const = 0;
    virtual RGB GetInSetColor() const = 0;
};


std::vector<std::vector<RGB>> ComputeRGBMandelbrot(const ColorMap& colors, const int maxIterations, const double realMin, const double imMin, 
    const double realMax, const double imMax, const size_t sizeX, const size_t sizeY);