#ifndef DECK_H
#define DECK_H

#include <vector>
#include <iostream>
#include "card.h"

class Deck {
private:
    std::vector<Card> cards;

public:
    static const std::vector<int> Ranks;
    static const std::vector<int> Suits;

    Deck();
    Card DrawCard();
    void Shuffle();
    int Size() const;
};



#endif // DECK_H