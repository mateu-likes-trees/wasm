#include <stdio.h>
#include <string>

using namespace std;

extern "C"
{
    int add(int first, int second)
    {
        return first + second;
    }

    const char *greet()
    {
        return "hello";
    }
}
