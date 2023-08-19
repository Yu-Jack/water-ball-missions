#ifndef HAND_H
#define HAND_H

#include <vector>
#include <iostream>
#include "card.h"


class Hand {
private:
    std::vector<Card> cards;

public:
    void AddCard(const Card& c) {
        cards.push_back(c);
    }

    Card PickCard(int order) {
        Card c = cards[order];
        cards.erase(cards.begin() + order);
        return c;
    }

    int Size() const {
        return static_cast<int>(cards.size());
    }

    void List() const {
        for (int i = 0; i < static_cast<int>(cards.size()); ++i) {
            Card c = cards[i];
            std::cout << i << ": " << c.String() << std::endl;
        }
    }
};


#endif // HAND_H
