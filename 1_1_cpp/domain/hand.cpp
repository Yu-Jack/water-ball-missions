#include <vector>
#include <iostream>
#include "hand.h"
#include "card.h"


void Hand::AddCard(const Card& c) {
    cards.push_back(c);
}

Card Hand::PickCard(int order) {
    Card c = cards[order];
    cards.erase(cards.begin() + order);
    return c;
}

int Hand::Size() const {
    return static_cast<int>(cards.size());
}

void Hand::List() const {
    for (int i = 0; i < static_cast<int>(cards.size()); ++i) {
        Card c = cards[i];
        std::cout << i << ": " << c.String() << std::endl;
    }
}