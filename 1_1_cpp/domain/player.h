#ifndef PLAYER_H
#define PLAYER_H

#include <iostream>
#include <cstdlib>
#include <string>
#include <ctime>
#include <vector>
#include "player_i.h"
#include "showdown.h"
#include "card.h"
#include "hand.h"
#include "deck.h"


using PlayerName = std::string;

class exchangeHand {
public:
    exchangeHand(PlayerI* exchangee, int countdown) {
        this->exchangee = exchangee;
        this->countdown = countdown;
    }

    PlayerI* exchangee;
    int countdown;
};

class PlayerImpl : public PlayerI {

public:
    PlayerName name;
    int point;
    Hand* hand;
    bool exchangeAble;
    Showdown* showdown;
    exchangeHand* exchangeHandData;

    PlayerImpl(Showdown* s) : point(0), hand(new Hand()), exchangeAble(true), showdown(s) {}

    PlayerName GetName() const override {
        return name;
    }

    int GetPoint() const override {
        return point;
    }

    void GainPoint() override {
        point++;
    }

    Hand* GetHands() override {
        return hand;
    }

    void SetHands(Hand* h) override {
        hand = h;
    }

    void ExchangeHand(PlayerI* exchangee) override {
        Hand* p1Hand = GetHands();
        Hand* p2Hand = exchangee->GetHands();

        SetHands(p2Hand);
        exchangee->SetHands(p1Hand);
    }

    void DrawCard(Deck* deck) override {
        hand->AddCard(deck->DrawCard());
    }

    void TakeTurn1(PlayerI* player) override {
        player->MakeExchangeHandDecision();
    }

    void TakeTurn2(PlayerI* player) override {
        player->PickCard();
    }

    virtual ~PlayerImpl() {
        delete hand;
        delete exchangeHandData;
    }
};

class HumanPlayer : public PlayerImpl {
public:
    HumanPlayer(Showdown* s) : PlayerImpl(s) {}

    void Naming() override {
        std::string input;
        std::cout << "What is your name?: ";
        std::cin >> input;
        std::cout << "Hello, " << input << "!" << std::endl;
        name = input;
    }

    void PickCard() override {
        std::cout << "your hand:" << std::endl;
        hand->List();

        std::string input;
        std::cout << "Which card do you want to pick? (card number): ";
        std::cin >> input;
        int order = std::stoi(input);

        showdown->PutCardOnTable(hand->PickCard(order), name);
    }

    void MakeExchangeHandDecision() override {
        if (!exchangeAble) {
            exchangeHandData->countdown--;
            if (exchangeHandData->countdown == 0) {
                ExchangeHand(exchangeHandData->exchangee);
            }
            return;
        }

        std::string input;

        while (input != "y" && input != "n") {
            if (!input.empty()) {
                std::cout << "Invalid input" << std::endl;
            }

            std::cout << "Do you want to exchange your hand? (y/n): ";
            std::cin >> input;
        }

        if (input == "n") {
            return;
        }

        std::cout << "Who do you want to exchange your hand with? (Player name): ";
        std::cin >> input;

        for (PlayerI* player : showdown->players) {
            if (player->GetName() == input) {
                exchangeAble = false;
                ExchangeHand(player);
                exchangeHandData = new exchangeHand(player, 3);
            }
        }

        if (exchangeAble) {
            std::cout << "Invalid name" << std::endl;
            return;
        }
    }
};

class AIPlayer : public PlayerImpl {
public:
    AIPlayer(Showdown* s) : PlayerImpl(s) {}

    void Naming() override {
        std::string nameStr = std::to_string(std::rand() % 100000);
        name = nameStr;
        std::cout << "Hello, " << nameStr << "!" << std::endl;
    }

    void PickCard() override {
        // always pick the first card for AI Player
        showdown->PutCardOnTable(hand->PickCard(0), name);
    }

    void MakeExchangeHandDecision() override {
        if (!exchangeAble) {
            exchangeHandData->countdown--;
            if (exchangeHandData->countdown == 0) {
                ExchangeHand(exchangeHandData->exchangee);
            }
            return;
        }

        bool exchangeAble = rand() % 2 == 1;

        if (!exchangeAble) {
            return;
        }

        exchangeAble = false;
        int order = rand() % showdown->players.size();
        PlayerI* targetPlayer = showdown->players[order];
        ExchangeHand(targetPlayer);
        exchangeHandData = new exchangeHand(targetPlayer, 3);

        std::cout << "AIPlayer: " << GetName() << " exchanged hand with " << targetPlayer->GetName() << std::endl;
    }
};

inline PlayerI* NewHumanPlayer(Showdown* showdown) {
    return new HumanPlayer(showdown);
}

inline PlayerI* NewAIPlayer(Showdown* showdown) {
    return new AIPlayer(showdown);
}


#endif // PLAYER_H
