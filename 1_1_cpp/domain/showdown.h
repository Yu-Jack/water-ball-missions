#ifndef SHOWDOWN_H
#define SHOWDOWN_H

#include <map>
#include <iostream>
#include "player_i.h"
#include "card.h"
#include "hand.h"
#include "deck.h"

using PlayerName = std::string;

const int MaxTurn = 13;
class Showdown {

public:
    int turn;
    Deck* deck;
    std::vector<PlayerI*> players;
    std::map<PlayerName, Card> table; // key is a HumanPlayer name

    Showdown(Deck* d) : turn(0), deck(d) {}

    void AddPlayer(PlayerI* p) {
        players.push_back(p);
    }

    void Start() {
        for (PlayerI* p : players) {
            p->Naming();
        }

        deck->Shuffle();

        while (deck->Size() > 0) {
            for (PlayerI* p : players) {
                p->DrawCard(deck);
            }
        }

        // Modify the game rule to prevent all hands from being empty
        while (turn < MaxTurn) {
            // First action: decide whether to exchange hands
            for (PlayerI* p : players) {
                p->TakeTurn1(p);
            }

            // Second action: play a card
            for (PlayerI* p : players) {
                p->TakeTurn2(p);
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
            for (PlayerI* p : players) {
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
        for (PlayerI* p : players) {
            std::cout << "Player gets " << p->GetName() << " point " << p->GetPoint() << "!" << std::endl;
            if (p->GetPoint() > point) {
                point = p->GetPoint();
                winner = p->GetName();
            }
        }

        std::cout << "Winner is " << winner << "! Point is " << point << std::endl;
    }

    void PutCardOnTable(const Card& c, const PlayerName& p) {
        table[p] = c;
    }
};


#endif // SHOWDOWN_H
