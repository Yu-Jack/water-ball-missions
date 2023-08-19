#ifndef HAND_H
#define HAND_H

#include <vector>
#include "card.h"


class Hand {
private:
    std::vector<Card> cards;

public:
    void AddCard(const Card& c);
    Card PickCard(int order);
    int Size() const;
    void List() const;
};


#endif // HAND_H
