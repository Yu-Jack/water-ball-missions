#include "player.h"

void HumanPlayer::Naming()  {
    std::string input;
    std::cout << "What is your name?: ";
    std::cin >> input;
    std::cout << "Hello, " << input << "!" << std::endl;
    name = input;
}

void HumanPlayer::PickCard()  {
    std::cout << "your hand:" << std::endl;
    hand->List();

    std::string input;
    std::cout << "Which card do you want to pick? (card number): ";
    std::cin >> input;
    int order = std::stoi(input);

    showdown->PutCardOnTable(hand->PickCard(order), name);
}

void HumanPlayer::MakeExchangeHandDecision()  {
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

    for (IPlayer* player : showdown->players) {
        if (player->GetName() == input) {
            exchangeAble = false;
            ExchangeHand(player);
            exchangeHandData->countdown = 3;
            exchangeHandData->exchangee = player;
        }
    }

    if (exchangeAble) {
        std::cout << "Invalid name" << std::endl;
        return;
    }
}