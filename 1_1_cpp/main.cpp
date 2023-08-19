#include <iostream>
#include "./domain/card.h"

int main() {

    Card mycard1 = Card(5, 3);
    Card mycard2 = Card(2, 1);

    std::cout << mycard1.Bigger(mycard2) << std::endl;

    std::cout << mycard1.String() << std::endl;
    std::cout << mycard2.String() << std::endl;

    return 0;
}