#ifndef SHOWDOWN_H
#define SHOWDOWN_H

#include <map>
#include <iostream>
#include "player_i.h"
#include "card.h"
#include "hand.h"
#include "deck.h"


class Showdown {

public:
    int turn;
    Deck* deck;
    std::vector<PlayerI*> players;
    std::map<PlayerName, Card> table; // key is a HumanPlayer name

    Showdown(Deck* d);
    void AddPlayer(PlayerI* p);
    void Start();
    void PutCardOnTable(const Card& c, const PlayerName& p);
};


#endif // SHOWDOWN_H
