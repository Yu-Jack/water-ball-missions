#include <iostream>
#include "./domain/showdown.h"
#include "./domain/player.h"
#include "./domain/card.h"
#include "./domain/hand.h"
#include "./domain/deck.h"


int main() {
    std::srand(static_cast<unsigned int>(std::time(nullptr)));

    Deck deck;
    Showdown showdown(&deck);

    showdown.AddPlayer(new HumanPlayer(&showdown));
    showdown.AddPlayer(new AIPlayer(&showdown));
    showdown.AddPlayer(new AIPlayer(&showdown));
    showdown.AddPlayer(new AIPlayer(&showdown));

    showdown.Start();

    return 0;
}