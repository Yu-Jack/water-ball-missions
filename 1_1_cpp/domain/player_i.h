#ifndef PLAYER_I_H
#define PLAYER_I_H

#include <string>
#include "card.h"
#include "hand.h"
#include "deck.h"

using PlayerName = std::string;
class PlayerI {
public:
    virtual void Naming() = 0;
    virtual void GainPoint() = 0;
    virtual void PickCard() = 0;
    virtual void ExchangeHand(PlayerI* exchangee) = 0;
    virtual void MakeExchangeHandDecision() = 0;
    virtual void DrawCard(Deck* deck) = 0;
    virtual PlayerName GetName() const = 0;
    virtual Hand* GetHands() = 0;
    virtual void SetHands(Hand* hand) = 0;
    virtual int GetPoint() const = 0;
    virtual void TakeTurn1(PlayerI* player) = 0;
    virtual void TakeTurn2(PlayerI* player) = 0;
    virtual ~PlayerI() {}
};

#endif // PLAYER_I_H
