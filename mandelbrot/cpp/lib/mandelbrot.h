#include <optional>
#include <vector>

std::optional<int> ComputeMandelbrot(const int maxIterations, const double c0Real, const double c0Im);

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

struct MandelbrotParameters
{
    int maxIterations;
    double realMin, imMin;
    double realMax,  imMax;
    size_t sizeX, sizeY;
};

std::vector<std::vector<char>> ComputeAsciiMandelbrot(const MandelbrotParameters params);