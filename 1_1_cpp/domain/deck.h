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

    Deck() {
        cards.reserve(52);

        for (int r : Ranks) {
            for (int s : Suits) {
                cards.push_back(Card(r, s));
            }
        }
    }

    Card DrawCard() {
        Card c = cards[0]; // pop first one
        cards.erase(cards.begin());
        return c;
    }

    void Shuffle() {
        // do nothing for first version
        std::cout << "Shuffle" << std::endl;
    }

    int Size() const {
        return static_cast<int>(cards.size());
    }
};

const std::vector<int> Deck::Ranks = {
    Card::RANK2, Card::RANK3, Card::RANK4, Card::RANK5,
    Card::RANK6, Card::RANK7, Card::RANK8, Card::RANK9,
    Card::RANK10, Card::RANKJ, Card::RANKQ, Card::RANKK, Card::RANKA
};

const std::vector<int> Deck::Suits = {
    Card::SuitClub, Card::SuitDiamond, Card::SuitHeart, Card::SuitSpade
};


#endif // DECK_H