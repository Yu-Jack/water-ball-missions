#include <iostream>
#include "./domain/deck.h"
#include "./domain/player/player.h"
#include "./domain/showdown.h"

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