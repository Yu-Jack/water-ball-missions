#ifndef SHOWDOWN_H
#define SHOWDOWN_H

#include <map>
#include <iostream>
#include "./player/i_player.h"


class Showdown {

public:
    int turn;
    Deck* deck;
    std::vector<IPlayer*> players;
    std::map<PlayerName, Card> table; // key is a HumanPlayer name

    Showdown(Deck* d);
    void AddPlayer(IPlayer* p);
    void Start();
    void PutCardOnTable(const Card& c, const PlayerName& p);
};


#endif // SHOWDOWN_H
