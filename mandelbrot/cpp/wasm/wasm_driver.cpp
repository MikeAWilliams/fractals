#include "../lib/rgb.h"
#include "../lib/mandelbrot.h"

#include <iostream>

RGB* globalRGB = nullptr;

extern "C" { RGB* simplified_mandelbrot(int realMin, int imMin, int realMax, int imMax, int sizeX, int sizeY)
{
    delete globalRGB;
    MandelbrotParameters params
        {30
        ,static_cast<double>(realMin)
        ,static_cast<double>(imMin)
        ,static_cast<double>(realMax)
        ,static_cast<double>(imMax)
        ,static_cast<size_t>(sizeX)
        ,static_cast<size_t>(sizeY)};
    RGB outOfSetColor{255, 255, 255};
    RGB inSetColor{0, 0, 0};

    SingleShadeRGB blackToWhite{params.maxIterations, std::move(outOfSetColor), std::move(inSetColor)};
    std::cout << "Computing Mandelbrot\n";
    
    auto rgbResult {ComputeMandelbrot<RGB>(params, blackToWhite, blackToWhite)};
    
    std::cout << "Done computing Mandelbrot.\n";
    globalRGB = new RGB[sizeX * sizeY];

    for(int y{0}; y < params.sizeY; ++y)
    {
        for(int x{0}; x < params.sizeX; ++x)
        {
            RGB& rgb {rgbResult[y][x]};
            int index = y * params.sizeX + x;
            globalRGB[index].r = rgb.r;
            globalRGB[index].g = rgb.g;
            globalRGB[index].b = rgb.b;
        }
    }
    std::cout << "Done Plotting\n";
    return globalRGB;
}}