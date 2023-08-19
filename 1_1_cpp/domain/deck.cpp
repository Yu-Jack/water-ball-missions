
#include "deck.h"
#include <iostream>

const std::vector<int> Deck::Ranks = {
    Card::RANK2, Card::RANK3, Card::RANK4, Card::RANK5,
    Card::RANK6, Card::RANK7, Card::RANK8, Card::RANK9,
    Card::RANK10, Card::RANKJ, Card::RANKQ, Card::RANKK, Card::RANKA
};

const std::vector<int> Deck::Suits = {
    Card::SuitClub, Card::SuitDiamond, Card::SuitHeart, Card::SuitSpade
};

Deck::Deck() {
    cards.reserve(52);

    for (int r : Ranks) {
        for (int s : Suits) {
            cards.push_back(Card(r, s));
        }
    }
}

Card Deck::DrawCard() {
    Card c = cards[0]; // pop first one
    cards.erase(cards.begin());
    return c;
}

void Deck::Shuffle() {
    // do nothing for first version
    std::cout << "Shuffle" << std::endl;
}

int Deck::Size() const {
    return static_cast<int>(cards.size());
}