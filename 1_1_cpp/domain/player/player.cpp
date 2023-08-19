#include "player.h"

PlayerImpl::PlayerImpl(Showdown* s): point(0), hand(new Hand()), exchangeAble(true), showdown(s), exchangeHandData(new exchangeHand()) {}


PlayerName PlayerImpl::GetName() const {
    return name;
}

int PlayerImpl::GetPoint() const  {
    return point;
}
void PlayerImpl::GainPoint()  {
    point++;
}
Hand* PlayerImpl::GetHands()  {
    return hand;
}
void PlayerImpl::SetHands(Hand* h)  {
    hand = h;
}

void PlayerImpl::ExchangeHand(IPlayer* exchangee) {
    Hand* p1Hand = GetHands();
    Hand* p2Hand = exchangee->GetHands();

    SetHands(p2Hand);
    exchangee->SetHands(p1Hand);
}

void PlayerImpl::DrawCard(Deck* deck) {
    hand->AddCard(deck->DrawCard());
}

void PlayerImpl::TakeTurn1(IPlayer* player) {
    player->MakeExchangeHandDecision();
}

void PlayerImpl::TakeTurn2(IPlayer* player) {
    player->PickCard();
}