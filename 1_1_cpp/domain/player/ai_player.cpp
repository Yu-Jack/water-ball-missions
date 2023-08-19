#include "player.h"

AIPlayer::AIPlayer(Showdown* s): PlayerImpl(s) {}

void AIPlayer::Naming() {
    std::string nameStr = std::to_string(std::rand() % 100000);
    name = nameStr;
    std::cout << "Hello, " << nameStr << "!" << std::endl;
}

void AIPlayer::PickCard() {
    // always pick the first card for AI Player
    showdown->PutCardOnTable(hand->PickCard(0), name);
}

void AIPlayer::MakeExchangeHandDecision() {
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
    IPlayer* targetPlayer = showdown->players[order];
    ExchangeHand(targetPlayer);
    exchangeHandData->countdown = 3;
    exchangeHandData->exchangee = targetPlayer;

    std::cout << "AIPlayer: " << GetName() << " exchanged hand with " << targetPlayer->GetName() << std::endl;
}