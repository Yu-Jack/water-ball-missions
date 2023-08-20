#ifndef PLAYER_H
#define PLAYER_H

#include <iostream>
#include <cstdlib>
#include <string>
#include <ctime>
#include <vector>
#include "i_player.h"
#include "../showdown.h"

using PlayerName = std::string;

struct exchangeHand {
    IPlayer* exchangee;
    int countdown;
};

class PlayerImpl : public IPlayer {

public:
    PlayerName name;
    int point;
    Hand* hand;
    bool exchangeAble;
    Showdown* showdown;
    exchangeHand* exchangeHandData;

    PlayerImpl(Showdown* s);
    PlayerName GetName() const override;
    int GetPoint() const override;
    void GainPoint() override;
    Hand* GetHands() override;
    void SetHands(Hand* h) override;
    void ExchangeHand(IPlayer* exchangee) override;
    void DrawCard(Deck* deck) override;
    void TakeTurn1() override;
    void TakeTurn2() override;

    virtual ~PlayerImpl() {
        delete hand;
        delete exchangeHandData;
    }
};

class AIPlayer : public PlayerImpl {
public:
    using PlayerImpl::PlayerImpl; // Inheriting Constructor, it's equal to below one.
    // AIPlayer(Showdown* s);
    void Naming() override ;
    void PickCard() override;
    void MakeExchangeHandDecision() override;
};

class HumanPlayer : public PlayerImpl {
public:
    using PlayerImpl::PlayerImpl; 
    // HumanPlayer(Showdown* s);
    void Naming() override;
    void PickCard() override;
    void MakeExchangeHandDecision() override;
};

#endif // PLAYER_H
