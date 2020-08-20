#pragma once
struct RGB
{
    int r, g, b;
};

class SingleShadeRGB
{
public:
    SingleShadeRGB(int maxIt,RGB outColor, RGB inColor);
    RGB operator()();
    RGB operator()(int);
private:
    int maxIterations;
    RGB outOfSetColor;
    RGB inSetColor;

};