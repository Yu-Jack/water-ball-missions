
#include "showdown.h"

const int MaxTurn = 13;

using PlayerName = std::string;

Showdown::Showdown(Deck* d) {
    deck = d;
}

void Showdown::AddPlayer(IPlayer* p) {
    players.push_back(p);
}

void Showdown::Start() {
    for (IPlayer* p : players) {
        p->Naming();
    }

    deck->Shuffle();

    while (deck->Size() > 0) {
        for (IPlayer* p : players) {
            p->DrawCard(deck);
        }
    }

    // Modify the game rule to prevent all hands from being empty
    while (turn < MaxTurn) {
        // First action: decide whether to exchange hands
        for (IPlayer* p : players) {
            p->TakeTurn1();
        }

        // Second action: play a card
        for (IPlayer* p : players) {
            p->TakeTurn2();
        }

        // Find which player has the bigger card
        PlayerName biggerOne = "";
        for (const auto& pair : table) {
            Card c = pair.second;
            std::cout << "Player: " << pair.first << " shows " << c.String() << std::endl;

            if (biggerOne == "") {
                biggerOne = pair.first;
                continue;
            }

            if (c.Bigger(table[biggerOne])) {
                biggerOne = pair.first;
            }
        }

        // Give the winner a point
        for (IPlayer* p : players) {
            if (p->GetName() == biggerOne) {
                std::cout << p->GetName() << " got a point!" << std::endl;
                p->GainPoint();
            }
        }

        turn++;
    }

    // Find the winner
    int point = 0;
    PlayerName winner = "";
    for (IPlayer* p : players) {
        std::cout << "Player gets " << p->GetName() << " point " << p->GetPoint() << "!" << std::endl;
        if (p->GetPoint() > point) {
            point = p->GetPoint();
            winner = p->GetName();
        }
    }

    std::cout << "Winner is " << winner << "! Point is " << point << std::endl;
}
void Showdown::PutCardOnTable(const Card& c, const PlayerName& p) {
    table[p] = c;
}