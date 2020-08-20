#include "../lib/rgb.h"
#include "../lib/mandelbrot.h"

#include "pngwriter.h"

#include <iostream>

int main()
{
    MandelbrotParameters params{30, -2.0, -1.0, 1.0, 1.0, 1920, 1080};
    RGB outOfSetColor{65535, 65535, 65535};
    RGB inSetColor{0, 0, 0};

    SingleShadeRGB blackToWhite{params.maxIterations, std::move(outOfSetColor), std::move(inSetColor)};
    std::cout << "Computing Mandelbrot\n";
  
    auto rgbResult {ComputeMandelbrot<RGB>(params, blackToWhite, blackToWhite)};
    
    std::cout << "Done computing Mandelbrot\n";

    pngwriter outFile(params.sizeX, params.sizeY, 0, "out.png"); 

    for(int y{0}; y < params.sizeY; ++y)
    {
        for(int x{0}; x < params.sizeX; ++x)
        {
            RGB& rgb {rgbResult[y][x]};
            //std::cout << rgb.r << " " << rgb.g << " " << rgb.b << std::endl;
            outFile.plot(x, y, rgb.r, rgb.g, rgb.b);
            //outFile.plot(x, y, 65535, 0, 0);
        }
    }

    outFile.close();
}